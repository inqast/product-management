package service

import (
	"context"
	"errors"
	"fmt"
	log "route256/libs/logger"
	"route256/loms/internal/domain/model"

	"github.com/opentracing/opentracing-go"
)

var ErrNotEnoughItems = errors.New("not enough items")

func (s *Service) CreateOrder(
	ctx context.Context,
	userID int64,
	items []*model.OrderItem,
) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "domain/CreateOrder")
	defer span.Finish()

	span.SetTag("user_id", userID)

	var orderId int64

	err := s.transactionManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		// Костылик от затенения.
		var orderErr error
		orderId, orderErr = s.orderRepo.CreateOrder(ctxTx, userID)
		if orderErr != nil {
			return orderErr
		}

		if err := s.statusSender.SendMessage(ctx, userID, orderId, model.OrderNew); err != nil {
			log.Warnf("error sending status to broker: %s", err.Error())
		}

		err := s.addOrderItems(ctxTx, orderId, items)
		if err != nil {
			return fmt.Errorf("failed to ser order items: %w", err)
		}

		return nil
	})
	if err != nil {
		return 0, fmt.Errorf("failed to create order: %w", err)
	}

	err = s.transactionManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		for _, item := range items {
			stocksToReserve, err := s.getStocksToReserve(ctxTx, item)
			if err != nil {
				return fmt.Errorf("failed to get reserve for sku %d: %w", item.SKU, err)
			}

			err = s.reserveStock(ctxTx, item.SKU, stocksToReserve)
			if err != nil {
				return err
			}
		}

		err = s.orderRepo.SetOrderStatus(ctxTx, orderId, model.OrderAwaitingPayment)
		if err != nil {
			return fmt.Errorf("failed to ser order status awaiting payment: %w", err)
		}

		if err := s.statusSender.SendMessage(ctx, userID, orderId, model.OrderAwaitingPayment); err != nil {
			log.Warnf("error sending status to broker: %s", err.Error())
		}

		return nil
	})
	if err != nil {
		if statusErr := s.orderRepo.SetOrderStatus(ctx, orderId, model.OrderFailed); statusErr != nil {
			log.Warnf("failed to set status orderFailed: %s", statusErr.Error())
		}

		if err := s.statusSender.SendMessage(ctx, userID, orderId, model.OrderFailed); err != nil {
			log.Warnf("error sending status to broker: %s", err.Error())
		}

		return 0, fmt.Errorf("failed to create order: %w", err)
	}

	return orderId, nil
}

func (s *Service) addOrderItems(ctx context.Context, orderId int64, items []*model.OrderItem) error {
	for _, item := range items {
		err := s.orderRepo.AddOrderItem(ctx, orderId, item.SKU, item.Count)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) reserveStock(ctx context.Context, sku uint32, stocksToReserve []*model.Stock) error {
	for _, stockToReserve := range stocksToReserve {
		if err := s.stocksRepo.SetStock(
			ctx,
			sku,
			stockToReserve.WarehouseID,
			stockToReserve.Count,
			stockToReserve.Reserved,
		); err != nil {
			return fmt.Errorf("failed to create reserve for sku %d: %w", sku, err)
		}
	}

	return nil
}

func (s *Service) getStocksToReserve(ctx context.Context, item *model.OrderItem) ([]*model.Stock, error) {
	stocks, err := s.stocksRepo.GetStocks(ctx, item.SKU)
	if err != nil {
		return nil, err
	}

	var (
		reservedCount uint64
	)

	stocksToReserve := make([]*model.Stock, 0, len(stocks))

	for _, stock := range stocks {
		left := uint64(item.Count) - reservedCount
		if left == 0 {
			break
		}

		if stock.Count >= left {
			stocksToReserve = append(stocksToReserve, &model.Stock{
				WarehouseID: stock.WarehouseID,
				Count:       stock.Count - left,
				Reserved:    stock.Reserved + left,
			})
			reservedCount += left
		} else {
			stocksToReserve = append(stocksToReserve, &model.Stock{
				WarehouseID: stock.WarehouseID,
				Count:       0,
				Reserved:    stock.Reserved + stock.Count,
			})
			reservedCount += stock.Count
		}
	}

	if reservedCount != uint64(item.Count) {
		return nil, ErrNotEnoughItems
	}

	for _, stockToReserve := range stocksToReserve {
		fmt.Println(stockToReserve)
	}

	return stocksToReserve, nil
}
