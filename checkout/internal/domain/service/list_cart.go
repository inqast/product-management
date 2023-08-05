package service

import (
	"context"
	"fmt"
	"route256/checkout/internal/domain/model"

	"github.com/opentracing/opentracing-go"
)

func (s *Service) ListCart(
	ctx context.Context,
	userID int64,
) (*model.Cart, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "domain/ListCart")
	defer span.Finish()

	span.SetTag("user_id", userID)

	cartItems, err := s.repo.GetItems(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("error getting cart items: %w", err)
	}

	skus := make([]uint32, 0, len(cartItems))
	for _, cartItem := range cartItems {
		skus = append(skus, cartItem.SKU)
	}

	products, err := s.productsClient.GetProductAsync(ctx, skus)
	if err != nil {
		return nil, fmt.Errorf("error getting products info: %w", err)
	}

	var totalPrice uint32
	for _, cartItem := range cartItems {
		product := products[cartItem.SKU]

		cartItem.Name = product.Name
		cartItem.Price = product.Price

		totalPrice += cartItem.Price * uint32(cartItem.Count)
	}

	return &model.Cart{
		Items:      cartItems,
		TotalPrice: totalPrice,
	}, nil
}
