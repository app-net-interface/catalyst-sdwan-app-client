// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	site "github.com/app-net-interface/catalyst-sdwan-app-client/site"
)

// Site is an autogenerated mock type for the Site type
type Site struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, siteID
func (_m *Site) Get(ctx context.Context, siteID string) (*site.Site, error) {
	ret := _m.Called(ctx, siteID)

	var r0 *site.Site
	if rf, ok := ret.Get(0).(func(context.Context, string) *site.Site); ok {
		r0 = rf(ctx, siteID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*site.Site)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, siteID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx
func (_m *Site) List(ctx context.Context) ([]*site.Site, error) {
	ret := _m.Called(ctx)

	var r0 []*site.Site
	if rf, ok := ret.Get(0).(func(context.Context) []*site.Site); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*site.Site)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewSite interface {
	mock.TestingT
	Cleanup(func())
}

// NewSite creates a new instance of Site. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSite(t mockConstructorTestingTNewSite) *Site {
	mock := &Site{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}