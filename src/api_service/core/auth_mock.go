package core

import (
	"github.com/stretchr/testify/mock"
)

type MockAuthenticator struct {
	mock.Mock
}

func (m *MockAuthenticator) Authenticate(username string, password string) (token string, err error) {
	args := m.Called(username, password)

	if rf, ok := args.Get(0).(func(username string, password string) (token string, err error)); ok {
		return rf(username, password)
	}
	return args.String(0), args.Error(1)
}

func (m *MockAuthenticator) Validate(token string) (user *User, err error) {
	args := m.Called(token)

	if rf, ok := args.Get(0).(func(token string) (user *User, err error)); ok {
		return rf(token)
	}
	return args.Get(0).(*User), args.Error(1)
}
