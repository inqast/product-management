package order

import (
	"route256/loms/internal/domain/model"
	"route256/loms/internal/repository/schema"
)

func mapDomainOrder(order schema.Order, items []schema.Item) *model.Order {
	return &model.Order{
		ID:     order.ID,
		UserID: order.UserID,
		Status: model.OrderStatus(order.Status),
		Items:  mapDomainOrderItems(items),
	}
}

func mapDomainOrderItems(items []schema.Item) []*model.OrderItem {
	domainItems := make([]*model.OrderItem, len(items))

	for i, item := range items {
		domainItems[i] = mapDomainOrderItem(item)
	}

	return domainItems
}

func mapDomainOrderItem(item schema.Item) *model.OrderItem {
	return &model.OrderItem{
		SKU:   uint32(item.SKU),
		Count: uint16(item.Count),
	}
}
