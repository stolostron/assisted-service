// Code generated by mockery v1.0.0. DO NOT EDIT.

package restapi

import (
	context "context"

	installer "github.com/filanov/bm-inventory/restapi/operations/installer"
	middleware "github.com/go-openapi/runtime/middleware"

	mock "github.com/stretchr/testify/mock"
)

// MockInstallerAPI is an autogenerated mock type for the InstallerAPI type
type MockInstallerAPI struct {
	mock.Mock
}

// DeregisterCluster provides a mock function with given fields: ctx, params
func (_m *MockInstallerAPI) DeregisterCluster(ctx context.Context, params installer.DeregisterClusterParams) middleware.Responder {
	ret := _m.Called(ctx, params)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(context.Context, installer.DeregisterClusterParams) middleware.Responder); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}

// DeregisterHost provides a mock function with given fields: ctx, params
func (_m *MockInstallerAPI) DeregisterHost(ctx context.Context, params installer.DeregisterHostParams) middleware.Responder {
	ret := _m.Called(ctx, params)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(context.Context, installer.DeregisterHostParams) middleware.Responder); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}

// DisableHost provides a mock function with given fields: ctx, params
func (_m *MockInstallerAPI) DisableHost(ctx context.Context, params installer.DisableHostParams) middleware.Responder {
	ret := _m.Called(ctx, params)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(context.Context, installer.DisableHostParams) middleware.Responder); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}

// DownloadClusterFiles provides a mock function with given fields: ctx, params
func (_m *MockInstallerAPI) DownloadClusterFiles(ctx context.Context, params installer.DownloadClusterFilesParams) middleware.Responder {
	ret := _m.Called(ctx, params)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(context.Context, installer.DownloadClusterFilesParams) middleware.Responder); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}

// DownloadClusterISO provides a mock function with given fields: ctx, params
func (_m *MockInstallerAPI) DownloadClusterISO(ctx context.Context, params installer.DownloadClusterISOParams) middleware.Responder {
	ret := _m.Called(ctx, params)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(context.Context, installer.DownloadClusterISOParams) middleware.Responder); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}

// EnableHost provides a mock function with given fields: ctx, params
func (_m *MockInstallerAPI) EnableHost(ctx context.Context, params installer.EnableHostParams) middleware.Responder {
	ret := _m.Called(ctx, params)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(context.Context, installer.EnableHostParams) middleware.Responder); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}

// GenerateClusterISO provides a mock function with given fields: ctx, params
func (_m *MockInstallerAPI) GenerateClusterISO(ctx context.Context, params installer.GenerateClusterISOParams) middleware.Responder {
	ret := _m.Called(ctx, params)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(context.Context, installer.GenerateClusterISOParams) middleware.Responder); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}

// GetCluster provides a mock function with given fields: ctx, params
func (_m *MockInstallerAPI) GetCluster(ctx context.Context, params installer.GetClusterParams) middleware.Responder {
	ret := _m.Called(ctx, params)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(context.Context, installer.GetClusterParams) middleware.Responder); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}

// GetCredentials provides a mock function with given fields: ctx, params
func (_m *MockInstallerAPI) GetCredentials(ctx context.Context, params installer.GetCredentialsParams) middleware.Responder {
	ret := _m.Called(ctx, params)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(context.Context, installer.GetCredentialsParams) middleware.Responder); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}

// GetHost provides a mock function with given fields: ctx, params
func (_m *MockInstallerAPI) GetHost(ctx context.Context, params installer.GetHostParams) middleware.Responder {
	ret := _m.Called(ctx, params)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(context.Context, installer.GetHostParams) middleware.Responder); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}

// GetNextSteps provides a mock function with given fields: ctx, params
func (_m *MockInstallerAPI) GetNextSteps(ctx context.Context, params installer.GetNextStepsParams) middleware.Responder {
	ret := _m.Called(ctx, params)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(context.Context, installer.GetNextStepsParams) middleware.Responder); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}

// InstallCluster provides a mock function with given fields: ctx, params
func (_m *MockInstallerAPI) InstallCluster(ctx context.Context, params installer.InstallClusterParams) middleware.Responder {
	ret := _m.Called(ctx, params)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(context.Context, installer.InstallClusterParams) middleware.Responder); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}

// ListClusters provides a mock function with given fields: ctx, params
func (_m *MockInstallerAPI) ListClusters(ctx context.Context, params installer.ListClustersParams) middleware.Responder {
	ret := _m.Called(ctx, params)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(context.Context, installer.ListClustersParams) middleware.Responder); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}

// ListHosts provides a mock function with given fields: ctx, params
func (_m *MockInstallerAPI) ListHosts(ctx context.Context, params installer.ListHostsParams) middleware.Responder {
	ret := _m.Called(ctx, params)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(context.Context, installer.ListHostsParams) middleware.Responder); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}

// PostStepReply provides a mock function with given fields: ctx, params
func (_m *MockInstallerAPI) PostStepReply(ctx context.Context, params installer.PostStepReplyParams) middleware.Responder {
	ret := _m.Called(ctx, params)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(context.Context, installer.PostStepReplyParams) middleware.Responder); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}

// RegisterCluster provides a mock function with given fields: ctx, params
func (_m *MockInstallerAPI) RegisterCluster(ctx context.Context, params installer.RegisterClusterParams) middleware.Responder {
	ret := _m.Called(ctx, params)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(context.Context, installer.RegisterClusterParams) middleware.Responder); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}

// RegisterHost provides a mock function with given fields: ctx, params
func (_m *MockInstallerAPI) RegisterHost(ctx context.Context, params installer.RegisterHostParams) middleware.Responder {
	ret := _m.Called(ctx, params)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(context.Context, installer.RegisterHostParams) middleware.Responder); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}

// SetDebugStep provides a mock function with given fields: ctx, params
func (_m *MockInstallerAPI) SetDebugStep(ctx context.Context, params installer.SetDebugStepParams) middleware.Responder {
	ret := _m.Called(ctx, params)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(context.Context, installer.SetDebugStepParams) middleware.Responder); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}

// UpdateCluster provides a mock function with given fields: ctx, params
func (_m *MockInstallerAPI) UpdateCluster(ctx context.Context, params installer.UpdateClusterParams) middleware.Responder {
	ret := _m.Called(ctx, params)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(context.Context, installer.UpdateClusterParams) middleware.Responder); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}

// UpdateHostInstallProgress provides a mock function with given fields: ctx, params
func (_m *MockInstallerAPI) UpdateHostInstallProgress(ctx context.Context, params installer.UpdateHostInstallProgressParams) middleware.Responder {
	ret := _m.Called(ctx, params)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(context.Context, installer.UpdateHostInstallProgressParams) middleware.Responder); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}
