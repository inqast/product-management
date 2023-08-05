package product

import (
	"context"
	"route256/checkout/internal/domain/model"
	"route256/libs/workerpool"

	"github.com/opentracing/opentracing-go"
)

func (c *Client) GetProductAsync(ctx context.Context, skus []uint32) (map[uint32]*model.Product, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productClient/GetProducts")
	defer span.Finish()

	wp := workerpool.NewPool(ctx, c.GetProduct, 5)

	for _, sku := range skus {
		wp.AddTask(sku)
	}

	wp.Run()

	products := make(map[uint32]*model.Product, len(skus))

	for result := range wp.GetResults() {
		if result.Err != nil {
			wp.Stop()
			return nil, result.Err
		}

		products[result.Output.SKU] = result.Output
	}

	return products, nil
}
