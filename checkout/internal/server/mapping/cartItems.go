package mapping

import (
	"route256/checkout/internal/domain/model"
	api "route256/checkout/pkg/checkout/v1"
)

func MapContractCartItems(domainItems []*model.CartItem) []*api.Item {
	contractItems := make([]*api.Item, len(domainItems))

	for i, domainItem := range domainItems {
		contractItems[i] = MapContractCartItem(domainItem)
	}

	return contractItems
}

func MapContractCartItem(domainItem *model.CartItem) *api.Item {
	return &api.Item{
		Sku:   domainItem.SKU,
		Count: uint32(domainItem.Count),
		Name:  domainItem.Name,
		Price: domainItem.Price,
	}
}
