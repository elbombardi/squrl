package core

import (
	"github.com/stretchr/testify/mock"
)

type MockLinksManager struct {
	mock.Mock
}

func (m *MockLinksManager) Shorten(longUrl string, user *User) (*Link, error) {
	arg := m.Called(longUrl, user)
	if rf, ok := arg.Get(0).(func(longUrl string, user *User) (*Link, error)); ok {
		return rf(longUrl, user)
	}
	return arg.Get(0).(*Link), arg.Error(1)
}

func (m *MockLinksManager) Update(params *LinkUpdateParams, user *User) (*Link, error) {
	args := m.Called(params, user)
	if rf, ok := args.Get(0).(func(params *LinkUpdateParams, user *User) (*Link, error)); ok {
		return rf(params, user)
	}
	return args.Get(0).(*Link), args.Error(1)
}
