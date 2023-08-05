package notification

import (
	"route256/notifications/internal/domain/model"
	"route256/notifications/internal/repository/schema"
)

func mapDomainNotifications(items []schema.Notification) []*model.Notification {
	domainItems := make([]*model.Notification, len(items))

	for i, item := range items {
		domainItems[i] = mapDomainNotification(item)
	}

	return domainItems
}

func mapDomainNotification(item schema.Notification) *model.Notification {
	return &model.Notification{
		UserID:  item.UserID,
		OrderID: item.OrderID,
		Status:  item.Status,
	}
}
