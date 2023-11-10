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
		CategoryID:  prod.CategoryID,
		Name:        prod.Name,
		ImageUrl:    prod.ImageUrl,
		Description: prod.Description,
		Price:       prod.Price,
		Qty:         1,
		ProductID:   prod.ID,
		OrderID:     order.ID,
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
			CategoryID:  product.CategoryID,
			Name:        product.Name,
			ImageUrl:    product.ImageUrl,
			Description: product.Description,
			Price:       product.Price,
			Qty:         1,
			ProductID:   product.ID,
			OrderID:     order.ID,
		}

		order, err := testQueries.CreateOrderItem(context.Background(), arg)

		require.NoError(t, err)
		require.NotEmpty(t, order)
		require.Equal(t, arg.ProductID, order.ProductID)
	}

	orders, err := testQueries.ListOrderItemByOrderId(context.Background(), order.ID)

	require.NoError(t, err)
	require.NotEmpty(t, orders)
	require.Len(t, orders, length)

	for _, item := range orders {
		require.NotEmpty(t, order)
		require.Equal(t, item.OrderID, order.ID)
	}
}
