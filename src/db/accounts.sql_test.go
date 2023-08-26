package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// CheckPrefixExists(ctx context.Context, prefix string) (bool, error)
// CheckUsernameExists(ctx context.Context, username string) (bool, error)
// GetAccountByPrefix(ctx context.Context, prefix string) (Account, error)
// GetAccountByUsername(ctx context.Context, username string) (Account, error)
// InsertNewAccount(ctx context.Context, arg InsertNewAccountParams) error
// UpdateAccountStatusByUsername(ctx context.Context, arg UpdateAccountStatusByUsernameParams) error

func TestAccounts(t *testing.T) {

	ctx := context.Background()

	prefix := "tst"
	exists, err := testStore.CheckPrefixExists(ctx, prefix)
	require.NoError(t, err, "Error should be nil")
	require.False(t, exists, "Prefix should not exist")

	username := "test"
	exists, err = testStore.CheckUsernameExists(ctx, username)
	require.NoError(t, err, "Error should be nil")
	require.False(t, exists, "Username should not exist")

	account, err := testStore.GetAccountByPrefix(ctx, prefix)
	require.Error(t, err, "Error should not be nil")
	require.EqualError(t, err, "sql: no rows in result set", "Error should be sql: no rows in result set")
	require.Equal(t, Account{}, account, "Account should be empty")

	account, err = testStore.GetAccountByUsername(ctx, username)
	require.Error(t, err, "Error should not be nil")
	require.EqualError(t, err, "sql: no rows in result set", "Error should be sql: no rows in result set")
	require.Equal(t, Account{}, account, "Account should be empty")

	err = testStore.InsertNewAccount(ctx, InsertNewAccountParams{
		Prefix:         prefix,
		Username:       username,
		Email:          "email@gmail.com",
		HashedPassword: "$2a$1",
	})
	require.NoError(t, err, "Error should be nil")

	exists, err = testStore.CheckPrefixExists(ctx, prefix)
	require.NoError(t, err, "Error should be nil")
	require.True(t, exists, "Prefix should exist")

	exists, err = testStore.CheckUsernameExists(ctx, username)
	require.NoError(t, err, "Error should be nil")
	require.True(t, exists, "Username should exist")

	account, err = testStore.GetAccountByPrefix(ctx, prefix)
	require.NoError(t, err, "Error should be nil")
	require.Equal(t, int32(1), account.ID, "ID should be 1")
	require.Equal(t, prefix, account.Prefix, "Prefix should match")
	require.Equal(t, username, account.Username, "Username should match")
	require.Equal(t, true, account.Enabled, "Enabled should be true")

	exists, err = testStore.CheckPrefixExists(ctx, prefix)
	require.NoError(t, err, "Error should be nil")
	require.True(t, exists, "Prefix should exist")

	exists, err = testStore.CheckUsernameExists(ctx, username)
	require.NoError(t, err, "Error should be nil")
	require.True(t, exists, "Username should exist")

	account, err = testStore.GetAccountByPrefix(ctx, prefix)
	require.NoError(t, err, "Error should be nil")
	require.Equal(t, int32(1), account.ID, "ID should be 1")
	require.Equal(t, prefix, account.Prefix, "Prefix should match")
	require.Equal(t, username, account.Username, "Username should match")
	require.Equal(t, true, account.Enabled, "Enabled should be true")

	account, err = testStore.GetAccountByUsername(ctx, username)
	require.NoError(t, err, "Error should be nil")
	require.Equal(t, int32(1), account.ID, "ID should be 1")
	require.Equal(t, prefix, account.Prefix, "Prefix should match")
	require.Equal(t, username, account.Username, "Username should match")
	require.Equal(t, true, account.Enabled, "Enabled should be true")

	err = testStore.UpdateAccountStatusByUsername(ctx, UpdateAccountStatusByUsernameParams{
		Username: username,
		Enabled:  false,
	})
	require.NoError(t, err, "Error should be nil")

	account, err = testStore.GetAccountByUsername(ctx, username)
	require.NoError(t, err, "Error should be nil")
	require.Equal(t, int32(1), account.ID, "ID should be 1")
	require.Equal(t, prefix, account.Prefix, "Prefix should match")
	require.Equal(t, username, account.Username, "Username should match")
	require.Equal(t, false, account.Enabled, "Enabled should be false")

}
