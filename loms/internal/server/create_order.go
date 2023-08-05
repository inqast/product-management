package server

import (
	"context"
	"route256/loms/internal/server/mapping"
	api "route256/loms/pkg/loms/v1"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateOrder(ctx context.Context, req *api.CreateOrderRequest) (*api.CreateOrderResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server/CreateOrder")
	defer span.Finish()

	span.SetTag("user_id", req.User)

	orderId, err := s.service.CreateOrder(ctx, req.User, mapping.MapDomainOrderItems(req.Items))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &api.CreateOrderResponse{
		OrderID: orderId,
	}, nil
}
