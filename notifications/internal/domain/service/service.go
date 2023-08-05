//go:generate mockery --filename sender_mock.go --name sender --exported
//go:generate mockery --filename repository_mock.go --name repository --exported

package service

import (
	"context"
	"route256/notifications/internal/domain/model"
)

type sender interface {
	Send(ctx context.Context, notification string) error
}

type repository interface {
	Add(ctx context.Context, userID int64, orderID int64, status string) error
	GetLst(ctx context.Context, userID int64) ([]*model.Notification, error)
}

type Service struct {
	sender sender
	repo   repository
}

func New(snd sender, repo repository) *Service {
	return &Service{
		sender: snd,
		repo:   repo,
	}
}
