package loms

import (
	"context"
	"route256/checkout/internal/domain/model"
	api "route256/checkout/internal/pb/loms/v1"

	"github.com/opentracing/opentracing-go"
)

func (c *Client) CreateOrder(ctx context.Context, userId int64, items []*model.CartItem) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "lomsClient/CreateOrder")
	defer span.Finish()

	span.SetTag("user_id", userId)

	req := &api.CreateOrderRequest{
		User:  userId,
		Items: mapOrderItems(items),
	}

	rs, err := c.c.CreateOrder(ctx, req)
	if err != nil {
		return 0, err
	}

	return rs.GetOrderId(), nil
}

func mapOrderItems(cartItems []*model.CartItem) []*api.Item {
	orderItems := make([]*api.Item, len(cartItems))

	for i, cartItem := range cartItems {
		orderItems[i] = &api.Item{
			Sku:   cartItem.SKU,
			Count: uint32(cartItem.Count),
		}
	}

	return orderItems
}
