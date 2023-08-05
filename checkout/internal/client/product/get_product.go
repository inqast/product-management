package product

import (
	"context"
	"route256/checkout/internal/domain/model"
	api "route256/checkout/internal/pb/product/v1"

	"github.com/opentracing/opentracing-go"
)

func (c *Client) GetProduct(ctx context.Context, sku uint32) (*model.Product, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productClient/GetProduct")
	defer span.Finish()

	span.SetTag("sku", sku)

	err := c.limiter.Wait(ctx)
	if err != nil {
		return nil, err
	}

	req := &api.GetProductRequest{
		Token: c.token,
		Sku:   sku,
	}

	rs, err := c.c.GetProduct(ctx, req)
	if err != nil {
		return nil, err
	}

	return &model.Product{
		SKU:   sku,
		Name:  rs.GetName(),
		Price: rs.GetPrice(),
	}, nil
}
