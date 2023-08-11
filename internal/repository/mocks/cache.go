// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dewzzjr/ais/internal/repository (interfaces: CacheArticle)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/dewzzjr/ais/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockCacheArticle is a mock of CacheArticle interface.
type MockCacheArticle struct {
	ctrl     *gomock.Controller
	recorder *MockCacheArticleMockRecorder
}

// MockCacheArticleMockRecorder is the mock recorder for MockCacheArticle.
type MockCacheArticleMockRecorder struct {
	mock *MockCacheArticle
}

// NewMockCacheArticle creates a new mock instance.
func NewMockCacheArticle(ctrl *gomock.Controller) *MockCacheArticle {
	mock := &MockCacheArticle{ctrl: ctrl}
	mock.recorder = &MockCacheArticleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheArticle) EXPECT() *MockCacheArticleMockRecorder {
	return m.recorder
}

// Del mocks base method.
func (m *MockCacheArticle) Del(arg0 context.Context, arg1 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Del", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockCacheArticleMockRecorder) Del(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockCacheArticle)(nil).Del), varargs...)
}

// Get mocks base method.
func (m *MockCacheArticle) Get(arg0 context.Context, arg1 string) (model.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(model.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCacheArticleMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCacheArticle)(nil).Get), arg0, arg1)
}

// Set mocks base method.
func (m *MockCacheArticle) Set(arg0 context.Context, arg1 string, arg2 model.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockCacheArticleMockRecorder) Set(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCacheArticle)(nil).Set), arg0, arg1, arg2)
}
