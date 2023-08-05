// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"
	model "route256/checkout/internal/domain/model"

	mock "github.com/stretchr/testify/mock"
)

// CartRepo is an autogenerated mock type for the cartRepo type
type CartRepo struct {
	mock.Mock
}

// DeleteItem provides a mock function with given fields: ctx, userID, sku
func (_m *CartRepo) DeleteItem(ctx context.Context, userID int64, sku uint32) error {
	ret := _m.Called(ctx, userID, sku)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, uint32) error); ok {
		r0 = rf(ctx, userID, sku)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteItems provides a mock function with given fields: ctx, userID
func (_m *CartRepo) DeleteItems(ctx context.Context, userID int64) error {
	ret := _m.Called(ctx, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetItemCount provides a mock function with given fields: ctx, userID, sku
func (_m *CartRepo) GetItemCount(ctx context.Context, userID int64, sku uint32) (uint16, error) {
	ret := _m.Called(ctx, userID, sku)

	var r0 uint16
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, uint32) (uint16, error)); ok {
		return rf(ctx, userID, sku)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, uint32) uint16); ok {
		r0 = rf(ctx, userID, sku)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, uint32) error); ok {
		r1 = rf(ctx, userID, sku)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetItems provides a mock function with given fields: ctx, userID
func (_m *CartRepo) GetItems(ctx context.Context, userID int64) ([]*model.CartItem, error) {
	ret := _m.Called(ctx, userID)

	var r0 []*model.CartItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]*model.CartItem, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []*model.CartItem); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.CartItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetItemCount provides a mock function with given fields: ctx, userID, sku, count
func (_m *CartRepo) SetItemCount(ctx context.Context, userID int64, sku uint32, count uint16) error {
	ret := _m.Called(ctx, userID, sku, count)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, uint32, uint16) error); ok {
		r0 = rf(ctx, userID, sku, count)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCartRepo creates a new instance of CartRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCartRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *CartRepo {
	mock := &CartRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}