package server

import (
	"route256/notifications/internal/domain/model"
	api "route256/notifications/pkg/notifications/v1"
)

func mapNotifications(items []*model.Notification) []*api.Notification {
	domainItems := make([]*api.Notification, len(items))

	for i, item := range items {
		domainItems[i] = mapNotification(item)
	}

	return domainItems
}

func mapNotification(item *model.Notification) *api.Notification {
	return &api.Notification{
		OrderId: uint32(item.OrderID),
		Status:  item.Status,
	}
}
