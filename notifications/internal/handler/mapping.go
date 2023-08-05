package handler

import "route256/notifications/internal/domain/model"

func mapDomainNotification(statusMsg *statusMessage) *model.Notification {
	return &model.Notification{
		UserID:  statusMsg.UserID,
		OrderID: statusMsg.OrderID,
		Status:  statusMsg.Status,
	}
}
