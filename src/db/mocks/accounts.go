package mocks

import (
	"context"

	"github.com/elbombardi/squrl/src/db"
	"github.com/stretchr/testify/mock"
)

type MockAccountRepository struct {
	mock.Mock
}

func (r *MockAccountRepository) CheckPrefixExists(ctx context.Context, prefix string) (bool, error) {

	args := r.Called(ctx, prefix)

	if rf, ok := args.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, prefix)
	}
	return args.Bool(0), args.Error(1)
}

func (r *MockAccountRepository) CheckUsernameExists(ctx context.Context, username string) (bool, error) {
	args := r.Called(ctx, username)

	if rf, ok := args.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, username)
	}
	return args.Bool(0), args.Error(1)
}

func (r *MockAccountRepository) GetAccountByPrefix(ctx context.Context, prefix string) (db.Account, error) {
	args := r.Called(ctx, prefix)

	if rf, ok := args.Get(0).(func(context.Context, string) (db.Account, error)); ok {
		return rf(ctx, prefix)
	}
	return args.Get(0).(db.Account), args.Error(1)
}

func (r *MockAccountRepository) GetAccountByUsername(ctx context.Context, username string) (db.Account, error) {
	args := r.Called(ctx, username)

	if rf, ok := args.Get(0).(func(context.Context, string) (db.Account, error)); ok {
		return rf(ctx, username)
	}
	return args.Get(0).(db.Account), args.Error(1)
}

func (r *MockAccountRepository) InsertNewAccount(ctx context.Context, params db.InsertNewAccountParams) error {
	args := r.Called(ctx, params)

	if rf, ok := args.Get(0).(func(context.Context, db.InsertNewAccountParams) error); ok {
		return rf(ctx, params)
	}
	return args.Error(0)
}

func (r *MockAccountRepository) UpdateAccountStatusByUsername(ctx context.Context, params db.UpdateAccountStatusByUsernameParams) error {
	args := r.Called(ctx, params)

	if rf, ok := args.Get(0).(func(context.Context, db.UpdateAccountStatusByUsernameParams) error); ok {
		return rf(ctx, params)
	}
	return args.Error(0)
}
