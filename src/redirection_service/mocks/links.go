package mocks

import (
	"github.com/elbombardi/squrl/src/redirection_service/core"
	"github.com/stretchr/testify/mock"
)

type MockLinksManager struct {
	mock.Mock
}

func (m *MockLinksManager) Resolve(params *core.ResolveLinkParams) (*core.Link, error) {
	args := m.Called(params)

	if rf, ok := args.Get(0).(func(*core.ResolveLinkParams) (*core.Link, error)); ok {
		return rf(params)
	}
	return args.Get(0).(*core.Link), args.Error(1)
}
