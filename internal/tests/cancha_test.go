package tests

import (
	"testing"

	"github.com/MervanXD/backend_ClubYa/internal/models/espacio"
	"github.com/MervanXD/backend_ClubYa/internal/models/espacio/mocks"
	"github.com/stretchr/testify/assert"
)

func TestObtenerCanchasHorarios_OK(t *testing.T) {
	mockRepo := new(mocks.CanchaRepository)
	mockCanchas := []espacio.CanchaHorarioDTO{
		{
			Espacio: espacio.Cancha{
				Espacio: espacio.Espacio{
					Id:     1,
					Nombre: "Cancha 1",
				},
			},
			Fecha:      "2024-06-12",
			HoraInicio: "10:00",
			HoraFinal:  "11:00",
		},
	}
	mockRepo.On("ObtenerCanchasHorarios").Return(mockCanchas, nil)

	result, err := mockRepo.ObtenerCanchasHorarios()

	assert.NoError(t, err)
	assert.Equal(t, mockCanchas, result)
	mockRepo.AssertExpectations(t)
}

func TestObtenerCanchasHorarios_Error(t *testing.T) {
	mockRepo := new(mocks.CanchaRepository)
	mockRepo.On("ObtenerCanchasHorarios").Return(nil, assert.AnError)

	result, err := mockRepo.ObtenerCanchasHorarios()

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
