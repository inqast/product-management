package sender

import (
	"context"
	"errors"
	"route256/libs/broker"
	"route256/libs/broker/kafka"
	"route256/libs/mw/tracing"
	"route256/loms/internal/domain/model"
	commonMapping "route256/loms/internal/mapping"

	"github.com/opentracing/opentracing-go"
)

type StatusSender struct {
	sender *broker.KafkaSender
}

func New(
	producer *kafka.Producer,
	topic string,
) *StatusSender {
	return &StatusSender{
		sender: broker.NewKafkaSender(
			producer, topic,
		),
	}
}

func (s *StatusSender) SendMessage(ctx context.Context, userID, orderID int64, status model.OrderStatus) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "sender/SendMessage")
	defer span.Finish()

	span.SetTag("order_id", orderID)

	statusStr, ok := commonMapping.OrderStatusesToString[status]
	if !ok {
		return tracing.MarkSpanWithError(ctx, errors.New("mapping error: invalid status"))
	}

	return s.sender.SendMessage(&statusMessage{
		UserID:  userID,
		OrderID: orderID,
		Status:  statusStr,
	})
}
