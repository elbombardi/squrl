package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertNewClick(t *testing.T) {
	setup()
	defer teardown()

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

	err = testStore.InsertNewClick(ctx, InsertNewClickParams{
		LinkID:    1,
		UserAgent: sql.NullString{String: "Mozilla/5.0", Valid: true},
		IpAddress: sql.NullString{String: "127.0.0.1", Valid: true},
	})
	require.NoError(t, err, "Error should be nil")
}
