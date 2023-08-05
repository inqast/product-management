package service

import (
	"context"
	"fmt"
	"route256/loms/internal/domain/model"

	"github.com/opentracing/opentracing-go"
)

func (s *Service) ListOrder(
	ctx context.Context,
	orderID int64,
) (*model.Order, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "domain/ListOrder")
	defer span.Finish()

	span.SetTag("order_id", orderID)

	order, err := s.orderRepo.GetOrder(ctx, orderID)
	if err != nil {
		return nil, fmt.Errorf("error getting order from db id %d: %w", orderID, err)
	}

	return order, nil
}
