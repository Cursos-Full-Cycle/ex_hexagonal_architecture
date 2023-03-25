package application_test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"go_hexagonal/application"
	uuid "github.com/satori/go.uuid"
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

func TestProduct_Disabled(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Price = 0
	product.Status = application.ENABLED	
	err := product.Disable()	
	require.Nil(t, err)	
	
	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero to disable the product", err.Error())	
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Price = 10
	product.Status = application.DISABLED
	product.ID = uuid.NewV4().String()
	
	_, err := product.IsValid()
	require.Nil(t, err)
	
	product.Status = "invalid"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())
	
	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)
	
	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())
	
}

func TestProduct_GetId(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	require.Equal(t, product.ID, product.GetId())
}

func TestProduct_GetName(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	require.Equal(t, product.Name, product.GetName())
}

func TestProduct_GetStatus(t *testing.T) {
	product := application.Product{}
	product.Status = application.ENABLED
	require.Equal(t, product.Status, product.GetStatus())
}

func TestProduct_GetPrice(t *testing.T) {
	product := application.Product{}
	product.Price = 10
	require.Equal(t, product.Price, product.GetPrice())
	}	
	