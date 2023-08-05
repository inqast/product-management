package service

import (
	"context"
	"errors"
	"fmt"
	log "route256/libs/logger"
	"route256/libs/mw/tracing"
	"route256/loms/internal/domain/model"

	"github.com/opentracing/opentracing-go"
)

func (s *Service) OrderCancel(
	ctx context.Context,
	orderID int64,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "domain/OrderCancel")
	defer span.Finish()

	span.SetTag("order_id", orderID)

	order, err := s.orderRepo.GetOrder(ctx, orderID)
	if err != nil {
		return fmt.Errorf("error getting order from db id %d: %w", orderID, err)
	}

	if order.Status == model.OrderPayed {
		return tracing.MarkSpanWithError(ctx, errors.New("incorrect order status"))
	}

	err = s.transactionManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		for _, item := range order.Items {
			reservesToCancel, err := s.getReservesToCancel(ctxTx, item)
			if err != nil {
				return fmt.Errorf("failed to get reserves for sku %d: %w", item.SKU, err)
			}

			err = s.cancelReserves(ctxTx, item.SKU, reservesToCancel)
			if err != nil {
				return err
			}
		}

		err = s.orderRepo.SetOrderStatus(ctxTx, order.ID, model.OrderCancelled)
		if err != nil {
			return fmt.Errorf("error setting order status: %w", err)
		}

		if err := s.statusSender.SendMessage(ctx, order.UserID, order.ID, model.OrderCancelled); err != nil {
			log.Warnf("error sending status to broker: %s", err.Error())
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to cancel order: %w", err)
	}

	return nil
}

func (s *Service) cancelReserves(ctx context.Context, sku uint32, stocksToCancelReserve []*model.Stock) error {
	for _, stockToCancelReserve := range stocksToCancelReserve {
		if err := s.stocksRepo.SetStock(
			ctx,
			sku,
			stockToCancelReserve.WarehouseID,
			stockToCancelReserve.Count,
			stockToCancelReserve.Reserved,
		); err != nil {
			return fmt.Errorf("failed to cancel reserve for sku %d: %w", sku, err)
		}
	}

	return nil
}

func (s *Service) getReservesToCancel(ctx context.Context, item *model.OrderItem) ([]*model.Stock, error) {
	stocks, err := s.stocksRepo.GetStocks(ctx, item.SKU)
	if err != nil {
		return nil, err
	}

	var (
		cancelledCount uint64
	)

	stockToCancelReserve := make([]*model.Stock, 0, len(stocks))

	for _, stock := range stocks {
		left := uint64(item.Count) - cancelledCount
		if left == 0 {
			break
		}

		if stock.Reserved >= left {
			stockToCancelReserve = append(stockToCancelReserve, &model.Stock{
				WarehouseID: stock.WarehouseID,
				Count:       stock.Count + left,
				Reserved:    stock.Reserved - left,
			})
			cancelledCount += left
		} else {
			stockToCancelReserve = append(stockToCancelReserve, &model.Stock{
				WarehouseID: stock.WarehouseID,
				Count:       stock.Count + stock.Reserved,
				Reserved:    0,
			})
			cancelledCount += stock.Reserved
		}
	}

	if cancelledCount != uint64(item.Count) {
		return nil, ErrNotEnoughItems
	}

	return stockToCancelReserve, nil
}
