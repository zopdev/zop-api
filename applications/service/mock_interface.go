// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	store "github.com/zopdev/zop-api/applications/store"
	gomock "go.uber.org/mock/gomock"
	gofr "gofr.dev/pkg/gofr"
)

// MockApplicationService is a mock of ApplicationService interface.
type MockApplicationService struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationServiceMockRecorder
}

// MockApplicationServiceMockRecorder is the mock recorder for MockApplicationService.
type MockApplicationServiceMockRecorder struct {
	mock *MockApplicationService
}

// NewMockApplicationService creates a new mock instance.
func NewMockApplicationService(ctrl *gomock.Controller) *MockApplicationService {
	mock := &MockApplicationService{ctrl: ctrl}
	mock.recorder = &MockApplicationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationService) EXPECT() *MockApplicationServiceMockRecorder {
	return m.recorder
}

// AddApplication mocks base method.
func (m *MockApplicationService) AddApplication(ctx *gofr.Context, application *store.Application) (*store.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddApplication", ctx, application)
	ret0, _ := ret[0].(*store.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddApplication indicates an expected call of AddApplication.
func (mr *MockApplicationServiceMockRecorder) AddApplication(ctx, application interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddApplication", reflect.TypeOf((*MockApplicationService)(nil).AddApplication), ctx, application)
}

// FetchAllApplications mocks base method.
func (m *MockApplicationService) FetchAllApplications(ctx *gofr.Context) ([]store.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAllApplications", ctx)
	ret0, _ := ret[0].([]store.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAllApplications indicates an expected call of FetchAllApplications.
func (mr *MockApplicationServiceMockRecorder) FetchAllApplications(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAllApplications", reflect.TypeOf((*MockApplicationService)(nil).FetchAllApplications), ctx)
}
