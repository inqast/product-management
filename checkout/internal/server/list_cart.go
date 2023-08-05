package server

import (
	"context"
	"route256/checkout/internal/domain/model"
	"route256/checkout/internal/server/mapping"
	api "route256/checkout/pkg/checkout/v1"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListCart(ctx context.Context, req *api.ListCartRequest) (*api.ListCartResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server/ListCart")
	defer span.Finish()

	span.SetTag("user_id", req.User)

	cart, err := s.service.ListCart(ctx, req.User)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return mapContract(cart), nil
}

func mapContract(cart *model.Cart) *api.ListCartResponse {
	return &api.ListCartResponse{
		Items:      mapping.MapContractCartItems(cart.Items),
		TotalPrice: cart.TotalPrice,
	}
}
