package db

import (
	"context"
	"testing"

	"github.com/gafar-code/online-store/util"
	"github.com/stretchr/testify/require"
)

func createRandomProductCategory(t *testing.T) ProductCategory {

	name := util.RandomName()
	category, err := testQueries.CreateProductCategory(context.Background(), name)

	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, name, category.Name)

	require.NotZero(t, category.ID)

	return category
}

func TestCreateProductCategory(t *testing.T) {
	createRandomProductCategory(t)
}

func TestListProductCategory(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomProductCategory(t)
	}

	arg := ListProductCategoryParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListProductCategory(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
