package service

import (
	"context"
	"fmt"
	"route256/loms/internal/domain/model"

	"github.com/opentracing/opentracing-go"
)

func (s *Service) Stocks(ctx context.Context, SKU uint32) ([]*model.Stock, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "domain/Stocks")
	defer span.Finish()

	span.SetTag("sku", SKU)

	stocks, err := s.stocksRepo.GetStocks(ctx, SKU)
	if err != nil {
		return nil, fmt.Errorf("error getting stocks from db: %w", err)
	}

	return stocks, nil
}
