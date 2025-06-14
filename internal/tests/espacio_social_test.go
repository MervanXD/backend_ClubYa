package tests

import (
	"context"
	"testing"

	"github.com/MervanXD/backend_ClubYa/internal/models/espacio"
	"github.com/MervanXD/backend_ClubYa/internal/models/espacio/mocks"
	"github.com/stretchr/testify/assert"
)

func TestObtenerEspaciosSociales_OK(t *testing.T) {
	mockRepo := new(mocks.EspacioSocialRepository)
	mockEspacios := []espacio.EspacioSocial{
		{
			Espacio: espacio.Espacio{
				Id:     1,
				Nombre: "Salón Principal",
				Codigo: "ESP001",
			},
		},
	}
	mockRepo.On("ObtenerEspaciosSociales").Return(mockEspacios, nil)

	result, err := mockRepo.ObtenerEspaciosSociales()

	assert.NoError(t, err)
	assert.Equal(t, mockEspacios, result)
	mockRepo.AssertExpectations(t)
}

func TestObtenerEspaciosSociales_Error(t *testing.T) {
	mockRepo := new(mocks.EspacioSocialRepository)
	mockRepo.On("ObtenerEspaciosSociales").Return(nil, assert.AnError)

	result, err := mockRepo.ObtenerEspaciosSociales()

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestObtenerEspacioSocialPorID_OK(t *testing.T) {
	mockRepo := new(mocks.EspacioSocialRepository)
	expected := &espacio.EspacioSocial{
		Espacio: espacio.Espacio{
			Id:     1,
			Nombre: "Salón Principal",
			Codigo: "ESP001",
		},
	}
	ctx := context.Background()
	mockRepo.On("ObtenerEspacioSocialPorID", ctx, 1).Return(expected, nil)

	result, err := mockRepo.ObtenerEspacioSocialPorID(ctx, 1)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestObtenerEspacioSocialPorID_Error(t *testing.T) {
	mockRepo := new(mocks.EspacioSocialRepository)
	ctx := context.Background()
	mockRepo.On("ObtenerEspacioSocialPorID", ctx, 1).Return(nil, assert.AnError)

	result, err := mockRepo.ObtenerEspacioSocialPorID(ctx, 1)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
