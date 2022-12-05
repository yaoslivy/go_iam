// Package authorization is a generated GoMock package.
package authorization

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	ladon "github.com/ory/ladon"
)

// MockAuthorizationInterface is a mock of AuthorizationInterface interface.
type MockAuthorizationInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationInterfaceMockRecorder
}

// MockAuthorizationInterfaceMockRecorder is the mock recorder for MockAuthorizationInterface.
type MockAuthorizationInterfaceMockRecorder struct {
	mock *MockAuthorizationInterface
}

// NewMockAuthorizationInterface creates a new mock instance.
func NewMockAuthorizationInterface(ctrl *gomock.Controller) *MockAuthorizationInterface {
	mock := &MockAuthorizationInterface{ctrl: ctrl}
	mock.recorder = &MockAuthorizationInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizationInterface) EXPECT() *MockAuthorizationInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAuthorizationInterface) Create(arg0 *ladon.DefaultPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockAuthorizationInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAuthorizationInterface)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockAuthorizationInterface) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAuthorizationInterfaceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAuthorizationInterface)(nil).Delete), arg0)
}

// DeleteCollection mocks base method.
func (m *MockAuthorizationInterface) DeleteCollection(arg0 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockAuthorizationInterfaceMockRecorder) DeleteCollection(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockAuthorizationInterface)(nil).DeleteCollection), arg0)
}

// Get mocks base method.
func (m *MockAuthorizationInterface) Get(arg0 string) (*ladon.DefaultPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*ladon.DefaultPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAuthorizationInterfaceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAuthorizationInterface)(nil).Get), arg0)
}

// List mocks base method.
func (m *MockAuthorizationInterface) List(arg0 string) ([]*ladon.DefaultPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]*ladon.DefaultPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAuthorizationInterfaceMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAuthorizationInterface)(nil).List), arg0)
}

// LogGrantedAccessRequest mocks base method.
func (m *MockAuthorizationInterface) LogGrantedAccessRequest(arg0 *ladon.Request, arg1, arg2 ladon.Policies) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LogGrantedAccessRequest", arg0, arg1, arg2)
}

// LogGrantedAccessRequest indicates an expected call of LogGrantedAccessRequest.
func (mr *MockAuthorizationInterfaceMockRecorder) LogGrantedAccessRequest(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogGrantedAccessRequest", reflect.TypeOf((*MockAuthorizationInterface)(nil).LogGrantedAccessRequest), arg0, arg1, arg2)
}

// LogRejectedAccessRequest mocks base method.
func (m *MockAuthorizationInterface) LogRejectedAccessRequest(arg0 *ladon.Request, arg1, arg2 ladon.Policies) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LogRejectedAccessRequest", arg0, arg1, arg2)
}

// LogRejectedAccessRequest indicates an expected call of LogRejectedAccessRequest.
func (mr *MockAuthorizationInterfaceMockRecorder) LogRejectedAccessRequest(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogRejectedAccessRequest", reflect.TypeOf((*MockAuthorizationInterface)(nil).LogRejectedAccessRequest), arg0, arg1, arg2)
}

// Update mocks base method.
func (m *MockAuthorizationInterface) Update(arg0 *ladon.DefaultPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockAuthorizationInterfaceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAuthorizationInterface)(nil).Update), arg0)
}
