package server

import (
	"context"
	api "route256/notifications/pkg/notifications/v1"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetHistory(ctx context.Context, req *api.GetHistoryRequest) (*api.GetHistoryResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server/GetHistory")
	defer span.Finish()

	span.SetTag("user_id", req.User)

	notifications, err := s.service.GetHistory(ctx, req.User)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &api.GetHistoryResponse{
		Notifications: mapNotifications(notifications),
	}, nil
}
