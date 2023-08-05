package service

import (
	"context"
	"fmt"
	"route256/libs/mw/tracing"

	"github.com/opentracing/opentracing-go"
)

func (s *Service) DeleteFromCart(
	ctx context.Context,
	userID int64,
	SKU uint32,
	count uint16,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "domain/DeleteFromCart")
	defer span.Finish()

	span.SetTag("user_id", userID)

	currentCount, err := s.repo.GetItemCount(ctx, userID, SKU)
	if err != nil {
		return fmt.Errorf("error getting current count for SKU %d: %w", SKU, err)
	}

	switch {
	case currentCount < count:
		return tracing.MarkSpanWithError(ctx, fmt.Errorf("not enough items in cart SKU %d", SKU))
	case currentCount == count:
		err = s.repo.DeleteItem(ctx, userID, SKU)
		if err != nil {
			return fmt.Errorf("error adding item to cart for SKU %d: %w", SKU, err)
		}
	default:
		err = s.repo.SetItemCount(ctx, userID, SKU, currentCount-count)
		if err != nil {
			return fmt.Errorf("error adding item to cart for SKU %d: %w", SKU, err)
		}
	}

	return nil
}
