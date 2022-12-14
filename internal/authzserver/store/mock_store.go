// Package store is a generated GoMock package.
package store

import (
	reflect "reflect"

	v1 "go_iam/internal/pkg/proto/apiserver/v1"

	gomock "github.com/golang/mock/gomock"
	ladon "github.com/ory/ladon"
)

// MockFactory is a mock of Factory interface.
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory.
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance.
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// Policies mocks base method.
func (m *MockFactory) Policies() PolicyStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Policies")
	ret0, _ := ret[0].(PolicyStore)
	return ret0
}

// Policies indicates an expected call of Policies.
func (mr *MockFactoryMockRecorder) Policies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Policies", reflect.TypeOf((*MockFactory)(nil).Policies))
}

// Secrets mocks base method.
func (m *MockFactory) Secrets() SecretStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Secrets")
	ret0, _ := ret[0].(SecretStore)
	return ret0
}

// Secrets indicates an expected call of Secrets.
func (mr *MockFactoryMockRecorder) Secrets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Secrets", reflect.TypeOf((*MockFactory)(nil).Secrets))
}

// MockSecretStore is a mock of SecretStore interface.
type MockSecretStore struct {
	ctrl     *gomock.Controller
	recorder *MockSecretStoreMockRecorder
}

// MockSecretStoreMockRecorder is the mock recorder for MockSecretStore.
type MockSecretStoreMockRecorder struct {
	mock *MockSecretStore
}

// NewMockSecretStore creates a new mock instance.
func NewMockSecretStore(ctrl *gomock.Controller) *MockSecretStore {
	mock := &MockSecretStore{ctrl: ctrl}
	mock.recorder = &MockSecretStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretStore) EXPECT() *MockSecretStoreMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockSecretStore) List() (map[string]*v1.SecretInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(map[string]*v1.SecretInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockSecretStoreMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSecretStore)(nil).List))
}

// MockPolicyStore is a mock of PolicyStore interface.
type MockPolicyStore struct {
	ctrl     *gomock.Controller
	recorder *MockPolicyStoreMockRecorder
}

// MockPolicyStoreMockRecorder is the mock recorder for MockPolicyStore.
type MockPolicyStoreMockRecorder struct {
	mock *MockPolicyStore
}

// NewMockPolicyStore creates a new mock instance.
func NewMockPolicyStore(ctrl *gomock.Controller) *MockPolicyStore {
	mock := &MockPolicyStore{ctrl: ctrl}
	mock.recorder = &MockPolicyStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPolicyStore) EXPECT() *MockPolicyStoreMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockPolicyStore) List() (map[string][]*ladon.DefaultPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(map[string][]*ladon.DefaultPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPolicyStoreMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPolicyStore)(nil).List))
}
