package service

import (
	"context"
	"errors"
	"reflect"
	"route256/checkout/internal/domain/model"
	"route256/checkout/internal/domain/service/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestService_ListCart(t *testing.T) {
	t.Parallel()

	t.Run("default", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		userID := int64(1)

		lomsMock := mocks.NewLoms(t)

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

		productsMock := mocks.NewProducts(t)
		products := map[uint32]*model.Product{
			1: {
				SKU:   1,
				Name:  "1",
				Price: 1,
			},
			2: {
				SKU:   2,
				Name:  "2",
				Price: 2,
			},
			3: {
				SKU:   3,
				Name:  "3",
				Price: 3,
			},
		}
		productsMock.On("GetProductAsync", mock.Anything,
			[]uint32{cartItems[0].SKU, cartItems[1].SKU, cartItems[2].SKU},
		).Return(products, nil)

		s := &Service{
			lomsClient:     lomsMock,
			productsClient: productsMock,
			repo:           repoMock,
		}

		want := &model.Cart{
			Items: []*model.CartItem{
				{
					SKU:   1,
					Count: 1,
					Name:  "1",
					Price: 1,
				},
				{
					SKU:   2,
					Count: 2,
					Name:  "2",
					Price: 2,
				},
				{
					SKU:   3,
					Count: 3,
					Name:  "3",
					Price: 3,
				},
			},
			TotalPrice: 14,
		}

		got, err := s.ListCart(ctx, userID)

		if err != nil {
			t.Errorf("ListCart() error = %v", err)
			return
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ListCart() got = %v, want %v", got, want)
		}
	})

	t.Run("repo error", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		userID := int64(1)

		lomsMock := mocks.NewLoms(t)

		repoMock := mocks.NewCartRepo(t)
		repoMock.On("GetItems", mock.Anything, userID).Return(nil, errors.New("some error"))

		productsMock := mocks.NewProducts(t)

		s := &Service{
			lomsClient:     lomsMock,
			productsClient: productsMock,
			repo:           repoMock,
		}

		_, err := s.ListCart(ctx, userID)

		if err == nil {
			t.Error("ListCart() repo error unhandled")
			return
		}
	})

	t.Run("products error", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		userID := int64(1)

		lomsMock := mocks.NewLoms(t)

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

		productsMock := mocks.NewProducts(t)
		productsMock.On("GetProductAsync", mock.Anything,
			[]uint32{cartItems[0].SKU, cartItems[1].SKU, cartItems[2].SKU},
		).Return(nil, errors.New("some error"))

		s := &Service{
			lomsClient:     lomsMock,
			productsClient: productsMock,
			repo:           repoMock,
		}

		_, err := s.ListCart(ctx, userID)

		if err == nil {
			t.Error("ListCart() products service error unhandled")
			return
		}

	})
}
