package server

import (
	"context"
	api "route256/checkout/pkg/checkout/v1"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Purchase(ctx context.Context, req *api.PurchaseRequest) (*api.PurchaseResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server/Purchase")
	defer span.Finish()

	span.SetTag("user_id", req.User)

	orderID, err := s.service.Purchase(
		ctx,
		req.User,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &api.PurchaseResponse{
		OrderId: orderID,
	}, nil
}
