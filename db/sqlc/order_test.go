package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomOrder(t *testing.T) Order {
	cust := createRandomCustomer(t)
	va := createRandomVirtualAccount(t)

	arg := CreateOrderParams{
		Status:           "WAITING_PAYMENT",
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

func TestGetOrder(t *testing.T) {
	order1 := createRandomOrder(t)

	order2, err := testQueries.GetOrder(context.Background(), order1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, order2)

	require.Equal(t, order1.ID, order2.ID)
	require.Equal(t, order1.CustomerID, order2.CustomerID)
	require.Equal(t, order1.VirtualAccountID, order2.VirtualAccountID)
	require.Equal(t, order1.Status, order2.Status)
	require.Equal(t, order1.CreatedAt, order2.CreatedAt)
	require.WithinDuration(t, order1.CreatedAt, order2.CreatedAt, time.Second)

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
