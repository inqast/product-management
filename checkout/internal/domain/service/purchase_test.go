package service

import (
	"context"
	"errors"
	"route256/checkout/internal/domain/model"
	"route256/checkout/internal/domain/service/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestServicePurchase(t *testing.T) {
	t.Parallel()

	t.Run("default", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		userID := int64(1)
		orderID := int64(2)

		productsMock := mocks.NewProducts(t)

		repoMock := mocks.NewCartRepo(t)
		cartItems := []*model.CartItem{
			{
				SKU:   1,
				Count: 1,
			},
			{
				SKU:   2,
				Count: 2,
			},
			{
				SKU:   3,
				Count: 3,
			},
		}
		repoMock.On("GetItems", mock.Anything, userID).Return(cartItems, nil)
		repoMock.On("DeleteItems", mock.Anything, userID).Return(nil)

		lomsMock := mocks.NewLoms(t)
		lomsMock.On("CreateOrder", mock.Anything, userID, cartItems).Return(orderID, nil)

		s := &Service{
			lomsClient:     lomsMock,
			productsClient: productsMock,
			repo:           repoMock,
		}

		want := orderID

		got, err := s.Purchase(ctx, userID)

		if err != nil {
			t.Errorf("Purchase() error = %v", err)
			return
		}

		if got != want {
			t.Errorf("Purchase() got = %v, want %v", got, want)
		}
	})

	t.Run("repo error get", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		userID := int64(1)

		productsMock := mocks.NewProducts(t)

		repoMock := mocks.NewCartRepo(t)
		repoMock.On("GetItems", mock.Anything, userID).Return(nil, errors.New("get error"))

		lomsMock := mocks.NewLoms(t)

		s := &Service{
			lomsClient:     lomsMock,
			productsClient: productsMock,
			repo:           repoMock,
		}

		_, err := s.Purchase(ctx, userID)

		if err == nil {
			t.Error("Purchase() repo get error unhandled")
			return
		}
	})

	t.Run("empty cart error", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		userID := int64(1)

		productsMock := mocks.NewProducts(t)

		repoMock := mocks.NewCartRepo(t)
		repoMock.On("GetItems", mock.Anything, userID).Return([]*model.CartItem{}, nil)

		lomsMock := mocks.NewLoms(t)

		s := &Service{
			lomsClient:     lomsMock,
			productsClient: productsMock,
			repo:           repoMock,
		}

		_, err := s.Purchase(ctx, userID)

		if err == nil {
			t.Error("Purchase() empty cart error unhandled")
			return
		}
	})

	t.Run("loms create order error", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		userID := int64(1)

		productsMock := mocks.NewProducts(t)

		repoMock := mocks.NewCartRepo(t)
		cartItems := []*model.CartItem{
			{
				SKU:   1,
				Count: 1,
			},
			{
				SKU:   2,
				Count: 2,
			},
			{
				SKU:   3,
				Count: 3,
			},
		}
		repoMock.On("GetItems", mock.Anything, userID).Return(cartItems, nil)

		lomsMock := mocks.NewLoms(t)
		lomsMock.On("CreateOrder", mock.Anything, userID, cartItems).Return(int64(0), errors.New("order err"))

		s := &Service{
			lomsClient:     lomsMock,
			productsClient: productsMock,
			repo:           repoMock,
		}

		_, err := s.Purchase(ctx, userID)

		if err == nil {
			t.Error("Purchase() loms create order error unhandled")
			return
		}
	})

	t.Run("repo delete error", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		userID := int64(1)
		orderID := int64(2)

		productsMock := mocks.NewProducts(t)

		repoMock := mocks.NewCartRepo(t)
		cartItems := []*model.CartItem{
			{
				SKU:   1,
				Count: 1,
			},
			{
				SKU:   2,
				Count: 2,
			},
			{
				SKU:   3,
				Count: 3,
			},
		}
		repoMock.On("GetItems", mock.Anything, userID).Return(cartItems, nil)
		repoMock.On("DeleteItems", mock.Anything, userID).Return(errors.New("delete error"))

		lomsMock := mocks.NewLoms(t)
		lomsMock.On("CreateOrder", mock.Anything, userID, cartItems).Return(orderID, nil)

		s := &Service{
			lomsClient:     lomsMock,
			productsClient: productsMock,
			repo:           repoMock,
		}

		want := orderID

		got, err := s.Purchase(ctx, userID)

		if err == nil {
			t.Error("Purchase() repo delete error unhandled")
			return
		}

		if got != want {
			t.Errorf("Purchase() got = %v, want %v", got, want)
		}
	})

}
