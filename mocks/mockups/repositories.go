// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/ports/repositories.go

// Package mockups is a generated GoMock package.
package mockups

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/mkantzer/emojiSorter/internal/core/domain"
)

// MockEmojiRepository is a mock of EmojiRepository interface.
type MockEmojiRepository struct {
	ctrl     *gomock.Controller
	recorder *MockEmojiRepositoryMockRecorder
}

// MockEmojiRepositoryMockRecorder is the mock recorder for MockEmojiRepository.
type MockEmojiRepositoryMockRecorder struct {
	mock *MockEmojiRepository
}

// NewMockEmojiRepository creates a new mock instance.
func NewMockEmojiRepository(ctrl *gomock.Controller) *MockEmojiRepository {
	mock := &MockEmojiRepository{ctrl: ctrl}
	mock.recorder = &MockEmojiRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmojiRepository) EXPECT() *MockEmojiRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockEmojiRepository) Get(name string) (domain.Emoji, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", name)
	ret0, _ := ret[0].(domain.Emoji)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockEmojiRepositoryMockRecorder) Get(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockEmojiRepository)(nil).Get), name)
}

// Save mocks base method.
func (m *MockEmojiRepository) Save(arg0 domain.Emoji) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockEmojiRepositoryMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockEmojiRepository)(nil).Save), arg0)
}