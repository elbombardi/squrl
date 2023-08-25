package core

import (
	"github.com/stretchr/testify/mock"
)

type MockAccountsManager struct {
	mock.Mock
}

func (m *MockAccountsManager) Create(params *CreateAccountParams, user *User) (*CreateAccountResponse, error) {
	args := m.Called(params, user)

	if rf, ok := args.Get(0).(func(*CreateAccountParams, *User) (*CreateAccountResponse, error)); ok {
		return rf(params, user)
	}
	return args.Get(0).(*CreateAccountResponse), args.Error(1)
}

func (m *MockAccountsManager) Update(params *UpdateAccountParams, user *User) (*UpdateAccountResponse, error) {
	args := m.Called(params, user)

	if rf, ok := args.Get(0).(func(*UpdateAccountParams, *User) (*UpdateAccountResponse, error)); ok {
		return rf(params, user)
	}
	return args.Get(0).(*UpdateAccountResponse), args.Error(1)
}
