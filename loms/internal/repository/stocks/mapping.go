package stocks

import (
	"route256/loms/internal/domain/model"
	"route256/loms/internal/repository/schema"
)

func mapDomainStocks(items []schema.StockItem) []*model.Stock {
	domainStocks := make([]*model.Stock, len(items))

	for i, item := range items {
		domainStocks[i] = mapDomainStock(item)
	}

	return domainStocks
}

func mapDomainStock(item schema.StockItem) *model.Stock {
	return &model.Stock{
		WarehouseID: item.WarehouseID,
		Count:       uint64(item.Count),
		Reserved:    uint64(item.Reserved),
	}
}
