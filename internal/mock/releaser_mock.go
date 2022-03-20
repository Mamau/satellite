// Code generated by MockGen. DO NOT EDIT.
// Source: satellite/internal/updater (interfaces: Releaser)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	updater "satellite/internal/updater"

	gomock "github.com/golang/mock/gomock"
)

// MockReleaser is a mock of Releaser interface.
type MockReleaser struct {
	ctrl     *gomock.Controller
	recorder *MockReleaserMockRecorder
}

// MockReleaserMockRecorder is the mock recorder for MockReleaser.
type MockReleaserMockRecorder struct {
	mock *MockReleaser
}

// NewMockReleaser creates a new mock instance.
func NewMockReleaser(ctrl *gomock.Controller) *MockReleaser {
	mock := &MockReleaser{ctrl: ctrl}
	mock.recorder = &MockReleaserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReleaser) EXPECT() *MockReleaserMockRecorder {
	return m.recorder
}

// FetchRelease mocks base method.
func (m *MockReleaser) FetchRelease() *updater.Release {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchRelease")
	ret0, _ := ret[0].(*updater.Release)
	return ret0
}

// FetchRelease indicates an expected call of FetchRelease.
func (mr *MockReleaserMockRecorder) FetchRelease() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchRelease", reflect.TypeOf((*MockReleaser)(nil).FetchRelease))
}