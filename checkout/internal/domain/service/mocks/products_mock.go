// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"
	model "route256/checkout/internal/domain/model"

	mock "github.com/stretchr/testify/mock"
)

// Products is an autogenerated mock type for the products type
type Products struct {
	mock.Mock
}

// GetProduct provides a mock function with given fields: ctx, sku
func (_m *Products) GetProduct(ctx context.Context, sku uint32) (*model.Product, error) {
	ret := _m.Called(ctx, sku)

	var r0 *model.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) (*model.Product, error)); ok {
		return rf(ctx, sku)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32) *model.Product); ok {
		r0 = rf(ctx, sku)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, sku)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductAsync provides a mock function with given fields: ctx, skus
func (_m *Products) GetProductAsync(ctx context.Context, skus []uint32) (map[uint32]*model.Product, error) {
	ret := _m.Called(ctx, skus)

	var r0 map[uint32]*model.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []uint32) (map[uint32]*model.Product, error)); ok {
		return rf(ctx, skus)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []uint32) map[uint32]*model.Product); ok {
		r0 = rf(ctx, skus)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[uint32]*model.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []uint32) error); ok {
		r1 = rf(ctx, skus)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewProducts creates a new instance of Products. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProducts(t interface {
	mock.TestingT
	Cleanup(func())
}) *Products {
	mock := &Products{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
