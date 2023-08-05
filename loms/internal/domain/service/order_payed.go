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

func (s *Service) OrderPayed(
	ctx context.Context,
	orderID int64,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "domain/OrderPayed")
	defer span.Finish()

	span.SetTag("order_id", orderID)

	order, err := s.orderRepo.GetOrder(ctx, orderID)
	if err != nil {
		return fmt.Errorf("error getting order from db id %d: %w", orderID, err)
	}

	if order.Status != model.OrderAwaitingPayment {
		return tracing.MarkSpanWithError(ctx, errors.New("incorrect order status"))
	}

	err = s.transactionManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		for _, item := range order.Items {
			reservesToBuy, err := s.getReservesToBuy(ctxTx, item)
			if err != nil {
				return fmt.Errorf("failed to get reserves for sku %d: %w", item.SKU, err)
			}

			err = s.buyReserves(ctxTx, item.SKU, reservesToBuy)
			if err != nil {
				return err
			}
		}

		err = s.orderRepo.SetOrderStatus(ctxTx, order.ID, model.OrderPayed)
		if err != nil {
			return fmt.Errorf("error setting order status: %w", err)
		}

		if err := s.statusSender.SendMessage(ctx, order.UserID, order.ID, model.OrderPayed); err != nil {
			log.Warnf("error sending status to broker: %s", err.Error())
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to cancel order: %w", err)
	}

	return nil
}

func (s *Service) buyReserves(ctx context.Context, sku uint32, stocksToBuy []*model.Stock) error {
	for _, stockToBuy := range stocksToBuy {
		if stockToBuy.Count == 0 && stockToBuy.Reserved == 0 {
			if err := s.stocksRepo.DeleteStock(
				ctx,
				sku,
				stockToBuy.WarehouseID,
			); err != nil {
				return fmt.Errorf("failed to delete empty stock for sku %d: %w", sku, err)
			}

			continue
		}

		if err := s.stocksRepo.SetStock(
			ctx,
			sku,
			stockToBuy.WarehouseID,
			stockToBuy.Count,
			stockToBuy.Reserved,
		); err != nil {
			return fmt.Errorf("failed to buy items on stock for sku %d: %w", sku, err)
		}
	}

	return nil
}

func (s *Service) getReservesToBuy(ctx context.Context, item *model.OrderItem) ([]*model.Stock, error) {
	stocks, err := s.stocksRepo.GetStocks(ctx, item.SKU)
	if err != nil {
		return nil, err
	}

	var (
		boughtCount uint64
	)

	stockToBuy := make([]*model.Stock, 0, len(stocks))

	for _, stock := range stocks {
		left := uint64(item.Count) - boughtCount
		if left == 0 {
			break
		}

		if stock.Reserved >= left {
			stockToBuy = append(stockToBuy, &model.Stock{
				WarehouseID: stock.WarehouseID,
				Count:       stock.Count,
				Reserved:    stock.Reserved - left,
			})
			boughtCount += left
		} else {
			stockToBuy = append(stockToBuy, &model.Stock{
				WarehouseID: stock.WarehouseID,
				Count:       stock.Count,
				Reserved:    0,
			})
			boughtCount += stock.Reserved
		}
	}

	if boughtCount != uint64(item.Count) {
		return nil, tracing.MarkSpanWithError(ctx, ErrNotEnoughItems)
	}

	return stockToBuy, nil
}
