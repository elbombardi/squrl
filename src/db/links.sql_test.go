package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

// (ctx context.Context, arg UpdateLinkTrackingStatusParams) error

func TestCheckShortUrlKeyExists(t *testing.T) {
	ctx := context.Background()

	exists, err := testStore.CheckShortUrlKeyExists(ctx, CheckShortUrlKeyExistsParams{
		AccountID:   1,
		ShortUrlKey: sql.NullString{String: "test", Valid: true},
	})
	require.NoError(t, err, "Error should be nil")
	require.False(t, exists, "ShortUrlKey should not exist")

	err = testStore.InsertNewAccount(ctx, InsertNewAccountParams{
		Prefix:         "tst",
		Username:       "username",
		Email:          "email@gmail.com",
		HashedPassword: "$2a$1",
	})
	require.NoError(t, err, "Error should be nil")

	err = testStore.InsertNewLink(ctx, InsertNewLinkParams{
		AccountID:   1,
		ShortUrlKey: sql.NullString{String: "test", Valid: true},
		LongUrl:     "https://www.google.com",
	})
	require.NoError(t, err, "Error should be nil")

	exists, err = testStore.CheckShortUrlKeyExists(ctx, CheckShortUrlKeyExistsParams{
		AccountID:   1,
		ShortUrlKey: sql.NullString{String: "test", Valid: true},
	})
	require.NoError(t, err, "Error should be nil")
	require.True(t, exists, "ShortUrlKey should exist")
}

func TestGetLinkByAccountIDAndShortURLKey(t *testing.T) {
	ctx := context.Background()

	link, err := testStore.GetLinkByAccountIDAndShortURLKey(ctx, GetLinkByAccountIDAndShortURLKeyParams{
		AccountID:   1,
		ShortUrlKey: sql.NullString{String: "test", Valid: true},
	})
	require.Error(t, err, "Error should not be nil")
	require.EqualError(t, err, "sql: no rows in result set", "Error should be sql: no rows in result set")
	require.Equal(t, Link{}, link, "Link should be empty")

	err = testStore.InsertNewAccount(ctx, InsertNewAccountParams{
		Prefix:         "tst",
		Username:       "username",
		Email:          "email@gmail.com",
		HashedPassword: "$2a$1",
	})
	require.NoError(t, err, "Error should be nil")

	err = testStore.InsertNewLink(ctx, InsertNewLinkParams{
		AccountID:   1,
		ShortUrlKey: sql.NullString{String: "test", Valid: true},
		LongUrl:     "https://www.google.com",
	})
	require.NoError(t, err, "Error should be nil")

	link, err = testStore.GetLinkByAccountIDAndShortURLKey(ctx, GetLinkByAccountIDAndShortURLKeyParams{
		AccountID:   1,
		ShortUrlKey: sql.NullString{String: "test", Valid: true},
	})
	require.NoError(t, err, "Error should be nil")
	require.Equal(t, int32(1), link.ID, "ID should be 1")
	require.Equal(t, int32(1), link.AccountID, "AccountID should be 1")
	require.Equal(t, "test", link.ShortUrlKey.String, "ShortUrlKey should match")
	require.Equal(t, "https://www.google.com", link.LongUrl, "LongUrl should match")
	require.Equal(t, true, link.Enabled, "Enabled should be true")
	require.Equal(t, true, link.TrackingEnabled, "TrackingEnabled should be true")

}

