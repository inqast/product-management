package service

import (
	"context"
	"fmt"
	"route256/notifications/internal/domain/model"

	"github.com/opentracing/opentracing-go"
)

const notificationFmt = "order id=%d, status changed to %s"

func (s *Service) ProcessStatusChange(ctx context.Context, notification *model.Notification) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "domain/ProcessStatusChange")
	defer span.Finish()

	span.SetTag("order_id", notification.OrderID)

	notificationText := fmt.Sprintf(notificationFmt, notification.OrderID, notification.Status)

	if err := s.repo.Add(ctx, notification.UserID, notification.OrderID, notification.Status); err != nil {
		return fmt.Errorf("error saving history: %w", err)
	}

	return s.sender.Send(ctx, notificationText)
}
