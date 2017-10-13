// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nebulasio/go-nebulas/rpc/pb (interfaces: APIServiceClient)

// Package mock_pb is a generated GoMock package.
package mock_pb

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	pb "github.com/nebulasio/go-nebulas/rpc/pb"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockAPIServiceClient is a mock of APIServiceClient interface
type MockAPIServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockAPIServiceClientMockRecorder
}

// MockAPIServiceClientMockRecorder is the mock recorder for MockAPIServiceClient
type MockAPIServiceClientMockRecorder struct {
	mock *MockAPIServiceClient
}

// NewMockAPIServiceClient creates a new mock instance
func NewMockAPIServiceClient(ctrl *gomock.Controller) *MockAPIServiceClient {
	mock := &MockAPIServiceClient{ctrl: ctrl}
	mock.recorder = &MockAPIServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPIServiceClient) EXPECT() *MockAPIServiceClientMockRecorder {
	return m.recorder
}

// GetAccountState mocks base method
func (m *MockAPIServiceClient) GetAccountState(arg0 context.Context, arg1 *pb.GetAccountStateRequest, arg2 ...grpc.CallOption) (*pb.GetAccountStateResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccountState", varargs...)
	ret0, _ := ret[0].(*pb.GetAccountStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountState indicates an expected call of GetAccountState
func (mr *MockAPIServiceClientMockRecorder) GetAccountState(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountState", reflect.TypeOf((*MockAPIServiceClient)(nil).GetAccountState), varargs...)
}

// SendTransaction mocks base method
func (m *MockAPIServiceClient) SendTransaction(arg0 context.Context, arg1 *pb.SendTransactionRequest, arg2 ...grpc.CallOption) (*pb.SendTransactionResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendTransaction", varargs...)
	ret0, _ := ret[0].(*pb.SendTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTransaction indicates an expected call of SendTransaction
func (mr *MockAPIServiceClientMockRecorder) SendTransaction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransaction", reflect.TypeOf((*MockAPIServiceClient)(nil).SendTransaction), varargs...)
}