func TestUpdateLinkLongURL(t *testing.T) {
	ctx := context.Background()

	err := testStore.InsertNewAccount(ctx, InsertNewAccountParams{
		Prefix:         "tst",
		Username:       "username",
		Email:          "email@gmail.com",
		HashedPassword: "$2a$1",
	})
	require.NoError(t, err, "Error should be nil")

	err = testStore.InsertNewLink(ctx, InsertNewLinkParams{
		AccountID:   1,
		ShortUrlKey: sql.NullString{String: "test", Valid: true},
		LongUrl:     "https://www.google.com",
	})
	require.NoError(t, err, "Error should be nil")

	err = testStore.UpdateLinkLongURL(ctx, UpdateLinkLongURLParams{
		LongUrl: "https://www.yahoo.com",
		ID:      1,
	})
	require.NoError(t, err, "Error should be nil")

	link, err := testStore.GetLinkByAccountIDAndShortURLKey(ctx, GetLinkByAccountIDAndShortURLKeyParams{
		AccountID:   1,
		ShortUrlKey: sql.NullString{String: "test", Valid: true},
	})
	require.NoError(t, err, "Error should be nil")
	require.Equal(t, int32(1), link.ID, "ID should be 1")
	require.Equal(t, int32(1), link.AccountID, "AccountID should be 1")
	require.Equal(t, "test", link.ShortUrlKey.String, "ShortUrlKey should match")
	require.Equal(t, "https://www.yahoo.com", link.LongUrl, "LongUrl should match")
}

func TestUpdateLinkTrackingStatus(t *testing.T) {
	ctx := context.Background()

	err := testStore.InsertNewAccount(ctx, InsertNewAccountParams{
		Prefix:         "tst",
		Username:       "username",
		Email:          "email@gmail.com",
		HashedPassword: "$2a$1",
	})
	require.NoError(t, err, "Error should be nil")

	err = testStore.InsertNewLink(ctx, InsertNewLinkParams{
		AccountID:   1,
		ShortUrlKey: sql.NullString{String: "test", Valid: true},
		LongUrl:     "https://www.google.com",
	})
	require.NoError(t, err, "Error should be nil")

	err = testStore.UpdateLinkTrackingStatus(ctx, UpdateLinkTrackingStatusParams{
		ID:              1,
		TrackingEnabled: false,
	})
	require.NoError(t, err, "Error should be nil")

	link, err := testStore.GetLinkByAccountIDAndShortURLKey(ctx, GetLinkByAccountIDAndShortURLKeyParams{
		AccountID:   1,
		ShortUrlKey: sql.NullString{String: "test", Valid: true},
	})
	require.NoError(t, err, "Error should be nil")
	require.Equal(t, int32(1), link.ID, "ID should be 1")
	require.Equal(t, int32(1), link.AccountID, "AccountID should be 1")
	require.Equal(t, "test", link.ShortUrlKey.String, "ShortUrlKey should match")
	require.Equal(t, false, link.TrackingEnabled, "TrackingEnabled should be false")
}

func TestUpdateLinkStatus(t *testing.T) {
	ctx := context.Background()

	err := testStore.InsertNewAccount(ctx, InsertNewAccountParams{
		Prefix:         "tst",
		Username:       "username",
		Email:          "email@gmail.com",
		HashedPassword: "$2a$1",
	})
	require.NoError(t, err, "Error should be nil")

	err = testStore.InsertNewLink(ctx, InsertNewLinkParams{
		AccountID:   1,
		ShortUrlKey: sql.NullString{String: "test", Valid: true},
		LongUrl:     "https://www.google.com",
	})
	require.NoError(t, err, "Error should be nil")

	err = testStore.UpdateLinkStatus(ctx, UpdateLinkStatusParams{
		ID:      1,
		Enabled: false,
	})
	require.NoError(t, err, "Error should be nil")

	link, err := testStore.GetLinkByAccountIDAndShortURLKey(ctx, GetLinkByAccountIDAndShortURLKeyParams{
		AccountID:   1,
		ShortUrlKey: sql.NullString{String: "test", Valid: true},
	})
	require.NoError(t, err, "Error should be nil")
	require.Equal(t, int32(1), link.ID, "ID should be 1")
	require.Equal(t, int32(1), link.AccountID, "AccountID should be 1")
	require.Equal(t, "test", link.ShortUrlKey.String, "ShortUrlKey should match")
	require.Equal(t, false, link.Enabled, "Enabled should be false")
}
