package tests

import (
	"testing"

	"github.com/MervanXD/backend_ClubYa/internal/models/espacio"
	"github.com/MervanXD/backend_ClubYa/internal/models/reserva"
	"github.com/MervanXD/backend_ClubYa/internal/models/reserva/mocks"
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
	"github.com/stretchr/testify/assert"
)

func TestReservarEspacio_OK(t *testing.T) {
	mockRepo := new(mocks.ReservaRepository)
	mockRepo.On("ReservarEspacio", 1, 2, 3).Return(nil)

	err := mockRepo.ReservarEspacio(1, 2, 3)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestReservarEspacio_Error(t *testing.T) {
	mockRepo := new(mocks.ReservaRepository)
	mockRepo.On("ReservarEspacio", 1, 2, 3).Return(assert.AnError)

	err := mockRepo.ReservarEspacio(1, 2, 3)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestReservarEspacioSocial_OK(t *testing.T) {
	mockRepo := new(mocks.ReservaRepository)
	reservaEjemplo := reserva.ReservaEspacio{IdSocio: 1}
	mockRepo.On("ReservarEspacioSocial", reservaEjemplo).Return(nil)

	err := mockRepo.ReservarEspacioSocial(reservaEjemplo)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestReservarEspacioSocial_Error(t *testing.T) {
	mockRepo := new(mocks.ReservaRepository)
	reservaEjemplo := reserva.ReservaEspacio{IdSocio: 1}
	mockRepo.On("ReservarEspacioSocial", reservaEjemplo).Return(assert.AnError)

	err := mockRepo.ReservarEspacioSocial(reservaEjemplo)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAnulacionReservaEspacioSocial_OK(t *testing.T) {
	mockRepo := new(mocks.ReservaRepository)
	mockRepo.On("AnulacionReservaEspacioSocial", 1, 2, 3, 4, "motivo").Return(nil)

	err := mockRepo.AnulacionReservaEspacioSocial(1, 2, 3, 4, "motivo")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAnulacionReservaEspacioSocial_Error(t *testing.T) {
	mockRepo := new(mocks.ReservaRepository)
	mockRepo.On("AnulacionReservaEspacioSocial", 1, 2, 3, 4, "motivo").Return(assert.AnError)

	err := mockRepo.AnulacionReservaEspacioSocial(1, 2, 3, 4, "motivo")

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestObtenerReservasEspaciosSocialesSocio_OK(t *testing.T) {
	mockRepo := new(mocks.ReservaRepository)
	expected := []reserva.ReservaEspacioSocialRequest{
		{
			Id:           1,
			Estado:       tipos.Estado(0),
			FechaReserva: "2023-10-01",
			Fecha:        "2023-10-01",
			HoraInicio:   "10:00",
			HoraFin:      "11:00",
			Espacio: espacio.EspacioSocial{
				Espacio: espacio.Espacio{
					Id:        1,
					Codigo:    "ESP001",
					Nombre:    "Espacio 1",
					Ubicacion: tipos.Ubicacion(0),
				},
			},
		},
		{
			Id:           2,
			Estado:       tipos.Estado(1),
			FechaReserva: "2023-10-02",
			Fecha:        "2023-10-02",
			HoraInicio:   "12:00",
			HoraFin:      "13:00",
			Espacio: espacio.EspacioSocial{
				Espacio: espacio.Espacio{
					Id:        1,
					Codigo:    "ESP001",
					Nombre:    "Espacio 1",
					Ubicacion: tipos.Ubicacion(1),
				},
			},
		},
	}
	mockRepo.On("ObtenerReservasEspaciosSocialesSocio", 1).Return(expected, nil)

	result, err := mockRepo.ObtenerReservasEspaciosSocialesSocio(1)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestObtenerReservasEspaciosSocialesSocio_Error(t *testing.T) {
	mockRepo := new(mocks.ReservaRepository)
	mockRepo.On("ObtenerReservasEspaciosSocialesSocio", 1).Return(nil, assert.AnError)

	result, err := mockRepo.ObtenerReservasEspaciosSocialesSocio(1)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestObtenerReservasCanchasSocio_OK(t *testing.T) {
	mockRepo := new(mocks.ReservaRepository)
	expected := []reserva.ReservaCanchaRequest{
		{
			Id:         1,
			Estado:     tipos.Estado(0),
			Fecha:      "2023-10-01",
			HoraInicio: "10:00",
			HoraFin:    "11:00",
			Espacio: espacio.Cancha{
				Espacio: espacio.Espacio{
					Id:        1,
					Codigo:    "CAN001",
					Nombre:    "Cancha 1",
					Ubicacion: tipos.Ubicacion(0),
				},
			},
		},
		{
			Id:         2,
			Estado:     tipos.Estado(1),
			Fecha:      "2023-10-02",
			HoraInicio: "12:00",
			HoraFin:    "13:00",
			Espacio:    espacio.Cancha{},
		},
	}
	mockRepo.On("ObtenerReservasCanchasSocio", 1).Return(expected, nil)

	result, err := mockRepo.ObtenerReservasCanchasSocio(1)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestObtenerReservasCanchasSocio_Error(t *testing.T) {
	mockRepo := new(mocks.ReservaRepository)
	mockRepo.On("ObtenerReservasCanchasSocio", 1).Return(nil, assert.AnError)

	result, err := mockRepo.ObtenerReservasCanchasSocio(1)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
