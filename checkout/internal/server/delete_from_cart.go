package server

import (
	"context"
	api "route256/checkout/pkg/checkout/v1"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteFromCart(ctx context.Context, req *api.DeleteFromCartRequest) (*emptypb.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server/DeleteFromCart")
	defer span.Finish()

	span.SetTag("user_id", req.User)

	err := s.service.DeleteFromCart(
		ctx,
		req.User,
		req.Sku,
		uint16(req.Count),
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
