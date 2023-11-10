package db

import (
	"context"
	"testing"
	"time"

	"github.com/gafar-code/online-store/util"
	"github.com/stretchr/testify/require"
)

func createRandomProduct(t *testing.T) Product {
	category := createRandomProductCategory(t)

	arg := CreateProductParams{
		Name:        util.RandomName(),
		ImageUrl:    util.RandomString(40),
		Description: util.RandomString(40),
		Price:       util.RandomInt(1000000000, 9999999999),
		CategoryID:  category.ID,
		Qty:         util.RandomInt(90, 1000),
	}

	product, err := testQueries.CreateProduct(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.Name, product.Name)
	require.Equal(t, arg.ImageUrl, product.ImageUrl)
	require.Equal(t, arg.Description, product.Description)
	require.Equal(t, arg.Price, product.Price)
	require.Equal(t, arg.CategoryID, product.CategoryID)
	require.Equal(t, arg.Qty, product.Qty)

	require.NotZero(t, product.ID)
	require.NotZero(t, product.CreatedAt)

	return product
}

func TestCreateProduct(t *testing.T) {
	createRandomProduct(t)
}

func TestGetProduct(t *testing.T) {
	prod1 := createRandomProduct(t)
	prod2, err := testQueries.GetProductDetail(context.Background(), prod1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, prod2)

	require.Equal(t, prod1.ID, prod2.ID)
	require.Equal(t, prod1.Name, prod2.Name)
	require.Equal(t, prod1.ImageUrl, prod2.ImageUrl)
	require.Equal(t, prod1.Description, prod2.Description)
	require.Equal(t, prod1.Price, prod2.Price)
	require.Equal(t, prod1.CategoryID, prod2.CategoryID)
	require.Equal(t, prod1.Qty, prod2.Qty)
	require.WithinDuration(t, prod1.CreatedAt, prod2.CreatedAt, time.Second)
}
func TestListProduct(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomProduct(t)
	}

	arg := ListProductParams{
		Limit:  10,
		Offset: 0,
	}
	prod, err := testQueries.ListProduct(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, prod)

	require.GreaterOrEqual(t, len(prod), 10)
}
func TestGetProductByCategoryId(t *testing.T) {
	category := createRandomProductCategory(t)

	arg := CreateProductParams{
		Name:        util.RandomName(),
		ImageUrl:    util.RandomString(40),
		Description: util.RandomString(40),
		Price:       util.RandomInt(1000000000, 9999999999),
		CategoryID:  category.ID,
		Qty:         util.RandomInt(90, 1000),
	}

	length := 10

	for i := 0; i < length; i++ {
		_, err := testQueries.CreateProduct(context.Background(), arg)

		require.NoError(t, err)
	}

	prodArg := GetProductByCategoryIdParams{
		CategoryID: category.ID,
		Limit:      int32(length),
		Offset:     0,
	}

	products, err := testQueries.GetProductByCategoryId(context.Background(), prodArg)

	require.NoError(t, err)
	require.NotEmpty(t, products)
	require.Len(t, products, length)

	for _, product := range products {
		require.NotEmpty(t, product)
		require.Equal(t, product.CategoryID, category.ID)
	}
}
