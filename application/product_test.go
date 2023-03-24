package application_test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"go_hexagonal/application"
)

func TestProduct_Enabled(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Price = 10
	product.Status = application.DISABLED	
	err := product.Enable()	
	require.Nil(t, err)	
	
	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())	
}
