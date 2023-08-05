package service

import (
	"context"
	"errors"
	"fmt"
	"route256/libs/mw/tracing"

	"github.com/opentracing/opentracing-go"
)

func (s *Service) Purchase(
	ctx context.Context,
	userID int64,
) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "domain/Purchase")
	defer span.Finish()

	span.SetTag("user_id", userID)

	cartItems, err := s.repo.GetItems(ctx, userID)
	if err != nil {
		return 0, fmt.Errorf("error getting cart items: %w", err)
	}

	if len(cartItems) == 0 {
		return 0, tracing.MarkSpanWithError(ctx, errors.New("empty cart"))
	}

	orderID, err := s.lomsClient.CreateOrder(ctx, userID, cartItems)
	if err != nil {
		return 0, fmt.Errorf("failed to create order for user %d", userID)
	}

	err = s.repo.DeleteItems(ctx, userID)
	if err != nil {
		return orderID, fmt.Errorf("failed to clean cart for user %d", userID)
	}

	return orderID, nil
}
