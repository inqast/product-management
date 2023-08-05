package service

import (
	"context"
	"route256/loms/internal/domain/model"
)

type stocksRepo interface {
	GetStocks(ctx context.Context, sku uint32) ([]*model.Stock, error)
	SetStock(ctx context.Context, sku uint32, warehouseID int64, count uint64, reserved uint64) error
	DeleteStock(ctx context.Context, sku uint32, warehouseID int64) error
}

type orderRepo interface {
	CreateOrder(ctx context.Context, userID int64) (int64, error)
	AddOrderItem(ctx context.Context, orderID int64, sku uint32, count uint16) error
	GetOrder(ctx context.Context, orderID int64) (*model.Order, error)
	SetOrderStatus(ctx context.Context, orderID int64, status model.OrderStatus) error
}

type transactionManager interface {
	RunRepeatableRead(ctx context.Context, fn func(ctxTx context.Context) error) error
}

type statusSender interface {
	SendMessage(ctx context.Context, userID, orderID int64, status model.OrderStatus) error
}

type Service struct {
	stocksRepo         stocksRepo
	orderRepo          orderRepo
	transactionManager transactionManager
	statusSender       statusSender
}

func New(
	stocksRepo stocksRepo,
	orderRepo orderRepo,
	transactionManager transactionManager,
	statusSender statusSender,
) *Service {
	return &Service{
		stocksRepo:         stocksRepo,
		orderRepo:          orderRepo,
		transactionManager: transactionManager,
		statusSender:       statusSender,
	}
}
