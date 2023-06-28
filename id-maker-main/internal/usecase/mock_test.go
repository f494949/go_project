// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package usecase_test is a generated GoMock package.
package usecase_test

import (
	entity "id-maker/internal/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSegment is a mock of Segment interface.
type MockSegment struct {
	ctrl     *gomock.Controller
	recorder *MockSegmentMockRecorder
}

// MockSegmentMockRecorder is the mock recorder for MockSegment.
type MockSegmentMockRecorder struct {
	mock *MockSegment
}

// NewMockSegment creates a new mock instance.
func NewMockSegment(ctrl *gomock.Controller) *MockSegment {
	mock := &MockSegment{ctrl: ctrl}
	mock.recorder = &MockSegmentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSegment) EXPECT() *MockSegmentMockRecorder {
	return m.recorder
}

// CreateTag mocks base method.
func (m *MockSegment) CreateTag(arg0 *entity.Segments) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTag", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTag indicates an expected call of CreateTag.
func (mr *MockSegmentMockRecorder) CreateTag(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTag", reflect.TypeOf((*MockSegment)(nil).CreateTag), arg0)
}

// GetId mocks base method.
func (m *MockSegment) GetId(arg0 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetId", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetId indicates an expected call of GetId.
func (mr *MockSegmentMockRecorder) GetId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetId", reflect.TypeOf((*MockSegment)(nil).GetId), arg0)
}

// SnowFlakeGetId mocks base method.
func (m *MockSegment) SnowFlakeGetId() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SnowFlakeGetId")
	ret0, _ := ret[0].(int64)
	return ret0
}

// SnowFlakeGetId indicates an expected call of SnowFlakeGetId.
func (mr *MockSegmentMockRecorder) SnowFlakeGetId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnowFlakeGetId", reflect.TypeOf((*MockSegment)(nil).SnowFlakeGetId))
}

// MockSegmentRepo is a mock of SegmentRepo interface.
type MockSegmentRepo struct {
	ctrl     *gomock.Controller
	recorder *MockSegmentRepoMockRecorder
}

// MockSegmentRepoMockRecorder is the mock recorder for MockSegmentRepo.
type MockSegmentRepoMockRecorder struct {
	mock *MockSegmentRepo
}

// NewMockSegmentRepo creates a new mock instance.
func NewMockSegmentRepo(ctrl *gomock.Controller) *MockSegmentRepo {
	mock := &MockSegmentRepo{ctrl: ctrl}
	mock.recorder = &MockSegmentRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSegmentRepo) EXPECT() *MockSegmentRepoMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockSegmentRepo) Add(arg0 *entity.Segments) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockSegmentRepoMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockSegmentRepo)(nil).Add), arg0)
}

// GetList mocks base method.
func (m *MockSegmentRepo) GetList() ([]entity.Segments, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList")
	ret0, _ := ret[0].([]entity.Segments)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList.
func (mr *MockSegmentRepoMockRecorder) GetList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockSegmentRepo)(nil).GetList))
}

// GetNextId mocks base method.
func (m *MockSegmentRepo) GetNextId(arg0 string) (*entity.Segments, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextId", arg0)
	ret0, _ := ret[0].(*entity.Segments)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextId indicates an expected call of GetNextId.
func (mr *MockSegmentRepoMockRecorder) GetNextId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextId", reflect.TypeOf((*MockSegmentRepo)(nil).GetNextId), arg0)
}
