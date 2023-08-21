package mocks

import (
	"context"

	"github.com/elbombardi/squrl/src/db"
	"github.com/stretchr/testify/mock"
)

type MockLinkRepository struct {
	mock.Mock
}

func (r *MockLinkRepository) CheckShortUrlKeyExists(ctx context.Context, params db.CheckShortUrlKeyExistsParams) (bool, error) {
	args := r.Called(ctx, params)

	if rf, ok := args.Get(0).(func(context.Context, db.CheckShortUrlKeyExistsParams) (bool, error)); ok {
		return rf(ctx, params)
	}
	return args.Bool(0), args.Error(1)
}

func (r *MockLinkRepository) GetLinkByAccountIDAndShortURLKey(ctx context.Context, params db.GetLinkByAccountIDAndShortURLKeyParams) (db.Link, error) {
	args := r.Called(ctx, params)

	if rf, ok := args.Get(0).(func(context.Context, db.GetLinkByAccountIDAndShortURLKeyParams) (db.Link, error)); ok {
		return rf(ctx, params)
	}
	return args.Get(0).(db.Link), args.Error(1)
}

func (r *MockLinkRepository) InsertNewLink(ctx context.Context, params db.InsertNewLinkParams) error {
	args := r.Called(ctx, params)

	if rf, ok := args.Get(0).(func(context.Context, db.InsertNewLinkParams) error); ok {
		return rf(ctx, params)
	}
	return args.Error(0)
}

func (r *MockLinkRepository) UpdateLinkLongURL(ctx context.Context, params db.UpdateLinkLongURLParams) error {
	args := r.Called(ctx, params)

	if rf, ok := args.Get(0).(func(context.Context, db.UpdateLinkLongURLParams) error); ok {
		return rf(ctx, params)
	}
	return args.Error(0)
}

func (r *MockLinkRepository) UpdateLinkStatus(ctx context.Context, params db.UpdateLinkStatusParams) error {
	args := r.Called(ctx, params)

	if rf, ok := args.Get(0).(func(context.Context, db.UpdateLinkStatusParams) error); ok {
		return rf(ctx, params)
	}
	return args.Error(0)
}

func (r *MockLinkRepository) UpdateLinkTrackingStatus(ctx context.Context, params db.UpdateLinkTrackingStatusParams) error {
	args := r.Called(ctx, params)

	if rf, ok := args.Get(0).(func(context.Context, db.UpdateLinkTrackingStatusParams) error); ok {
		return rf(ctx, params)
	}
	return args.Error(0)
}
