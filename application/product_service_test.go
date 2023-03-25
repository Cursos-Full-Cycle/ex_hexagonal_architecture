package application_test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"go_hexagonal/application"
	"github.com/golang/mock/gomock"
	mock_application "go_hexagonal/application/mocks"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()
	
	service := application.ProductService{
		Persistence: persistence,
	}
	
	result, err := service.Get("123")
	require.Nil(t, err)
	require.Equal(t, product, result)	
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
	
	service := application.ProductService { Persistence: persistence }
	
	result, err := service.Create("product 1", 10)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_EnableDisable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	product := mock_application.NewMockProductInterface(ctrl)
	product.EXPECT().Enable().Return(nil).AnyTimes()
	product.EXPECT().Disable().Return(nil).AnyTimes()
	
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
	
	service := application.ProductService { Persistence: persistence }
	
	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
	
	result, err = service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}