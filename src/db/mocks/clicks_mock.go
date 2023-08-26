package db

import (
	"context"

	"github.com/elbombardi/squrl/src/db"
	"github.com/stretchr/testify/mock"
)

type MockClickRepository struct {
	mock.Mock
}

func (r *MockClickRepository) InsertNewClick(ctx context.Context, params db.InsertNewClickParams) error {
	args := r.Called(ctx, params)

	if rf, ok := args.Get(0).(func(context.Context, db.InsertNewClickParams) error); ok {
		return rf(ctx, params)
	}
	return args.Error(0)
}
