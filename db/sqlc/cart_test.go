package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/gafar-code/online-store/util"
	"github.com/stretchr/testify/require"
)

func createRandomCart(t *testing.T) Cart {
	cust := createRandomCustomer(t)
	prod := createRandomProduct(t)

	arg := CreateCartParams{
		CustomerID: cust.ID,
		ProductID:  prod.ID,
		Qty:        1,
	}

	cart, err := testQueries.CreateCart(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, cart)

	require.Equal(t, arg.CustomerID, cart.CustomerID)
	require.Equal(t, arg.ProductID, cart.ProductID)
	require.Equal(t, int64(1), cart.Qty)

	require.NotZero(t, cart.ID)
	require.NotZero(t, cart.CreatedAt)

	return cart
}

func TestCreateCart(t *testing.T) {
	createRandomCart(t)
}

func TestGetCartByCustomerId(t *testing.T) {
	length := 10
	cust1 := createRandomCustomer(t)

	for i := 0; i < length; i++ {

		product := createRandomProduct(t)
		arg := CreateCartParams{
			CustomerID: cust1.ID,
			ProductID:  product.ID,
		}
		_, err := testQueries.CreateCart(context.Background(), arg)
		require.NoError(t, err)
	}

	arg := GetCartByCustomerIdParams{
		CustomerID: cust1.ID,
		Limit:      int32(length),
		Offset:     0,
	}

	carts, err := testQueries.GetCartByCustomerId(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, carts)
	require.Len(t, carts, length)

	for _, cart := range carts {
		require.NotEmpty(t, cart)
		require.Equal(t, cart.CustomerID, cust1.ID)
	}
}

func TestGetCart(t *testing.T) {
	cart1 := createRandomCart(t)
	cart2, err := testQueries.GetCart(context.Background(), cart1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, cart2)

	require.Equal(t, cart1.CustomerID, cart2.CustomerID)
	require.Equal(t, cart1.ProductID, cart2.ProductID)
	require.Equal(t, cart1.Qty, cart2.Qty)
}

func TestDeleteCart(t *testing.T) {
	cart := createRandomCart(t)
	err := testQueries.DeleteCart(context.Background(), cart.ID)

	require.NoError(t, err)
	found, err := testQueries.GetCart(context.Background(), cart.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, found)
}

func TestUpdateCart(t *testing.T) {
	cart := createRandomCart(t)

	arg := UpdateCartParams{
		ID:  cart.ID,
		Qty: util.RandomInt(1, 10),
	}

	newCart, err := testQueries.UpdateCart(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, newCart)

	require.Equal(t, cart.ID, newCart.ID)
	require.Equal(t, cart.CustomerID, newCart.CustomerID)
	require.Equal(t, cart.ProductID, newCart.ProductID)
	require.Equal(t, arg.Qty, newCart.Qty)
}
