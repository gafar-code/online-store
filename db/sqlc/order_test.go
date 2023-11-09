package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomOrder(t *testing.T) Order {
	cust := createRandomCustomer(t)
	va := createRandomVirtualAccount(t)

	arg := CreateOrderParams{
		CustomerID:       cust.ID,
		VirtualAccountID: va.ID,
	}

	order, err := testQueries.CreateOrder(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, order)

	require.Equal(t, arg.CustomerID, order.CustomerID)
	require.Equal(t, arg.VirtualAccountID, order.VirtualAccountID)

	require.NotZero(t, order.ID)
	require.NotZero(t, order.CreatedAt)

	return order
}

func TestCreateOrder(t *testing.T) {
	createRandomOrder(t)
}

func TestListOrderByCustomerId(t *testing.T) {
	cust := createRandomCustomer(t)

	length := 10

	for i := 0; i < length; i++ {
		va := createRandomVirtualAccount(t)

		arg := CreateOrderParams{
			CustomerID:       cust.ID,
			VirtualAccountID: va.ID,
		}

		order, err := testQueries.CreateOrder(context.Background(), arg)

		require.NoError(t, err)
		require.NotEmpty(t, order)
		require.Equal(t, arg.CustomerID, order.CustomerID)
	}

	arg := ListOrderByCustomerIdParams{
		CustomerID: cust.ID,
		Limit:      10,
		Offset:     0,
	}

	orders, err := testQueries.ListOrderByCustomerId(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, orders)
	require.Len(t, orders, length)

	for _, order := range orders {
		require.NotEmpty(t, order)
		require.Equal(t, order.CustomerID, cust.ID)
	}
}
