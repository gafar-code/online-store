package db

import (
	"context"
	"testing"
	"time"

	"github.com/gafar-code/online-store/util"
	"github.com/stretchr/testify/require"
)

func createRandomProof(t *testing.T) Order {
	// create account
	arg := CreateCustomerParams{
		Name:     util.RandomName(),
		Email:    util.RandomEmail(),
		Password: util.RandomPassword(),
		Address:  util.RandomString(40),
	}

	customer, err := testQueries.CreateCustomer(context.Background(), arg)
	require.NoError(t, err)

	// create product category
	name := util.RandomName()
	category, err := testQueries.CreateProductCategory(context.Background(), name)
	require.NoError(t, err)

	// create product
	prodArg := CreateProductParams{
		Name:        util.RandomName(),
		ImageUrl:    util.RandomString(40),
		Description: util.RandomString(40),
		Price:       util.RandomInt(1000000000, 9999999999),
		CategoryID:  category.ID,
		Qty:         util.RandomInt(90, 1000),
	}

	product, err := testQueries.CreateProduct(context.Background(), prodArg)
	require.NoError(t, err)

	// create virtual account
	vaArg := CreateVirtualAccountParams{
		Name:           util.RandomName(),
		Description:    util.RandomString(40),
		RekeningNumber: util.RandomInt(1000000000, 9999999999),
	}

	va, err := testQueries.CreateVirtualAccount(context.Background(), vaArg)
	require.NoError(t, err)

	// create order
	orderArg := CreateOrderParams{
		CustomerID:       customer.ID,
		Amount:           util.RandomInt(1000000000, 9999999999),
		Status:           "WAITING_PAYMENT",
		VirtualAccountID: va.ID,
		ExpiredAt:        time.Now().Add(time.Hour * 24),
	}

	order, err := testQueries.CreateOrder(context.Background(), orderArg)
	require.NoError(t, err)

	// create order item
	orderItemArg := CreateOrderItemParams{
		CategoryID:  category.ID,
		Name:        product.Name,
		ImageUrl:    product.ImageUrl,
		Description: product.Description,
		Price:       product.Price,
		Qty:         10,
		ProductID:   product.ID,
		OrderID:     order.ID,
	}

	_, err = testQueries.CreateOrderItem(context.Background(), orderItemArg)
	require.NoError(t, err)

	// order prove
	crtOrderArg := CreateOrderProofParams{
		OrderID:        order.ID,
		NameHolder:     customer.Name,
		RekeningNumber: va.RekeningNumber,
		ImageUrl:       "https://google.com",
	}

	newOrder, err := testQueries.CreateOrderProof(context.Background(), crtOrderArg)
	require.NoError(t, err)
	require.Equal(t, crtOrderArg.OrderID, order.ID)

	// approove order
	approveArg := UpdateOrderParams{
		ID:     newOrder.ID,
		Status: "PAID",
	}

	uptOrder, err := testQueries.UpdateOrder(context.Background(), approveArg)
	require.NoError(t, err)
	require.Equal(t, approveArg.ID, uptOrder.ID)
	require.Equal(t, approveArg.Status, uptOrder.Status)

	trxArg := CreateTransactionParams{
		CustomerID: order.CustomerID,
		Status:     "ON_PROGRESS",
		IssuedAt:   order.IssuedAt,
		OrderID:    order.ID,
	}

	trx, err := testQueries.CreateTransaction(context.Background(), trxArg)

	require.NoError(t, err)
	require.Equal(t, trx.CustomerID, trxArg.CustomerID)
	require.Equal(t, trx.OrderID, trxArg.OrderID)
	require.Equal(t, trx.Status, trxArg.Status)

	return order
}

func TestCreateProofOrder(t *testing.T) {
	createRandomProof(t)
}

func TestGetProofOrder(t *testing.T) {
	order := createRandomProof(t)

	proof, err := testQueries.GetOrderProof(context.Background(), int32(order.ID))

	require.NoError(t, err)
	require.NotEmpty(t, proof)
	require.Equal(t, order.ID, proof.OrderID)
}
