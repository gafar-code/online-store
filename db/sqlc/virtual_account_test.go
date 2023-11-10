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

	customer, err := testQueries.CreateVirtualAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, customer)

	require.Equal(t, arg.Name, customer.Name)
	require.Equal(t, arg.Name, customer.Name)
	require.Equal(t, arg.Description, customer.Description)
	require.Equal(t, arg.RekeningNumber, customer.RekeningNumber)

	require.NotZero(t, customer.ID)
	require.NotZero(t, customer.CreatedAt)

	return customer
}

func TestCreateVirtualAccount(t *testing.T) {
	createRandomVirtualAccount(t)
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
