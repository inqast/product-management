package cart

import (
	"route256/checkout/internal/domain/model"
	"route256/checkout/internal/repository/schema"
)

func mapDomainModels(items []schema.Item) []*model.CartItem {
	domainItems := make([]*model.CartItem, len(items))

	for i, item := range items {
		domainItems[i] = mapDomainModel(item)
	}

	return domainItems
}

func mapDomainModel(item schema.Item) *model.CartItem {
	return &model.CartItem{
		SKU:   uint32(item.SKU),
		Count: uint16(item.Count),
	}
}
