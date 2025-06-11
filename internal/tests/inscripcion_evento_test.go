package tests

import (
	"testing"
	"time"

	"github.com/MervanXD/backend_ClubYa/internal/models/evento"
	"github.com/MervanXD/backend_ClubYa/internal/models/inscripcion_evento"
	"github.com/MervanXD/backend_ClubYa/internal/models/inscripcion_evento/mocks"
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
	"github.com/stretchr/testify/assert"
)

func TestRegistrarInscripcion_OK(t *testing.T) {
	mockRepo := new(mocks.InscripcionEventoRepository)
	mockRepo.On("RegistrarInscripcion", 1, 2, 3).Return(nil)

	err := mockRepo.RegistrarInscripcion(1, 2, 3)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRegistrarInscripcion_Error(t *testing.T) {
	mockRepo := new(mocks.InscripcionEventoRepository)
	mockRepo.On("RegistrarInscripcion", 1, 2, 3).Return(assert.AnError)

	err := mockRepo.RegistrarInscripcion(1, 2, 3)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestObtenerEventosSocio_OK(t *testing.T) {
	mockRepo := new(mocks.InscripcionEventoRepository)
	expected := []inscripcion_evento.InscripcionSocioDTO{
		{
			IdInscripcionEvento: 1,
			FechaInscripcion:    "2024-01-01",
			Estado:              tipos.Estado(0),
			Evento: evento.Evento{
				IdEvento:    10,
				Nombre:      "Torneo",
				Descripcion: "Torneo anual",
				Fecha:       time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
				Precio:      50.0,
				HoraInicio:  time.Date(0, 1, 1, 10, 0, 0, 0, time.UTC),
				HoraFin:     time.Date(0, 1, 1, 12, 0, 0, 0, time.UTC),
			},
		},
	}
	mockRepo.On("ObtenerEventosSocio", 1).Return(expected, nil)

	result, err := mockRepo.ObtenerEventosSocio(1)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestObtenerEventosSocio_Error(t *testing.T) {
	mockRepo := new(mocks.InscripcionEventoRepository)
	mockRepo.On("ObtenerEventosSocio", 1).Return(nil, assert.AnError)

	result, err := mockRepo.ObtenerEventosSocio(1)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
