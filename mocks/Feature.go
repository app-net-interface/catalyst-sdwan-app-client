// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/app-net-interface/catalyst-sdwan-app-client/common"

	feature "github.com/app-net-interface/catalyst-sdwan-app-client/feature"

	mock "github.com/stretchr/testify/mock"
)

// Feature is an autogenerated mock type for the Feature type
type Feature struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, template
func (_m *Feature) Create(ctx context.Context, template *feature.TemplateInput) (string, error) {
	ret := _m.Called(ctx, template)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *feature.TemplateInput) string); ok {
		r0 = rf(ctx, template)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *feature.TemplateInput) error); ok {
		r1 = rf(ctx, template)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, templateID
func (_m *Feature) Get(ctx context.Context, templateID string) (*feature.Template, error) {
	ret := _m.Called(ctx, templateID)

	var r0 *feature.Template
	if rf, ok := ret.Get(0).(func(context.Context, string) *feature.Template); ok {
		r0 = rf(ctx, templateID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*feature.Template)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, templateID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx
func (_m *Feature) List(ctx context.Context) ([]*feature.TemplateInput, error) {
	ret := _m.Called(ctx)

	var r0 []*feature.TemplateInput
	if rf, ok := ret.Get(0).(func(context.Context) []*feature.TemplateInput); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*feature.TemplateInput)
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

// ListByType provides a mock function with given fields: ctx, templateType
func (_m *Feature) ListByType(ctx context.Context, templateType string) ([]*feature.TemplateInput, error) {
	ret := _m.Called(ctx, templateType)

	var r0 []*feature.TemplateInput
	if rf, ok := ret.Get(0).(func(context.Context, string) []*feature.TemplateInput); ok {
		r0 = rf(ctx, templateType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*feature.TemplateInput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, templateType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, templateID, template
func (_m *Feature) Update(ctx context.Context, templateID string, template *feature.Template) (*common.UpdateResponse, error) {
	ret := _m.Called(ctx, templateID, template)

	var r0 *common.UpdateResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, *feature.Template) *common.UpdateResponse); ok {
		r0 = rf(ctx, templateID, template)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.UpdateResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *feature.Template) error); ok {
		r1 = rf(ctx, templateID, template)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFeature interface {
	mock.TestingT
	Cleanup(func())
}

// NewFeature creates a new instance of Feature. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFeature(t mockConstructorTestingTNewFeature) *Feature {
	mock := &Feature{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
