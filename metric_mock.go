// Automatically generated by MockGen. DO NOT EDIT!
// Source: metrics.go

package lile

import (
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// Mock of Monitor interface
type MockMonitor struct {
	ctrl     *gomock.Controller
	recorder *_MockMonitorRecorder
}

// Recorder for MockMonitor (not exported)
type _MockMonitorRecorder struct {
	mock *MockMonitor
}

func NewMockMonitor(ctrl *gomock.Controller) *MockMonitor {
	mock := &MockMonitor{ctrl: ctrl}
	mock.recorder = &_MockMonitorRecorder{mock}
	return mock
}

func (_m *MockMonitor) EXPECT() *_MockMonitorRecorder {
	return _m.recorder
}

func (_m *MockMonitor) InterceptRPC() (grpc.UnaryServerInterceptor, grpc.StreamServerInterceptor) {
	ret := _m.ctrl.Call(_m, "InterceptRPC")
	ret0, _ := ret[0].(grpc.UnaryServerInterceptor)
	ret1, _ := ret[1].(grpc.StreamServerInterceptor)
	return ret0, ret1
}

func (_mr *_MockMonitorRecorder) InterceptRPC() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InterceptRPC")
}

func (_m *MockMonitor) Register(_param0 *Service) error {
	ret := _m.ctrl.Call(_m, "Register", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockMonitorRecorder) Register(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Register", arg0)
}
