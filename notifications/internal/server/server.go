package server

import (
	"context"
	"route256/notifications/internal/domain/model"
	api "route256/notifications/pkg/notifications/v1"
)

type service interface {
	GetHistory(ctx context.Context, userID int64) ([]*model.Notification, error)
}

type Server struct {
	api.UnimplementedNotificationsServer
	service service
}

func New(service service) *Server {
	return &Server{
		service: service,
	}
}
