package db

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockLinkRepository struct {
	mock.Mock
}

func (r *MockLinkRepository) CheckShortUrlKeyExists(ctx context.Context, params CheckShortUrlKeyExistsParams) (bool, error) {
	args := r.Called(ctx, params)

	if rf, ok := args.Get(0).(func(context.Context, CheckShortUrlKeyExistsParams) (bool, error)); ok {
		return rf(ctx, params)
	}
	return args.Bool(0), args.Error(1)
}

func (r *MockLinkRepository) GetLinkByAccountIDAndShortURLKey(ctx context.Context, params GetLinkByAccountIDAndShortURLKeyParams) (Link, error) {
	args := r.Called(ctx, params)

	if rf, ok := args.Get(0).(func(context.Context, GetLinkByAccountIDAndShortURLKeyParams) (Link, error)); ok {
		return rf(ctx, params)
	}
	return args.Get(0).(Link), args.Error(1)
}

func (r *MockLinkRepository) InsertNewLink(ctx context.Context, params InsertNewLinkParams) error {
	args := r.Called(ctx, params)

	if rf, ok := args.Get(0).(func(context.Context, InsertNewLinkParams) error); ok {
		return rf(ctx, params)
	}
	return args.Error(0)
}

func (r *MockLinkRepository) UpdateLinkLongURL(ctx context.Context, params UpdateLinkLongURLParams) error {
	args := r.Called(ctx, params)

	if rf, ok := args.Get(0).(func(context.Context, UpdateLinkLongURLParams) error); ok {
		return rf(ctx, params)
	}
	return args.Error(0)
}

func (r *MockLinkRepository) UpdateLinkStatus(ctx context.Context, params UpdateLinkStatusParams) error {
	args := r.Called(ctx, params)

	if rf, ok := args.Get(0).(func(context.Context, UpdateLinkStatusParams) error); ok {
		return rf(ctx, params)
	}
	return args.Error(0)
}

func (r *MockLinkRepository) UpdateLinkTrackingStatus(ctx context.Context, params UpdateLinkTrackingStatusParams) error {
	args := r.Called(ctx, params)

	if rf, ok := args.Get(0).(func(context.Context, UpdateLinkTrackingStatusParams) error); ok {
		return rf(ctx, params)
	}
	return args.Error(0)
}
