//go:generate mockery --filename loms_mock.go --name loms --exported
//go:generate mockery --filename products_mock.go --name products --exported
//go:generate mockery --filename cart_repo_mock.go --name cartRepo --exported

package service

import (
	"context"
	"route256/checkout/internal/domain/model"
)

type loms interface {
	CreateOrder(ctx context.Context, userId int64, items []*model.CartItem) (int64, error)
	Stocks(ctx context.Context, sku uint32) ([]*model.Stock, error)
}

type products interface {
	GetProduct(ctx context.Context, sku uint32) (*model.Product, error)
	GetProductAsync(ctx context.Context, skus []uint32) (map[uint32]*model.Product, error)
}

type cartRepo interface {
	GetItemCount(ctx context.Context, userID int64, sku uint32) (uint16, error)
	SetItemCount(ctx context.Context, userID int64, sku uint32, count uint16) error
	GetItems(ctx context.Context, userID int64) ([]*model.CartItem, error)
	DeleteItem(ctx context.Context, userID int64, sku uint32) error
	DeleteItems(ctx context.Context, userID int64) error
}

type Service struct {
	lomsClient     loms
	productsClient products
	repo           cartRepo
}

func New(
	lomsClient loms,
	productsClient products,
	repo cartRepo,
) *Service {
	return &Service{
		lomsClient:     lomsClient,
		productsClient: productsClient,
		repo:           repo,
	}
}
