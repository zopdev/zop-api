// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -source=interface.go -destination=mock_interface.go -package=provider
//

// Package provider is a generated GoMock package.
package provider

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	gofr "gofr.dev/pkg/gofr"
)

// MockProvider is a mock of Provider interface.
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
	isgomock struct{}
}

// MockProviderMockRecorder is the mock recorder for MockProvider.
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance.
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// ListAllClusters mocks base method.
func (m *MockProvider) ListAllClusters(ctx *gofr.Context, cloudAccount *CloudAccount, credentials any) (*ClusterResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllClusters", ctx, cloudAccount, credentials)
	ret0, _ := ret[0].(*ClusterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllClusters indicates an expected call of ListAllClusters.
func (mr *MockProviderMockRecorder) ListAllClusters(ctx, cloudAccount, credentials any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllClusters", reflect.TypeOf((*MockProvider)(nil).ListAllClusters), ctx, cloudAccount, credentials)
}

// ListDeployments mocks base method.
func (m *MockProvider) ListDeployments(ctx *gofr.Context, cluster *Cluster, cloudAccount *CloudAccount, credentials any, namespace string) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDeployments", ctx, cluster, cloudAccount, credentials, namespace)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeployments indicates an expected call of ListDeployments.
func (mr *MockProviderMockRecorder) ListDeployments(ctx, cluster, cloudAccount, credentials, namespace any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeployments", reflect.TypeOf((*MockProvider)(nil).ListDeployments), ctx, cluster, cloudAccount, credentials, namespace)
}

// ListNamespace mocks base method.
func (m *MockProvider) ListNamespace(ctx *gofr.Context, cluster *Cluster, cloudAccount *CloudAccount, credentials any) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNamespace", ctx, cluster, cloudAccount, credentials)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNamespace indicates an expected call of ListNamespace.
func (mr *MockProviderMockRecorder) ListNamespace(ctx, cluster, cloudAccount, credentials any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNamespace", reflect.TypeOf((*MockProvider)(nil).ListNamespace), ctx, cluster, cloudAccount, credentials)
}

// ListServices mocks base method.
func (m *MockProvider) ListServices(ctx *gofr.Context, cluster *Cluster, cloudAccount *CloudAccount, credentials any, namespace string) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServices", ctx, cluster, cloudAccount, credentials, namespace)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices.
func (mr *MockProviderMockRecorder) ListServices(ctx, cluster, cloudAccount, credentials, namespace any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockProvider)(nil).ListServices), ctx, cluster, cloudAccount, credentials, namespace)
}
