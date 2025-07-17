package tests

import (
	"context"
	"testing"
	"time"

	detalledisponibilidad "github.com/MervanXD/backend_ClubYa/internal/models/detalle_disponibilidad"
	"github.com/MervanXD/backend_ClubYa/internal/models/detalle_disponibilidad/mocks"
	"github.com/MervanXD/backend_ClubYa/internal/models/espacio"
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
	"github.com/stretchr/testify/assert"
)

func TestActualizarEstadoDetalleDisponibilidad_OK(t *testing.T) {
	mockRepo := new(mocks.DetalleDisponibilidadRepository)
	mockRepo.On("ActualizarEstadoDetalleDisponibilidad", 1, 2, "Disponible").Return(nil)

	err := mockRepo.ActualizarEstadoDetalleDisponibilidad(1, 2, "Disponible")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestActualizarEstadoDetalleDisponibilidad_Error(t *testing.T) {
	mockRepo := new(mocks.DetalleDisponibilidadRepository)
	mockRepo.On("ActualizarEstadoDetalleDisponibilidad", 1, 2, "Disponible").Return(assert.AnError)

	err := mockRepo.ActualizarEstadoDetalleDisponibilidad(1, 2, "Disponible")

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestObtenerDisponibilidadEspacioSocialPorId_OK(t *testing.T) {
	mockRepo := new(mocks.DetalleDisponibilidadRepository)
	ctx := context.Background()
	expected := &detalledisponibilidad.DisponibilidadEspacioResponse{
		Espacio: espacio.EspacioSocial{
			Espacio: espacio.Espacio{
				Id:        1,
				Nombre:    "Salón Principal",
				Codigo:    "ESP001",
				Ubicacion: tipos.Ubicacion(0),
				Capacidad: 100,
				Costo:     200.0,
			},
			Actividad: tipos.Actividad(0),
		},
		HoraInicio:     "10:00",
		HoraFin:        "12:00",
		Fecha:          time.Date(2024, 6, 12, 0, 0, 0, 0, time.UTC),
		Disponibilidad: tipos.EstadoDisponibilidad(1),
	}
	mockRepo.On("ObtenerDisponibilidadEspacioSocialPorId", ctx, 1, 2, 3).Return(expected, nil)

	result, err := mockRepo.ObtenerDisponibilidadEspacioSocialPorId(ctx, 1, 2, 3)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestObtenerDisponibilidadEspacioSocialPorId_Error(t *testing.T) {
	mockRepo := new(mocks.DetalleDisponibilidadRepository)
	ctx := context.Background()
	mockRepo.On("ObtenerDisponibilidadEspacioSocialPorId", ctx, 1, 2, 3).Return(nil, assert.AnError)

	result, err := mockRepo.ObtenerDisponibilidadEspacioSocialPorId(ctx, 1, 2, 3)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestObtenerRangosInicioDisponibles_OK(t *testing.T) {
	mockRepo := new(mocks.DetalleDisponibilidadRepository)
	expected := []string{"08:00:00", "10:00:00", "14:00:00"}

	mockRepo.On("ObtenerRangosInicioDisponibles", 5, "2025-06-20").Return(expected, nil)

	result, err := mockRepo.ObtenerRangosInicioDisponibles(5, "2025-06-20")

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestObtenerRangosInicioDisponibles_Error(t *testing.T) {
	mockRepo := new(mocks.DetalleDisponibilidadRepository)
	mockRepo.On("ObtenerRangosInicioDisponibles", 5, "2025-06-20").Return(nil, assert.AnError)

	result, err := mockRepo.ObtenerRangosInicioDisponibles(5, "2025-06-20")

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
