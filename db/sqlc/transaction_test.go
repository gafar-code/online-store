package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomTransaction(t *testing.T, customerID int64) Transaction {

	order := createRandomOrder(t)

	arg := CreateTransactionParams{
		CustomerID: customerID,
		Status:     "PENDING",
		IssuedAt:   time.Now(),
		OrderID:    order.ID,
	}

	trx, err := testQueries.CreateTransaction(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, trx)

	require.Equal(t, arg.CustomerID, trx.CustomerID)
	require.Equal(t, arg.Status, trx.Status)

	require.NotZero(t, trx.ID)
	require.NotZero(t, trx.IssuedAt)

	return trx
}

func TestCreateTransaction(t *testing.T) {
	cust := createRandomCustomer(t)
	createRandomTransaction(t, cust.ID)
}

func TestListTransaction(t *testing.T) {
	cust := createRandomCustomer(t)

	length := 10

	for i := 0; i < length; i++ {
		createRandomTransaction(t, cust.ID)
	}

	arg := ListTransactionByCustomerIdParams{
		CustomerID: cust.ID,
		Limit:      int32(length),
		Offset:     0,
	}

	transactions, err := testQueries.ListTransactionByCustomerId(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, transactions, length)

	for _, trx := range transactions {
		require.NotEmpty(t, trx)
		require.Equal(t, trx.CustomerID, arg.CustomerID)
		require.NotZero(t, trx.IssuedAt)
	}
}
