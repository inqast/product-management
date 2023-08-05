package product

import (
	"route256/checkout/internal/pb/product/v1"

	"golang.org/x/time/rate"
)

type Client struct {
	c            product.ProductServiceClient
	token        string
	limiter      *rate.Limiter
	workersCount int
}

func New(
	c product.ProductServiceClient,
	token string,
	limit,
	workersCount int,
) *Client {
	return &Client{
		c:            c,
		token:        token,
		limiter:      rate.NewLimiter(rate.Limit(limit), limit),
		workersCount: workersCount,
	}
}
