package db

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockClickRepository struct {
	mock.Mock
}

func (r *MockClickRepository) InsertNewClick(ctx context.Context, params InsertNewClickParams) error {
	args := r.Called(ctx, params)

	if rf, ok := args.Get(0).(func(context.Context, InsertNewClickParams) error); ok {
		return rf(ctx, params)
	}
	return args.Error(0)
}
