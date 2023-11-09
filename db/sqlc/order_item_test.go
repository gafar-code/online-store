package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomOrderItem(t *testing.T) OrderItem {
	prod := createRandomProduct(t)
	order := createRandomOrder(t)

	arg := CreateOrderItemParams{
		ProductID: prod.ID,
		Qty:       1,
		OrderID:   order.ID,
	}

	orderItem, err := testQueries.CreateOrderItem(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, orderItem)

	require.Equal(t, arg.ProductID, orderItem.ProductID)
	require.Equal(t, arg.Qty, orderItem.Qty)
	require.Equal(t, arg.OrderID, orderItem.OrderID)

	require.NotZero(t, orderItem.ID)
	require.NotZero(t, orderItem.CreatedAt)

	return orderItem
}

func TestCreateOrderItem(t *testing.T) {
	createRandomOrderItem(t)
}

func TestListOrderItemByOrderId(t *testing.T) {
	order := createRandomOrder(t)
	product := createRandomProduct(t)

	length := 10

	for i := 0; i < length; i++ {

		arg := CreateOrderItemParams{
			ProductID: product.ID,
			Qty:       1,
			OrderID:   order.ID,
		}

		order, err := testQueries.CreateOrderItem(context.Background(), arg)

		require.NoError(t, err)
		require.NotEmpty(t, order)
		require.Equal(t, arg.ProductID, order.ProductID)
	}

	arg := ListOrderItemByOrderIdParams{
		OrderID: order.ID,
		Limit:   10,
		Offset:  0,
	}

	orders, err := testQueries.ListOrderItemByOrderId(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, orders)
	require.Len(t, orders, length)

	for _, item := range orders {
		require.NotEmpty(t, order)
		require.Equal(t, item.OrderID, order.ID)
	}
}
