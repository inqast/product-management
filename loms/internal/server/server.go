package server

import (
	"context"
	"route256/loms/internal/domain/model"
	api "route256/loms/pkg/loms/v1"
)

type service interface {
	OrderCancel(ctx context.Context, orderId int64) error
	CreateOrder(ctx context.Context, userId int64, items []*model.OrderItem) (int64, error)
	ListOrder(ctx context.Context, orderId int64) (*model.Order, error)
	OrderPayed(ctx context.Context, orderId int64) error
	Stocks(ctx context.Context, sku uint32) ([]*model.Stock, error)
}

type Server struct {
	api.UnimplementedLomsServer
	service service
}

func New(service service) *Server {
	return &Server{
		service: service,
	}
}
