package server

import (
	"context"
	"errors"
	"route256/libs/mw/tracing"
	"route256/loms/internal/domain/model"
	commonMapping "route256/loms/internal/mapping"
	"route256/loms/internal/server/mapping"
	api "route256/loms/pkg/loms/v1"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListOrder(ctx context.Context, req *api.ListOrderRequest) (*api.ListOrderResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server/ListOrder")
	defer span.Finish()

	span.SetTag("order_id", req.OrderID)

	order, err := s.service.ListOrder(ctx, req.OrderID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	status, ok := commonMapping.OrderStatusesToString[order.Status]
	if !ok {
		return nil, tracing.MarkSpanWithError(ctx, errors.New("mapping error: invalid status"))
	}

	return mapContract(order, status)
}

func mapContract(order *model.Order, status string) (*api.ListOrderResponse, error) {
	return &api.ListOrderResponse{
		User:   order.UserID,
		Status: status,
		Items:  mapping.MapContractOrderItems(order.Items),
	}, nil
}
