// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

import v1beta1 "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"

// RdmaDeviceSpec is an autogenerated mock type for the RdmaDeviceSpec type
type RdmaDeviceSpec struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0
func (_m *RdmaDeviceSpec) Get(_a0 string) []*v1beta1.DeviceSpec {
	ret := _m.Called(_a0)

	var r0 []*v1beta1.DeviceSpec
	if rf, ok := ret.Get(0).(func(string) []*v1beta1.DeviceSpec); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1beta1.DeviceSpec)
		}
	}

	return r0
}

// VerifyRdmaSpec provides a mock function with given fields: _a0
func (_m *RdmaDeviceSpec) VerifyRdmaSpec(_a0 []*v1beta1.DeviceSpec) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*v1beta1.DeviceSpec) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}