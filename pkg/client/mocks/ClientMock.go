// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/backplane-api/pkg/client (interfaces: ClientInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	Openapi "github.com/openshift/backplane-api/pkg/client"
)

// MockClientInterface is a mock of ClientInterface interface.
type MockClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClientInterfaceMockRecorder
}

// MockClientInterfaceMockRecorder is the mock recorder for MockClientInterface.
type MockClientInterfaceMockRecorder struct {
	mock *MockClientInterface
}

// NewMockClientInterface creates a new mock instance.
func NewMockClientInterface(ctrl *gomock.Controller) *MockClientInterface {
	mock := &MockClientInterface{ctrl: ctrl}
	mock.recorder = &MockClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientInterface) EXPECT() *MockClientInterfaceMockRecorder {
	return m.recorder
}

// CreateJob mocks base method.
func (m *MockClientInterface) CreateJob(arg0 context.Context, arg1 string, arg2 Openapi.CreateJob, arg3 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateJob", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJob indicates an expected call of CreateJob.
func (mr *MockClientInterfaceMockRecorder) CreateJob(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJob", reflect.TypeOf((*MockClientInterface)(nil).CreateJob), varargs...)
}

// CreateJobWithBody mocks base method.
func (m *MockClientInterface) CreateJobWithBody(arg0 context.Context, arg1, arg2 string, arg3 io.Reader, arg4 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateJobWithBody", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJobWithBody indicates an expected call of CreateJobWithBody.
func (mr *MockClientInterfaceMockRecorder) CreateJobWithBody(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJobWithBody", reflect.TypeOf((*MockClientInterface)(nil).CreateJobWithBody), varargs...)
}

// CreateRemediation mocks base method.
func (m *MockClientInterface) CreateRemediation(arg0 context.Context, arg1 string, arg2 *Openapi.CreateRemediationParams, arg3 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateRemediation", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRemediation indicates an expected call of CreateRemediation.
func (mr *MockClientInterfaceMockRecorder) CreateRemediation(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRemediation", reflect.TypeOf((*MockClientInterface)(nil).CreateRemediation), varargs...)
}

// CreateTestScriptRun mocks base method.
func (m *MockClientInterface) CreateTestScriptRun(arg0 context.Context, arg1 string, arg2 Openapi.CreateTestJob, arg3 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTestScriptRun", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTestScriptRun indicates an expected call of CreateTestScriptRun.
func (mr *MockClientInterfaceMockRecorder) CreateTestScriptRun(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTestScriptRun", reflect.TypeOf((*MockClientInterface)(nil).CreateTestScriptRun), varargs...)
}

// CreateTestScriptRunWithBody mocks base method.
func (m *MockClientInterface) CreateTestScriptRunWithBody(arg0 context.Context, arg1, arg2 string, arg3 io.Reader, arg4 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTestScriptRunWithBody", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTestScriptRunWithBody indicates an expected call of CreateTestScriptRunWithBody.
func (mr *MockClientInterfaceMockRecorder) CreateTestScriptRunWithBody(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTestScriptRunWithBody", reflect.TypeOf((*MockClientInterface)(nil).CreateTestScriptRunWithBody), varargs...)
}

// DeleteBackplaneClusterClusterId mocks base method.
func (m *MockClientInterface) DeleteBackplaneClusterClusterId(arg0 context.Context, arg1 string, arg2 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteBackplaneClusterClusterId", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBackplaneClusterClusterId indicates an expected call of DeleteBackplaneClusterClusterId.
func (mr *MockClientInterfaceMockRecorder) DeleteBackplaneClusterClusterId(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBackplaneClusterClusterId", reflect.TypeOf((*MockClientInterface)(nil).DeleteBackplaneClusterClusterId), varargs...)
}

// DeleteBackplaneRemediateClusterIdRemediation mocks base method.
func (m *MockClientInterface) DeleteBackplaneRemediateClusterIdRemediation(arg0 context.Context, arg1, arg2 string, arg3 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteBackplaneRemediateClusterIdRemediation", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBackplaneRemediateClusterIdRemediation indicates an expected call of DeleteBackplaneRemediateClusterIdRemediation.
func (mr *MockClientInterfaceMockRecorder) DeleteBackplaneRemediateClusterIdRemediation(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBackplaneRemediateClusterIdRemediation", reflect.TypeOf((*MockClientInterface)(nil).DeleteBackplaneRemediateClusterIdRemediation), varargs...)
}

// DeleteJob mocks base method.
func (m *MockClientInterface) DeleteJob(arg0 context.Context, arg1, arg2 string, arg3 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteJob", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteJob indicates an expected call of DeleteJob.
func (mr *MockClientInterfaceMockRecorder) DeleteJob(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteJob", reflect.TypeOf((*MockClientInterface)(nil).DeleteJob), varargs...)
}

// DeleteRemediation mocks base method.
func (m *MockClientInterface) DeleteRemediation(arg0 context.Context, arg1 string, arg2 *Openapi.DeleteRemediationParams, arg3 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteRemediation", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRemediation indicates an expected call of DeleteRemediation.
func (mr *MockClientInterfaceMockRecorder) DeleteRemediation(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRemediation", reflect.TypeOf((*MockClientInterface)(nil).DeleteRemediation), varargs...)
}

// GetAllJobs mocks base method.
func (m *MockClientInterface) GetAllJobs(arg0 context.Context, arg1 string, arg2 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllJobs", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllJobs indicates an expected call of GetAllJobs.
func (mr *MockClientInterfaceMockRecorder) GetAllJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllJobs", reflect.TypeOf((*MockClientInterface)(nil).GetAllJobs), varargs...)
}

// GetAssumeRoleSequence mocks base method.
func (m *MockClientInterface) GetAssumeRoleSequence(arg0 context.Context, arg1 string, arg2 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAssumeRoleSequence", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssumeRoleSequence indicates an expected call of GetAssumeRoleSequence.
func (mr *MockClientInterfaceMockRecorder) GetAssumeRoleSequence(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssumeRoleSequence", reflect.TypeOf((*MockClientInterface)(nil).GetAssumeRoleSequence), varargs...)
}

// GetBackplaneClusterClusterId mocks base method.
func (m *MockClientInterface) GetBackplaneClusterClusterId(arg0 context.Context, arg1 string, arg2 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBackplaneClusterClusterId", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackplaneClusterClusterId indicates an expected call of GetBackplaneClusterClusterId.
func (mr *MockClientInterfaceMockRecorder) GetBackplaneClusterClusterId(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackplaneClusterClusterId", reflect.TypeOf((*MockClientInterface)(nil).GetBackplaneClusterClusterId), varargs...)
}

// GetBackplaneRemediateClusterIdRemediation mocks base method.
func (m *MockClientInterface) GetBackplaneRemediateClusterIdRemediation(arg0 context.Context, arg1, arg2 string, arg3 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBackplaneRemediateClusterIdRemediation", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackplaneRemediateClusterIdRemediation indicates an expected call of GetBackplaneRemediateClusterIdRemediation.
func (mr *MockClientInterfaceMockRecorder) GetBackplaneRemediateClusterIdRemediation(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackplaneRemediateClusterIdRemediation", reflect.TypeOf((*MockClientInterface)(nil).GetBackplaneRemediateClusterIdRemediation), varargs...)
}

// GetCloudConsole mocks base method.
func (m *MockClientInterface) GetCloudConsole(arg0 context.Context, arg1 string, arg2 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCloudConsole", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCloudConsole indicates an expected call of GetCloudConsole.
func (mr *MockClientInterfaceMockRecorder) GetCloudConsole(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCloudConsole", reflect.TypeOf((*MockClientInterface)(nil).GetCloudConsole), varargs...)
}

// GetCloudCredentials mocks base method.
func (m *MockClientInterface) GetCloudCredentials(arg0 context.Context, arg1 string, arg2 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCloudCredentials", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCloudCredentials indicates an expected call of GetCloudCredentials.
func (mr *MockClientInterfaceMockRecorder) GetCloudCredentials(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCloudCredentials", reflect.TypeOf((*MockClientInterface)(nil).GetCloudCredentials), varargs...)
}

// GetJobLogs mocks base method.
func (m *MockClientInterface) GetJobLogs(arg0 context.Context, arg1, arg2 string, arg3 *Openapi.GetJobLogsParams, arg4 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetJobLogs", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobLogs indicates an expected call of GetJobLogs.
func (mr *MockClientInterfaceMockRecorder) GetJobLogs(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobLogs", reflect.TypeOf((*MockClientInterface)(nil).GetJobLogs), varargs...)
}

// GetRun mocks base method.
func (m *MockClientInterface) GetRun(arg0 context.Context, arg1, arg2 string, arg3 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRun", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRun indicates an expected call of GetRun.
func (mr *MockClientInterfaceMockRecorder) GetRun(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRun", reflect.TypeOf((*MockClientInterface)(nil).GetRun), varargs...)
}

// GetScriptsByCluster mocks base method.
func (m *MockClientInterface) GetScriptsByCluster(arg0 context.Context, arg1 string, arg2 *Openapi.GetScriptsByClusterParams, arg3 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetScriptsByCluster", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScriptsByCluster indicates an expected call of GetScriptsByCluster.
func (mr *MockClientInterfaceMockRecorder) GetScriptsByCluster(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScriptsByCluster", reflect.TypeOf((*MockClientInterface)(nil).GetScriptsByCluster), varargs...)
}

// GetTestScriptRun mocks base method.
func (m *MockClientInterface) GetTestScriptRun(arg0 context.Context, arg1, arg2 string, arg3 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTestScriptRun", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestScriptRun indicates an expected call of GetTestScriptRun.
func (mr *MockClientInterfaceMockRecorder) GetTestScriptRun(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestScriptRun", reflect.TypeOf((*MockClientInterface)(nil).GetTestScriptRun), varargs...)
}

// GetTestScriptRunLogs mocks base method.
func (m *MockClientInterface) GetTestScriptRunLogs(arg0 context.Context, arg1, arg2 string, arg3 *Openapi.GetTestScriptRunLogsParams, arg4 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTestScriptRunLogs", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestScriptRunLogs indicates an expected call of GetTestScriptRunLogs.
func (mr *MockClientInterfaceMockRecorder) GetTestScriptRunLogs(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestScriptRunLogs", reflect.TypeOf((*MockClientInterface)(nil).GetTestScriptRunLogs), varargs...)
}

// HeadBackplaneClusterClusterId mocks base method.
func (m *MockClientInterface) HeadBackplaneClusterClusterId(arg0 context.Context, arg1 string, arg2 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HeadBackplaneClusterClusterId", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeadBackplaneClusterClusterId indicates an expected call of HeadBackplaneClusterClusterId.
func (mr *MockClientInterfaceMockRecorder) HeadBackplaneClusterClusterId(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadBackplaneClusterClusterId", reflect.TypeOf((*MockClientInterface)(nil).HeadBackplaneClusterClusterId), varargs...)
}

// HeadBackplaneRemediateClusterIdRemediation mocks base method.
func (m *MockClientInterface) HeadBackplaneRemediateClusterIdRemediation(arg0 context.Context, arg1, arg2 string, arg3 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HeadBackplaneRemediateClusterIdRemediation", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeadBackplaneRemediateClusterIdRemediation indicates an expected call of HeadBackplaneRemediateClusterIdRemediation.
func (mr *MockClientInterfaceMockRecorder) HeadBackplaneRemediateClusterIdRemediation(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadBackplaneRemediateClusterIdRemediation", reflect.TypeOf((*MockClientInterface)(nil).HeadBackplaneRemediateClusterIdRemediation), varargs...)
}

// LoginCluster mocks base method.
func (m *MockClientInterface) LoginCluster(arg0 context.Context, arg1 string, arg2 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LoginCluster", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginCluster indicates an expected call of LoginCluster.
func (mr *MockClientInterfaceMockRecorder) LoginCluster(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginCluster", reflect.TypeOf((*MockClientInterface)(nil).LoginCluster), varargs...)
}

// OptionsBackplaneClusterClusterId mocks base method.
func (m *MockClientInterface) OptionsBackplaneClusterClusterId(arg0 context.Context, arg1 string, arg2 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OptionsBackplaneClusterClusterId", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OptionsBackplaneClusterClusterId indicates an expected call of OptionsBackplaneClusterClusterId.
func (mr *MockClientInterfaceMockRecorder) OptionsBackplaneClusterClusterId(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OptionsBackplaneClusterClusterId", reflect.TypeOf((*MockClientInterface)(nil).OptionsBackplaneClusterClusterId), varargs...)
}

// OptionsBackplaneRemediateClusterIdRemediation mocks base method.
func (m *MockClientInterface) OptionsBackplaneRemediateClusterIdRemediation(arg0 context.Context, arg1, arg2 string, arg3 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OptionsBackplaneRemediateClusterIdRemediation", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OptionsBackplaneRemediateClusterIdRemediation indicates an expected call of OptionsBackplaneRemediateClusterIdRemediation.
func (mr *MockClientInterfaceMockRecorder) OptionsBackplaneRemediateClusterIdRemediation(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OptionsBackplaneRemediateClusterIdRemediation", reflect.TypeOf((*MockClientInterface)(nil).OptionsBackplaneRemediateClusterIdRemediation), varargs...)
}

// PatchBackplaneClusterClusterId mocks base method.
func (m *MockClientInterface) PatchBackplaneClusterClusterId(arg0 context.Context, arg1 string, arg2 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchBackplaneClusterClusterId", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PatchBackplaneClusterClusterId indicates an expected call of PatchBackplaneClusterClusterId.
func (mr *MockClientInterfaceMockRecorder) PatchBackplaneClusterClusterId(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchBackplaneClusterClusterId", reflect.TypeOf((*MockClientInterface)(nil).PatchBackplaneClusterClusterId), varargs...)
}

// PatchBackplaneRemediateClusterIdRemediation mocks base method.
func (m *MockClientInterface) PatchBackplaneRemediateClusterIdRemediation(arg0 context.Context, arg1, arg2 string, arg3 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchBackplaneRemediateClusterIdRemediation", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PatchBackplaneRemediateClusterIdRemediation indicates an expected call of PatchBackplaneRemediateClusterIdRemediation.
func (mr *MockClientInterfaceMockRecorder) PatchBackplaneRemediateClusterIdRemediation(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchBackplaneRemediateClusterIdRemediation", reflect.TypeOf((*MockClientInterface)(nil).PatchBackplaneRemediateClusterIdRemediation), varargs...)
}

// PostBackplaneClusterClusterId mocks base method.
func (m *MockClientInterface) PostBackplaneClusterClusterId(arg0 context.Context, arg1 string, arg2 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostBackplaneClusterClusterId", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostBackplaneClusterClusterId indicates an expected call of PostBackplaneClusterClusterId.
func (mr *MockClientInterfaceMockRecorder) PostBackplaneClusterClusterId(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostBackplaneClusterClusterId", reflect.TypeOf((*MockClientInterface)(nil).PostBackplaneClusterClusterId), varargs...)
}

// PostBackplaneRemediateClusterIdRemediation mocks base method.
func (m *MockClientInterface) PostBackplaneRemediateClusterIdRemediation(arg0 context.Context, arg1, arg2 string, arg3 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostBackplaneRemediateClusterIdRemediation", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostBackplaneRemediateClusterIdRemediation indicates an expected call of PostBackplaneRemediateClusterIdRemediation.
func (mr *MockClientInterfaceMockRecorder) PostBackplaneRemediateClusterIdRemediation(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostBackplaneRemediateClusterIdRemediation", reflect.TypeOf((*MockClientInterface)(nil).PostBackplaneRemediateClusterIdRemediation), varargs...)
}

// PutBackplaneClusterClusterId mocks base method.
func (m *MockClientInterface) PutBackplaneClusterClusterId(arg0 context.Context, arg1 string, arg2 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutBackplaneClusterClusterId", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutBackplaneClusterClusterId indicates an expected call of PutBackplaneClusterClusterId.
func (mr *MockClientInterfaceMockRecorder) PutBackplaneClusterClusterId(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutBackplaneClusterClusterId", reflect.TypeOf((*MockClientInterface)(nil).PutBackplaneClusterClusterId), varargs...)
}

// PutBackplaneRemediateClusterIdRemediation mocks base method.
func (m *MockClientInterface) PutBackplaneRemediateClusterIdRemediation(arg0 context.Context, arg1, arg2 string, arg3 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutBackplaneRemediateClusterIdRemediation", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutBackplaneRemediateClusterIdRemediation indicates an expected call of PutBackplaneRemediateClusterIdRemediation.
func (mr *MockClientInterfaceMockRecorder) PutBackplaneRemediateClusterIdRemediation(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutBackplaneRemediateClusterIdRemediation", reflect.TypeOf((*MockClientInterface)(nil).PutBackplaneRemediateClusterIdRemediation), varargs...)
}

// TraceBackplaneClusterClusterId mocks base method.
func (m *MockClientInterface) TraceBackplaneClusterClusterId(arg0 context.Context, arg1 string, arg2 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TraceBackplaneClusterClusterId", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TraceBackplaneClusterClusterId indicates an expected call of TraceBackplaneClusterClusterId.
func (mr *MockClientInterfaceMockRecorder) TraceBackplaneClusterClusterId(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TraceBackplaneClusterClusterId", reflect.TypeOf((*MockClientInterface)(nil).TraceBackplaneClusterClusterId), varargs...)
}

// TraceBackplaneRemediateClusterIdRemediation mocks base method.
func (m *MockClientInterface) TraceBackplaneRemediateClusterIdRemediation(arg0 context.Context, arg1, arg2 string, arg3 ...Openapi.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TraceBackplaneRemediateClusterIdRemediation", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TraceBackplaneRemediateClusterIdRemediation indicates an expected call of TraceBackplaneRemediateClusterIdRemediation.
func (mr *MockClientInterfaceMockRecorder) TraceBackplaneRemediateClusterIdRemediation(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TraceBackplaneRemediateClusterIdRemediation", reflect.TypeOf((*MockClientInterface)(nil).TraceBackplaneRemediateClusterIdRemediation), varargs...)
}
