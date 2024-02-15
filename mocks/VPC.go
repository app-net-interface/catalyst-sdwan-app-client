// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	vpc "github.com/app-net-interface/catalyst-sdwan-app-client/vpc"
)

// VPC is an autogenerated mock type for the VPC type
type VPC struct {
	mock.Mock
}

// CreateVPCTag provides a mock function with given fields: ctx, tagName, _a2
func (_m *VPC) CreateVPCTag(ctx context.Context, tagName string, _a2 *vpc.VPC) (string, error) {
	ret := _m.Called(ctx, tagName, _a2)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, *vpc.VPC) string); ok {
		r0 = rf(ctx, tagName, _a2)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *vpc.VPC) error); ok {
		r1 = rf(ctx, tagName, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVPCTag provides a mock function with given fields: ctx, cloudType, tag
func (_m *VPC) DeleteVPCTag(ctx context.Context, cloudType string, tag string) (string, error) {
	ret := _m.Called(ctx, cloudType, tag)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, cloudType, tag)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, cloudType, tag)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, cloudType, vpcID
func (_m *VPC) Get(ctx context.Context, cloudType string, vpcID string) (*vpc.VPC, error) {
	ret := _m.Called(ctx, cloudType, vpcID)

	var r0 *vpc.VPC
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *vpc.VPC); ok {
		r0 = rf(ctx, cloudType, vpcID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vpc.VPC)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, cloudType, vpcID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: ctx, cloudType, vpcName
func (_m *VPC) GetByName(ctx context.Context, cloudType string, vpcName string) (*vpc.VPC, error) {
	ret := _m.Called(ctx, cloudType, vpcName)

	var r0 *vpc.VPC
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *vpc.VPC); ok {
		r0 = rf(ctx, cloudType, vpcName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vpc.VPC)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, cloudType, vpcName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, cloudType
func (_m *VPC) List(ctx context.Context, cloudType string) ([]*vpc.VPC, error) {
	ret := _m.Called(ctx, cloudType)

	var r0 []*vpc.VPC
	if rf, ok := ret.Get(0).(func(context.Context, string) []*vpc.VPC); ok {
		r0 = rf(ctx, cloudType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*vpc.VPC)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, cloudType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAll provides a mock function with given fields: ctx
func (_m *VPC) ListAll(ctx context.Context) ([]*vpc.VPC, error) {
	ret := _m.Called(ctx)

	var r0 []*vpc.VPC
	if rf, ok := ret.Get(0).(func(context.Context) []*vpc.VPC); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*vpc.VPC)
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

// ListWithParameters provides a mock function with given fields: ctx, parameters
func (_m *VPC) ListWithParameters(ctx context.Context, parameters *vpc.ListVPCParameters) ([]*vpc.VPC, error) {
	ret := _m.Called(ctx, parameters)

	var r0 []*vpc.VPC
	if rf, ok := ret.Get(0).(func(context.Context, *vpc.ListVPCParameters) []*vpc.VPC); ok {
		r0 = rf(ctx, parameters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*vpc.VPC)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *vpc.ListVPCParameters) error); ok {
		r1 = rf(ctx, parameters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListWithTag provides a mock function with given fields: ctx, cloudType
func (_m *VPC) ListWithTag(ctx context.Context, cloudType string) ([]*vpc.VPC, error) {
	ret := _m.Called(ctx, cloudType)

	var r0 []*vpc.VPC
	if rf, ok := ret.Get(0).(func(context.Context, string) []*vpc.VPC); ok {
		r0 = rf(ctx, cloudType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*vpc.VPC)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, cloudType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListWithTagWithParameters provides a mock function with given fields: ctx, parameters
func (_m *VPC) ListWithTagWithParameters(ctx context.Context, parameters *vpc.ListVPCTagParameters) ([]*vpc.VPC, error) {
	ret := _m.Called(ctx, parameters)

	var r0 []*vpc.VPC
	if rf, ok := ret.Get(0).(func(context.Context, *vpc.ListVPCTagParameters) []*vpc.VPC); ok {
		r0 = rf(ctx, parameters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*vpc.VPC)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *vpc.ListVPCTagParameters) error); ok {
		r1 = rf(ctx, parameters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewVPC interface {
	mock.TestingT
	Cleanup(func())
}

// NewVPC creates a new instance of VPC. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewVPC(t mockConstructorTestingTNewVPC) *VPC {
	mock := &VPC{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}