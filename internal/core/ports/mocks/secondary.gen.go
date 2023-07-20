// Code generated by MockGen. DO NOT EDIT.
// Source: secondary.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entity "backend-api/internal/core/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockChannelRepository is a mock of ChannelRepository interface.
type MockChannelRepository struct {
	ctrl     *gomock.Controller
	recorder *MockChannelRepositoryMockRecorder
}

// MockChannelRepositoryMockRecorder is the mock recorder for MockChannelRepository.
type MockChannelRepositoryMockRecorder struct {
	mock *MockChannelRepository
}

// NewMockChannelRepository creates a new mock instance.
func NewMockChannelRepository(ctrl *gomock.Controller) *MockChannelRepository {
	mock := &MockChannelRepository{ctrl: ctrl}
	mock.recorder = &MockChannelRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChannelRepository) EXPECT() *MockChannelRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockChannelRepository) Create(channel entity.Channel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", channel)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockChannelRepositoryMockRecorder) Create(channel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockChannelRepository)(nil).Create), channel)
}

// FindAll mocks base method.
func (m *MockChannelRepository) FindAll() ([]entity.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]entity.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockChannelRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockChannelRepository)(nil).FindAll))
}