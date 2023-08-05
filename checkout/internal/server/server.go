package server

import (
	"context"
	"route256/checkout/internal/domain/model"
	api "route256/checkout/pkg/checkout/v1"
)

type service interface {
	AddToCart(ctx context.Context, userId int64, sku uint32, count uint16) error
	DeleteFromCart(ctx context.Context, userId int64, sku uint32, count uint16) error
	ListCart(ctx context.Context, orderId int64) (*model.Cart, error)
	Purchase(ctx context.Context, userId int64) (int64, error)
}

type Server struct {
	api.UnimplementedCheckoutServer
	service service
}

func New(service service) *Server {
	return &Server{
		service: service,
	}
}
