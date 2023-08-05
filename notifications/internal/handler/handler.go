package handler

import (
	"context"
	"encoding/json"
	log "route256/libs/logger"
	"route256/notifications/internal/domain/model"

	"github.com/Shopify/sarama"
	"github.com/opentracing/opentracing-go"
)

type service interface {
	ProcessStatusChange(ctx context.Context, order *model.Notification) error
}

type Handler struct {
	service service
}

func New(srv service) *Handler {
	return &Handler{
		service: srv,
	}
}

func (h *Handler) HandleOrderStatus(ctx context.Context, message *sarama.ConsumerMessage) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "handler/HandleOrderStatus")
	defer span.Finish()

	statusMsg := &statusMessage{}
	err := json.Unmarshal(message.Value, statusMsg)
	if err != nil {
		log.Error("failed to unmarshall message")
		return
	}

	notification := mapDomainNotification(statusMsg)

	err = h.service.ProcessStatusChange(ctx, notification)
	if err != nil {
		log.Error("failed to process notification status for id ", statusMsg.OrderID, err.Error())
	}
}
