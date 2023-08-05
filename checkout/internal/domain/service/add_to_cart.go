package service

import (
	"context"
	"fmt"
	"route256/libs/mw/tracing"

	"github.com/opentracing/opentracing-go"
)

func (s *Service) AddToCart(
	ctx context.Context,
	userID int64,
	SKU uint32,
	count uint16,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "domain/AddToCart")
	defer span.Finish()

	span.SetTag("user_id", userID)

	stocks, err := s.lomsClient.Stocks(ctx, SKU)
	if err != nil {
		return fmt.Errorf("error processing loms Stocks for SKU %d: %w", SKU, err)
	}

	availableItems := uint64(0)
	for _, stock := range stocks {
		availableItems += stock.Count
	}

	currentCount, err := s.repo.GetItemCount(ctx, userID, SKU)
	if err != nil {
		return fmt.Errorf("error getting current count for SKU %d: %w", SKU, err)
	}

	if availableItems < uint64(count)+uint64(currentCount) {
		return tracing.MarkSpanWithError(ctx, fmt.Errorf("not enough items on stocks for SKU %d", SKU))
	}

	err = s.repo.SetItemCount(ctx, userID, SKU, count+currentCount)
	if err != nil {
		return fmt.Errorf("error adding item to cart for SKU %d: %w", SKU, err)
	}

	return nil
}
