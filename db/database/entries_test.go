package database

import (
	"context"
	"testing"

	"github.com/shama3541/simplebank/util"

	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	account := CreateRandomAccount(t)
	amount := util.RandomInt(0, account.Balance)
	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    amount,
	}
	entry_details, err := testQueries.CreateEntry(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entry_details)
	require.Equal(t, entry_details.AccountID, account.ID)
	require.Equal(t, entry_details.Amount, args.Amount)
	require.NotZero(t, entry_details.ID)
	require.NotZero(t, entry_details.CreatedAt)

}
