package db

import (
	"context"
	"testing"
	"time"

	"github.com/gafar-code/online-store/util"
	"github.com/stretchr/testify/require"
)

func createRandomCustomer(t *testing.T) Customer {
	arg := CreateCustomerParams{
		Name:     util.RandomName(),
		Email:    util.RandomEmail(),
		Password: util.RandomPassword(),
		Address:  util.RandomString(40),
		Token:    util.RandomString(24),
	}

	customer, err := testQueries.CreateCustomer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, customer)

	require.Equal(t, arg.Name, customer.Name)
	require.Equal(t, arg.Email, customer.Email)
	require.Equal(t, arg.Password, customer.Password)
	require.Equal(t, arg.Address, customer.Address)
	require.Equal(t, arg.Token, customer.Token)

	require.NotZero(t, customer.ID)
	require.NotZero(t, customer.CreatedAt)

	return customer
}

func TestCreateCustomer(t *testing.T) {
	createRandomCustomer(t)
}

func TestGetCustomer(t *testing.T) {
	cust1 := createRandomCustomer(t)
	cust2, err := testQueries.GetCustomer(context.Background(), cust1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, cust2)

	require.Equal(t, cust1.ID, cust2.ID)
	require.Equal(t, cust1.Name, cust2.Name)
	require.Equal(t, cust1.Email, cust2.Email)
	require.Equal(t, cust1.Password, cust2.Password)
	require.Equal(t, cust1.Address, cust2.Address)
	require.Equal(t, cust1.Token, cust2.Token)
	require.WithinDuration(t, cust1.CreatedAt, cust2.CreatedAt, time.Second)
}
