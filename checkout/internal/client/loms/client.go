package loms

import (
	"route256/checkout/internal/pb/loms/v1"
)

type Client struct {
	c loms.LomsClient
}

func New(c loms.LomsClient) *Client {
	return &Client{
		c: c,
	}
}
