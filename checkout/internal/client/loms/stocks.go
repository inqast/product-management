package loms

import (
	"context"
	"route256/checkout/internal/domain/model"
	"route256/checkout/internal/pb/loms/v1"

	"github.com/opentracing/opentracing-go"
)

func (c *Client) Stocks(ctx context.Context, sku uint32) ([]*model.Stock, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "lomsClient/Stocks")
	defer span.Finish()

	span.SetTag("sku", sku)

	req := &loms.StocksRequest{
		Sku: sku,
	}

	rs, err := c.c.Stocks(ctx, req)
	if err != nil {
		return nil, err
	}

	return mapDomainStocks(rs.GetStocks()), nil
}

func mapDomainStocks(contractStocks []*loms.Stock) []*model.Stock {
	domainStocks := make([]*model.Stock, len(contractStocks))

	for i, contractStock := range contractStocks {
		domainStocks[i] = mapDomainStock(contractStock)
	}

	return domainStocks
}

func mapDomainStock(contractStock *loms.Stock) *model.Stock {
	return &model.Stock{
		WarehouseID: contractStock.GetWarehouseId(),
		Count:       contractStock.GetCount(),
	}
}
