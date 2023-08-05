package mapping

import (
	"route256/loms/internal/domain/model"
	api "route256/loms/pkg/loms/v1"
)

func MapContractStocks(domainStocks []*model.Stock) []*api.Stock {
	contractItems := make([]*api.Stock, len(domainStocks))

	for i, domainStock := range domainStocks {
		contractItems[i] = MapContractStock(domainStock)
	}

	return contractItems
}

func MapContractStock(domainStock *model.Stock) *api.Stock {
	return &api.Stock{
		WarehouseID: domainStock.WarehouseID,
		Count:       domainStock.Count,
	}
}
