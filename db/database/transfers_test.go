package database

import (
	"context"
	"testing"

	"github.com/shama3541/simplebank/util"

	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)
	amount := util.RandomInt(0, account1.Balance)

	args := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        amount,
	}

	transfer_details, err := testQueries.CreateTransfer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfer_details)
	require.Equal(t, account1.ID, transfer_details.FromAccountID)
	require.Equal(t, account2.ID, transfer_details.ToAccountID)
	require.Equal(t, amount, transfer_details.Amount)

}
