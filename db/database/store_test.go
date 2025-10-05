package database

import (
	"context"
	"fmt"
	"testing"

	// "github.com/shama3541/simplebank/util"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDb)
	amount := 100
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)
	fmt.Println(">>before:", account1.Balance, account2.Balance)
	errs := make(chan error)
	results := make(chan TransferTxResult)
	n := 5
	for i := 0; i < n; i++ {

		go func() {
			result, err := store.TranferTx(context.Background(), TransferParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        int64(amount),
			})
			errs <- err
			results <- result
		}()
	}
	existed := make(map[int]bool)
	for i := 0; i < 5; i++ {

		err := <-errs
		require.NoError(t, err)
		result := <-results
		require.NotEmpty(t, result)

		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.NotZero(t, transfer.CreatedAt)
		require.Equal(t, transfer.Amount, int64(amount))
		require.NotZero(t, transfer.ID)

		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, fromEntry.AccountID, account1.ID)
		require.Equal(t, -int64(amount), fromEntry.Amount)
		require.NotZero(t, fromEntry.CreatedAt)
		require.NotZero(t, fromEntry.ID)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, toEntry.AccountID, account2.ID)
		require.Equal(t, int64(amount), toEntry.Amount)
		require.NotZero(t, toEntry.CreatedAt)
		require.NotZero(t, toEntry.ID)

		FromAccount := result.FromAccount
		require.NotEmpty(t, FromAccount)
		require.Equal(t, FromAccount.ID, account1.ID)

		ToAccount := result.ToAccount
		require.NotEmpty(t, ToAccount)
		require.Equal(t, ToAccount.ID, account2.ID)
		fmt.Println(">> tx:", FromAccount.Balance, ToAccount.Balance)
		diff1 := account1.Balance - FromAccount.Balance
		diff2 := ToAccount.Balance - account2.Balance

		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%int64(amount) == 0)

		k := int(diff1 / int64(amount))
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true

	}

	_, err := store.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	_, err = store.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

}

func TestDeadlockTxs(t *testing.T) {
	store := NewStore(testDb)

	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)
	amount := int64(10)
	errchan := make(chan error)
	resultchan := make(chan TransferTxResult)

	for i := 1; i <= 5; i++ {
		ToAccountID := account1.ID
		FromAccountId := account2.ID
		if i%2 == 0 {
			ToAccountID = account2.ID
			FromAccountId = account1.ID
		}

		go func() {
			result, err := store.TranferTx(context.Background(), TransferParams{
				FromAccountID: FromAccountId,
				ToAccountID:   ToAccountID,
				Amount:        amount,
			})

			errchan <- err
			resultchan <- result
		}()

	}

	for i := 1; i <= 5; i++ {
		err := <-errchan
		result := <-resultchan

		require.NotEmpty(t, result)
		require.NoError(t, err)
	}

}
