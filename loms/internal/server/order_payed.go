package server

import (
	"context"
	api "route256/loms/pkg/loms/v1"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) OrderPayed(ctx context.Context, req *api.OrderPayedRequest) (*emptypb.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server/OrderPayed")
	defer span.Finish()

	span.SetTag("order_id", req.OrderID)

	err := s.service.OrderPayed(ctx, req.OrderID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
