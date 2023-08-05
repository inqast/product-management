package service

import (
	"context"
	"fmt"
	"route256/notifications/internal/domain/model"

	"github.com/opentracing/opentracing-go"
)

func (s *Service) GetHistory(ctx context.Context, userID int64) ([]*model.Notification, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "domain/GetHistory")
	defer span.Finish()

	notifications, err := s.repo.GetLst(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("error getting history: %w", err)
	}

	return notifications, nil
}
