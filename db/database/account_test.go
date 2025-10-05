package database

import (
	"context"
	"database/sql"
	"testing"

	"github.com/shama3541/simplebank/util"

	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    util.RandomName(),
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

func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func TesGetAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.CreatedAt, account2.CreatedAt)
	require.Equal(t, account1.Currency, account2.Currency)

}

func TestUpdateAccount(t *testing.T) {
	account := CreateRandomAccount(t)
	args := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}
	_, err := testQueries.UpdateAccount(context.Background(), args)
	require.NoError(t, err)
	updatedAccount, err := testQueries.GetAccount(context.Background(), args.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)
	require.Equal(t, args.Balance, updatedAccount.Balance)
	require.Equal(t, args.ID, updatedAccount.ID)

}

func TestDeleteAccount(t *testing.T) {
	account := CreateRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)
	account2, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err, sql.ErrNoRows)
	require.Empty(t, account2)

}

func TestListAccount(t *testing.T) {
	for i := 0; i < 5; i++ {
		CreateRandomAccount(t)
	}

	args := ListAccountParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListAccount(context.Background(), args)
	require.NoError(t, err)
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
