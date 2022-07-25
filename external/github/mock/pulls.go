// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ksrof/trello-action/external/github (interfaces: Pulls)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	github "github.com/ksrof/trello-action/external/github"
	utils "github.com/ksrof/trello-action/utils"
)

// MockPulls is a mock of Pulls interface.
type MockPulls struct {
	ctrl     *gomock.Controller
	recorder *MockPullsMockRecorder
}

// MockPullsMockRecorder is the mock recorder for MockPulls.
type MockPullsMockRecorder struct {
	mock *MockPulls
}

// NewMockPulls creates a new mock instance.
func NewMockPulls(ctrl *gomock.Controller) *MockPulls {
	mock := &MockPulls{ctrl: ctrl}
	mock.recorder = &MockPullsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPulls) EXPECT() *MockPullsMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockPulls) Get(arg0 context.Context, arg1 []utils.Field) (*github.PullsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*github.PullsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPullsMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPulls)(nil).Get), arg0, arg1)
}
