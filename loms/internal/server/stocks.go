package server

import (
	"context"
	"route256/loms/internal/domain/model"
	"route256/loms/internal/server/mapping"
	api "route256/loms/pkg/loms/v1"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Stocks(ctx context.Context, req *api.StocksRequest) (*api.StocksResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server/Stocks")
	defer span.Finish()

	stocks, err := s.service.Stocks(ctx, req.Sku)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return mapStocksRequest(stocks), nil
}

func mapStocksRequest(stocks []*model.Stock) *api.StocksResponse {
	return &api.StocksResponse{
		Stocks: mapping.MapContractStocks(stocks),
	}
}
