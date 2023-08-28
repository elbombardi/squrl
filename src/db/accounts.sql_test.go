package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckPrefixExists(t *testing.T) {
	setup()
	defer teardown()
	ctx := context.Background()

	prefix := "tst"
	exists, err := testStore.CheckPrefixExists(ctx, prefix)
	require.NoError(t, err, "Error should be nil")
	require.False(t, exists, "Prefix should not exist")

	err = testStore.InsertNewAccount(ctx, InsertNewAccountParams{
		Prefix:         prefix,
		Username:       "username",
		Email:          "email@gmail.com",
		HashedPassword: "$2a$1",
	})
	require.NoError(t, err, "Error should be nil")

	exists, err = testStore.CheckPrefixExists(ctx, prefix)
	require.NoError(t, err, "Error should be nil")
	require.True(t, exists, "Prefix should exist")
}

func TestCheckUsernameExists(t *testing.T) {
	setup()
	defer teardown()
	ctx := context.Background()

	username := "username"
	exists, err := testStore.CheckUsernameExists(ctx, username)
	require.NoError(t, err, "Error should be nil")
	require.False(t, exists, "Username should not exist")

	err = testStore.InsertNewAccount(ctx, InsertNewAccountParams{
		Prefix:   "tst",
		Username: username,
		Email:    "email@gmail.com",
	})
	require.NoError(t, err, "Error should be nil")

	exists, err = testStore.CheckUsernameExists(ctx, username)
	require.NoError(t, err, "Error should be nil")
	require.True(t, exists, "Username should exist")

}

func TestGetAccountByPrefix(t *testing.T) {
	setup()
	defer teardown()
	ctx := context.Background()

	prefix := "tst"
	account, err := testStore.GetAccountByPrefix(ctx, prefix)
	require.Error(t, err, "Error should not be nil")
	require.Equal(t, sql.ErrNoRows, err, "Error should be sql.ErrNoRows")
	require.Equal(t, Account{}, account, "Account should be empty")

	err = testStore.InsertNewAccount(ctx, InsertNewAccountParams{
		Prefix:         prefix,
		Username:       "username",
		Email:          "email@gmail.com",
		HashedPassword: "$2a$1",
	})
	require.NoError(t, err, "Error should be nil")

	account, err = testStore.GetAccountByPrefix(ctx, prefix)
	require.NoError(t, err, "Error should be nil")
	require.Equal(t, int32(1), account.ID, "ID should be 1")
	require.Equal(t, prefix, account.Prefix, "Prefix should match")
	require.Equal(t, "username", account.Username, "Username should match")
	require.Equal(t, true, account.Enabled, "Enabled should be true")

}

func TestGetAccountByUsername(t *testing.T) {
	setup()
	defer teardown()
	ctx := context.Background()

	username := "username"
	account, err := testStore.GetAccountByUsername(ctx, username)
	require.Error(t, err, "Error should not be nil")
	require.Equal(t, sql.ErrNoRows, err, "Error should be sql.ErrNoRows")
	require.Equal(t, Account{}, account, "Account should be empty")

	err = testStore.InsertNewAccount(ctx, InsertNewAccountParams{
		Prefix:         "tst",
		Username:       username,
		Email:          "email@gmail.com",
		HashedPassword: "$2a$1",
	})
	require.NoError(t, err, "Error should be nil")

	account, err = testStore.GetAccountByUsername(ctx, username)
	require.NoError(t, err, "Error should be nil")
	require.Equal(t, int32(1), account.ID, "ID should be 1")
	require.Equal(t, "tst", account.Prefix, "Prefix should match")
	require.Equal(t, username, account.Username, "Username should match")
	require.Equal(t, true, account.Enabled, "Enabled should be true")

}

func TestUpdateAccountStatusByUsername(t *testing.T) {
	setup()
	defer teardown()
	ctx := context.Background()

	username := "username"
	account, err := testStore.GetAccountByUsername(ctx, username)
	require.Error(t, err, "Error should not be nil")
	require.Equal(t, sql.ErrNoRows, err, "Error should be sql.ErrNoRows")
	require.Equal(t, Account{}, account, "Account should be empty")

	err = testStore.InsertNewAccount(ctx, InsertNewAccountParams{
		Prefix:         "tdst",
		Username:       username,
		Email:          "email@gmail.com",
		HashedPassword: "$2a$1",
	})
	require.NoError(t, err, "Error should be nil")

	err = testStore.UpdateAccountStatusByUsername(ctx, UpdateAccountStatusByUsernameParams{
		Username: username,
		Enabled:  false,
	})
	require.NoError(t, err, "Error should be nil")

	account, err = testStore.GetAccountByUsername(ctx, username)
	require.NoError(t, err, "Error should be nil")
	require.Equal(t, int32(1), account.ID, "ID should be 1")
	require.Equal(t, "tst", account.Prefix, "Prefix should match")
	require.Equal(t, username, account.Username, "Username should match")
	require.Equal(t, false, account.Enabled, "Enabled should be false")

}
