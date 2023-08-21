package core

import (
	"github.com/elbombardi/squrl/src/api_service/core"
	"github.com/stretchr/testify/mock"
)

type MockAccountsManager struct {
	mock.Mock
}

func (m *MockAccountsManager) Create(params *core.CreateAccountParams, user *core.User) (*core.CreateAccountResponse, error) {
	args := m.Called(params, user)

	if rf, ok := args.Get(0).(func(*core.CreateAccountParams, *core.User) (*core.CreateAccountResponse, error)); ok {
		return rf(params, user)
	}
	return args.Get(0).(*core.CreateAccountResponse), args.Error(1)
}

func (m *MockAccountsManager) Update(params *core.UpdateAccountParams, user *core.User) (*core.UpdateAccountResponse, error) {
	args := m.Called(params, user)

	if rf, ok := args.Get(0).(func(*core.UpdateAccountParams, *core.User) (*core.UpdateAccountResponse, error)); ok {
		return rf(params, user)
	}
	return args.Get(0).(*core.UpdateAccountResponse), args.Error(1)
}
