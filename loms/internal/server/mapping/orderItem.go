package mapping

import (
	"route256/loms/internal/domain/model"
	api "route256/loms/pkg/loms/v1"
)

func MapDomainOrderItems(contractItems []*api.Item) []*model.OrderItem {
	domainItems := make([]*model.OrderItem, len(contractItems))

	for i, contractItem := range contractItems {
		domainItems[i] = MapDomainOrderItem(contractItem)
	}

	return domainItems
}

func MapDomainOrderItem(contractItem *api.Item) *model.OrderItem {
	return &model.OrderItem{
		SKU:   contractItem.GetSku(),
		Count: uint16(contractItem.GetCount()),
	}
}

func MapContractOrderItems(domainItems []*model.OrderItem) []*api.Item {
	contractItems := make([]*api.Item, len(domainItems))

	for i, domainItem := range domainItems {
		contractItems[i] = MapContractOrderItem(domainItem)
	}

	return contractItems
}

func MapContractOrderItem(domainItem *model.OrderItem) *api.Item {
	return &api.Item{
		Sku:   domainItem.SKU,
		Count: uint32(domainItem.Count),
	}
}
