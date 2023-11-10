package db

import (
	"context"
	"testing"

	"github.com/gafar-code/online-store/util"
	"github.com/stretchr/testify/require"
)

func createRandomVirtualAccount(t *testing.T) VirtualAccount {
	arg := CreateVirtualAccountParams{
		Name:           util.RandomName(),
		Description:    util.RandomString(40),
		RekeningNumber: util.RandomInt(1000000000, 9999999999),
	}

	va, err := testQueries.CreateVirtualAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, va)

	require.Equal(t, arg.Name, va.Name)
	require.Equal(t, arg.Name, va.Name)
	require.Equal(t, arg.Description, va.Description)
	require.Equal(t, arg.RekeningNumber, va.RekeningNumber)

	require.NotZero(t, va.ID)
	require.NotZero(t, va.CreatedAt)

	return va
}

func TestCreateVirtualAccount(t *testing.T) {
	createRandomVirtualAccount(t)
}

func TestListGetVirtualAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
	}
	acc1 := createRandomVirtualAccount(t)
	acc2, err := testQueries.GetVirtualAccount(context.Background(), acc1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, acc2)
	require.Equal(t, acc1.ID, acc2.ID)
	require.Equal(t, acc1.Name, acc2.Name)
	require.Equal(t, acc1.Description, acc2.Description)
	require.Equal(t, acc1.RekeningNumber, acc2.RekeningNumber)
}
func TestListVirtualAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomVirtualAccount(t)
	}

	accounts, err := testQueries.ListVirtualAccount(context.Background())

	require.NoError(t, err)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
