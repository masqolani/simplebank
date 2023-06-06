package db

import (
	"context"
	"github.com/masqolani/simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestQueriesCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestQueries_DeleteAccount(t *testing.T) {
	account := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)
}

func TestQueries_GetAccount(t *testing.T) {
	account := createRandomAccount(t)

	a, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, a)

	require.Equal(t, account.Owner, account.Owner)
	require.Equal(t, account.Balance, account.Balance)
	require.Equal(t, account.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestQueries_ListAccount(t *testing.T) {
	args := ListAccountParams{
		Limit:  0,
		Offset: 0,
	}

	_, err := testQueries.ListAccount(context.Background(), args)
	require.NoError(t, err)
}

func TestQueries_UpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	params := UpdateAccountParams{
		ID:      account.ID,
		Balance: 10000,
	}

	a, err := testQueries.UpdateAccount(context.Background(), params)
	require.NoError(t, err)
	require.Equal(t, account.ID, a.ID)
	require.Equal(t, params.Balance, a.Balance)
}
