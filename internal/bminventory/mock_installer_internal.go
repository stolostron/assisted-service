// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/assisted-service/internal/bminventory (interfaces: InstallerInternals)

// Package bminventory is a generated GoMock package.
package bminventory

import (
	context "context"
	strfmt "github.com/go-openapi/strfmt"
	gomock "github.com/golang/mock/gomock"
	common "github.com/openshift/assisted-service/internal/common"
	models "github.com/openshift/assisted-service/models"
	installer "github.com/openshift/assisted-service/restapi/operations/installer"
	io "io"
	types "k8s.io/apimachinery/pkg/types"
	reflect "reflect"
)

// MockInstallerInternals is a mock of InstallerInternals interface
type MockInstallerInternals struct {
	ctrl     *gomock.Controller
	recorder *MockInstallerInternalsMockRecorder
}

// MockInstallerInternalsMockRecorder is the mock recorder for MockInstallerInternals
type MockInstallerInternalsMockRecorder struct {
	mock *MockInstallerInternals
}

// NewMockInstallerInternals creates a new mock instance
func NewMockInstallerInternals(ctrl *gomock.Controller) *MockInstallerInternals {
	mock := &MockInstallerInternals{ctrl: ctrl}
	mock.recorder = &MockInstallerInternalsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInstallerInternals) EXPECT() *MockInstallerInternalsMockRecorder {
	return m.recorder
}

