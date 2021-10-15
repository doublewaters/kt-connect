// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/kt/exec/sshuttle/types.go

// Package sshuttle is a generated GoMock package.
package sshuttle

import (
	exec "os/exec"
	reflect "reflect"

	options "github.com/alibaba/kt-connect/pkg/kt/options"
	gomock "github.com/golang/mock/gomock"
)

// MockCliInterface is a mock of CliInterface interface.
type MockCliInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCliInterfaceMockRecorder
}

// MockCliInterfaceMockRecorder is the mock recorder for MockCliInterface.
type MockCliInterfaceMockRecorder struct {
	mock *MockCliInterface
}

// NewMockCliInterface creates a new mock instance.
func NewMockCliInterface(ctrl *gomock.Controller) *MockCliInterface {
	mock := &MockCliInterface{ctrl: ctrl}
	mock.recorder = &MockCliInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCliInterface) EXPECT() *MockCliInterfaceMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockCliInterface) Connect(opt *options.ConnectOptions, req *SSHVPNRequest) *exec.Cmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", opt, req)
	ret0, _ := ret[0].(*exec.Cmd)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockCliInterfaceMockRecorder) Connect(opt, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockCliInterface)(nil).Connect), opt, req)
}

// Install mocks base method.
func (m *MockCliInterface) Install() *exec.Cmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install")
	ret0, _ := ret[0].(*exec.Cmd)
	return ret0
}

// Install indicates an expected call of Install.
func (mr *MockCliInterfaceMockRecorder) Install() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockCliInterface)(nil).Install))
}

// Version mocks base method.
func (m *MockCliInterface) Version() *exec.Cmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(*exec.Cmd)
	return ret0
}

// Version indicates an expected call of Version.
func (mr *MockCliInterfaceMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockCliInterface)(nil).Version))
}
