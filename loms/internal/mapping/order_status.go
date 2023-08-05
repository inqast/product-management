package mapping

import "route256/loms/internal/domain/model"

var OrderStatusesToString = map[model.OrderStatus]string{
	model.OrderNew:             "new",
	model.OrderAwaitingPayment: "awaiting payment",
	model.OrderFailed:          "failed",
	model.OrderPayed:           "payed",
	model.OrderCancelled:       "cancelled",
}

var OrderStatusesFromString = map[string]model.OrderStatus{
	"new":              model.OrderNew,
	"awaiting payment": model.OrderAwaitingPayment,
	"failed":           model.OrderFailed,
	"payed":            model.OrderPayed,
	"cancelled":        model.OrderCancelled,
}