// AddOpenshiftVersion mocks base method
func (m *MockInstallerInternals) AddOpenshiftVersion(arg0 context.Context, arg1, arg2 string) (*models.OpenshiftVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOpenshiftVersion", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.OpenshiftVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddOpenshiftVersion indicates an expected call of AddOpenshiftVersion
func (mr *MockInstallerInternalsMockRecorder) AddOpenshiftVersion(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOpenshiftVersion", reflect.TypeOf((*MockInstallerInternals)(nil).AddOpenshiftVersion), arg0, arg1, arg2)
}

// CancelInstallationInternal mocks base method
func (m *MockInstallerInternals) CancelInstallationInternal(arg0 context.Context, arg1 installer.CancelInstallationParams) (*common.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelInstallationInternal", arg0, arg1)
	ret0, _ := ret[0].(*common.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelInstallationInternal indicates an expected call of CancelInstallationInternal
func (mr *MockInstallerInternalsMockRecorder) CancelInstallationInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelInstallationInternal", reflect.TypeOf((*MockInstallerInternals)(nil).CancelInstallationInternal), arg0, arg1)
}

// DeregisterClusterInternal mocks base method
func (m *MockInstallerInternals) DeregisterClusterInternal(arg0 context.Context, arg1 installer.DeregisterClusterParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeregisterClusterInternal", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeregisterClusterInternal indicates an expected call of DeregisterClusterInternal
func (mr *MockInstallerInternalsMockRecorder) DeregisterClusterInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterClusterInternal", reflect.TypeOf((*MockInstallerInternals)(nil).DeregisterClusterInternal), arg0, arg1)
}

// DeregisterHostInternal mocks base method
func (m *MockInstallerInternals) DeregisterHostInternal(arg0 context.Context, arg1 installer.DeregisterHostParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeregisterHostInternal", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeregisterHostInternal indicates an expected call of DeregisterHostInternal
func (mr *MockInstallerInternalsMockRecorder) DeregisterHostInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterHostInternal", reflect.TypeOf((*MockInstallerInternals)(nil).DeregisterHostInternal), arg0, arg1)
}

// DownloadClusterFilesInternal mocks base method
func (m *MockInstallerInternals) DownloadClusterFilesInternal(arg0 context.Context, arg1 installer.DownloadClusterFilesParams) (io.ReadCloser, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadClusterFilesInternal", arg0, arg1)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DownloadClusterFilesInternal indicates an expected call of DownloadClusterFilesInternal
func (mr *MockInstallerInternalsMockRecorder) DownloadClusterFilesInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadClusterFilesInternal", reflect.TypeOf((*MockInstallerInternals)(nil).DownloadClusterFilesInternal), arg0, arg1)
}

// GenerateClusterISOInternal mocks base method
func (m *MockInstallerInternals) GenerateClusterISOInternal(arg0 context.Context, arg1 installer.GenerateClusterISOParams) (*common.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateClusterISOInternal", arg0, arg1)
	ret0, _ := ret[0].(*common.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateClusterISOInternal indicates an expected call of GenerateClusterISOInternal
func (mr *MockInstallerInternalsMockRecorder) GenerateClusterISOInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateClusterISOInternal", reflect.TypeOf((*MockInstallerInternals)(nil).GenerateClusterISOInternal), arg0, arg1)
}

// GetClusterByKubeKey mocks base method
func (m *MockInstallerInternals) GetClusterByKubeKey(arg0 types.NamespacedName) (*common.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterByKubeKey", arg0)
	ret0, _ := ret[0].(*common.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterByKubeKey indicates an expected call of GetClusterByKubeKey
func (mr *MockInstallerInternalsMockRecorder) GetClusterByKubeKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterByKubeKey", reflect.TypeOf((*MockInstallerInternals)(nil).GetClusterByKubeKey), arg0)
}

// GetClusterInternal mocks base method
func (m *MockInstallerInternals) GetClusterInternal(arg0 context.Context, arg1 installer.GetClusterParams) (*common.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterInternal", arg0, arg1)
	ret0, _ := ret[0].(*common.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterInternal indicates an expected call of GetClusterInternal
func (mr *MockInstallerInternalsMockRecorder) GetClusterInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterInternal", reflect.TypeOf((*MockInstallerInternals)(nil).GetClusterInternal), arg0, arg1)
}

// GetCommonHostInternal mocks base method
func (m *MockInstallerInternals) GetCommonHostInternal(arg0 context.Context, arg1, arg2 string) (*common.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommonHostInternal", arg0, arg1, arg2)
	ret0, _ := ret[0].(*common.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommonHostInternal indicates an expected call of GetCommonHostInternal
func (mr *MockInstallerInternalsMockRecorder) GetCommonHostInternal(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommonHostInternal", reflect.TypeOf((*MockInstallerInternals)(nil).GetCommonHostInternal), arg0, arg1, arg2)
}

// GetCredentialsInternal mocks base method
func (m *MockInstallerInternals) GetCredentialsInternal(arg0 context.Context, arg1 installer.GetCredentialsParams) (*models.Credentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCredentialsInternal", arg0, arg1)
	ret0, _ := ret[0].(*models.Credentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredentialsInternal indicates an expected call of GetCredentialsInternal
func (mr *MockInstallerInternalsMockRecorder) GetCredentialsInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentialsInternal", reflect.TypeOf((*MockInstallerInternals)(nil).GetCredentialsInternal), arg0, arg1)
}

// GetHostByKubeKey mocks base method
func (m *MockInstallerInternals) GetHostByKubeKey(arg0 types.NamespacedName) (*common.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostByKubeKey", arg0)
	ret0, _ := ret[0].(*common.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostByKubeKey indicates an expected call of GetHostByKubeKey
func (mr *MockInstallerInternalsMockRecorder) GetHostByKubeKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostByKubeKey", reflect.TypeOf((*MockInstallerInternals)(nil).GetHostByKubeKey), arg0)
}

// InstallClusterInternal mocks base method
func (m *MockInstallerInternals) InstallClusterInternal(arg0 context.Context, arg1 installer.InstallClusterParams) (*common.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallClusterInternal", arg0, arg1)
	ret0, _ := ret[0].(*common.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstallClusterInternal indicates an expected call of InstallClusterInternal
func (mr *MockInstallerInternalsMockRecorder) InstallClusterInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallClusterInternal", reflect.TypeOf((*MockInstallerInternals)(nil).InstallClusterInternal), arg0, arg1)
}

// InstallSingleDay2HostInternal mocks base method
func (m *MockInstallerInternals) InstallSingleDay2HostInternal(arg0 context.Context, arg1, arg2 strfmt.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallSingleDay2HostInternal", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallSingleDay2HostInternal indicates an expected call of InstallSingleDay2HostInternal
func (mr *MockInstallerInternalsMockRecorder) InstallSingleDay2HostInternal(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallSingleDay2HostInternal", reflect.TypeOf((*MockInstallerInternals)(nil).InstallSingleDay2HostInternal), arg0, arg1, arg2)
}

// RegisterAddHostsClusterInternal mocks base method
func (m *MockInstallerInternals) RegisterAddHostsClusterInternal(arg0 context.Context, arg1 *types.NamespacedName, arg2 installer.RegisterAddHostsClusterParams) (*common.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterAddHostsClusterInternal", arg0, arg1, arg2)
	ret0, _ := ret[0].(*common.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterAddHostsClusterInternal indicates an expected call of RegisterAddHostsClusterInternal
func (mr *MockInstallerInternalsMockRecorder) RegisterAddHostsClusterInternal(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAddHostsClusterInternal", reflect.TypeOf((*MockInstallerInternals)(nil).RegisterAddHostsClusterInternal), arg0, arg1, arg2)
}

// RegisterClusterInternal mocks base method
func (m *MockInstallerInternals) RegisterClusterInternal(arg0 context.Context, arg1 *types.NamespacedName, arg2 installer.RegisterClusterParams) (*common.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterClusterInternal", arg0, arg1, arg2)
	ret0, _ := ret[0].(*common.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterClusterInternal indicates an expected call of RegisterClusterInternal
func (mr *MockInstallerInternalsMockRecorder) RegisterClusterInternal(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterClusterInternal", reflect.TypeOf((*MockInstallerInternals)(nil).RegisterClusterInternal), arg0, arg1, arg2)
}

// TransformClusterToDay2Internal mocks base method
func (m *MockInstallerInternals) TransformClusterToDay2Internal(arg0 context.Context, arg1 strfmt.UUID) (*common.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransformClusterToDay2Internal", arg0, arg1)
	ret0, _ := ret[0].(*common.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransformClusterToDay2Internal indicates an expected call of TransformClusterToDay2Internal
func (mr *MockInstallerInternalsMockRecorder) TransformClusterToDay2Internal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransformClusterToDay2Internal", reflect.TypeOf((*MockInstallerInternals)(nil).TransformClusterToDay2Internal), arg0, arg1)
}

// UpdateClusterInstallConfigInternal mocks base method
func (m *MockInstallerInternals) UpdateClusterInstallConfigInternal(arg0 context.Context, arg1 installer.UpdateClusterInstallConfigParams) (*common.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClusterInstallConfigInternal", arg0, arg1)
	ret0, _ := ret[0].(*common.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateClusterInstallConfigInternal indicates an expected call of UpdateClusterInstallConfigInternal
func (mr *MockInstallerInternalsMockRecorder) UpdateClusterInstallConfigInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClusterInstallConfigInternal", reflect.TypeOf((*MockInstallerInternals)(nil).UpdateClusterInstallConfigInternal), arg0, arg1)
}

// UpdateClusterNonInteractive mocks base method
func (m *MockInstallerInternals) UpdateClusterNonInteractive(arg0 context.Context, arg1 installer.UpdateClusterParams) (*common.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClusterNonInteractive", arg0, arg1)
	ret0, _ := ret[0].(*common.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateClusterNonInteractive indicates an expected call of UpdateClusterNonInteractive
func (mr *MockInstallerInternalsMockRecorder) UpdateClusterNonInteractive(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClusterNonInteractive", reflect.TypeOf((*MockInstallerInternals)(nil).UpdateClusterNonInteractive), arg0, arg1)
}

// UpdateDiscoveryIgnitionInternal mocks base method
func (m *MockInstallerInternals) UpdateDiscoveryIgnitionInternal(arg0 context.Context, arg1 installer.UpdateDiscoveryIgnitionParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDiscoveryIgnitionInternal", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDiscoveryIgnitionInternal indicates an expected call of UpdateDiscoveryIgnitionInternal
func (mr *MockInstallerInternalsMockRecorder) UpdateDiscoveryIgnitionInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDiscoveryIgnitionInternal", reflect.TypeOf((*MockInstallerInternals)(nil).UpdateDiscoveryIgnitionInternal), arg0, arg1)
}

// UpdateHostApprovedInternal mocks base method
func (m *MockInstallerInternals) UpdateHostApprovedInternal(arg0 context.Context, arg1, arg2 string, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHostApprovedInternal", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHostApprovedInternal indicates an expected call of UpdateHostApprovedInternal
func (mr *MockInstallerInternalsMockRecorder) UpdateHostApprovedInternal(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHostApprovedInternal", reflect.TypeOf((*MockInstallerInternals)(nil).UpdateHostApprovedInternal), arg0, arg1, arg2, arg3)
}

// UpdateHostIgnitionInternal mocks base method
func (m *MockInstallerInternals) UpdateHostIgnitionInternal(arg0 context.Context, arg1 installer.UpdateHostIgnitionParams) (*models.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHostIgnitionInternal", arg0, arg1)
	ret0, _ := ret[0].(*models.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateHostIgnitionInternal indicates an expected call of UpdateHostIgnitionInternal
func (mr *MockInstallerInternalsMockRecorder) UpdateHostIgnitionInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHostIgnitionInternal", reflect.TypeOf((*MockInstallerInternals)(nil).UpdateHostIgnitionInternal), arg0, arg1)
}

// UpdateHostInstallerArgsInternal mocks base method
func (m *MockInstallerInternals) UpdateHostInstallerArgsInternal(arg0 context.Context, arg1 installer.UpdateHostInstallerArgsParams) (*models.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHostInstallerArgsInternal", arg0, arg1)
	ret0, _ := ret[0].(*models.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateHostInstallerArgsInternal indicates an expected call of UpdateHostInstallerArgsInternal
func (mr *MockInstallerInternalsMockRecorder) UpdateHostInstallerArgsInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHostInstallerArgsInternal", reflect.TypeOf((*MockInstallerInternals)(nil).UpdateHostInstallerArgsInternal), arg0, arg1)
}
