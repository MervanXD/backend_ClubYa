package tests

import (
	"errors"
	"testing"
	"time"

	"github.com/MervanXD/backend_ClubYa/internal/models/evento"
	"github.com/MervanXD/backend_ClubYa/internal/models/evento/mocks"
	"github.com/stretchr/testify/assert"
)

func TestListarEventos_OK(t *testing.T) {
	mockRepo := new(mocks.EventoRepository)
	expected := []evento.Evento{
		{
			IdEvento:     1,
			Nombre:       "Evento de prueba",
			Descripcion:  "Descripción de prueba",
			Fecha:        time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC),
			HoraInicio:   time.Date(0, 1, 1, 10, 0, 0, 0, time.UTC),
			HoraFin:      time.Date(0, 1, 1, 12, 0, 0, 0, time.UTC),
			Aforo:        100,
			Invitados:    10,
			Precio:       20.0,
			NroInscritos: 30,
		},
	}
	mockRepo.On("ListarEventos").Return(expected, nil)

	result, err := mockRepo.ListarEventos()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestListarEventos_Error(t *testing.T) {
	mockRepo := new(mocks.EventoRepository)
	mockRepo.On("ListarEventos").Return(nil, errors.New("db error"))

	result, err := mockRepo.ListarEventos()

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestBuscarEventoPorID_OK(t *testing.T) {
	mockRepo := new(mocks.EventoRepository)
	expected := &evento.Evento{
		IdEvento:     1,
		Nombre:       "Evento X",
		Descripcion:  "Descripción",
		Fecha:        time.Date(2025, 6, 15, 0, 0, 0, 0, time.UTC),
		HoraInicio:   time.Date(0, 1, 1, 9, 0, 0, 0, time.UTC),
		HoraFin:      time.Date(0, 1, 1, 11, 0, 0, 0, time.UTC),
		Aforo:        50,
		Invitados:    5,
		Precio:       10.0,
		NroInscritos: 25,
		Estado:       1,
	}
	mockRepo.On("BuscarEventoPorID", 1).Return(expected, nil)

	result, err := mockRepo.BuscarEventoPorID(1)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestBuscarEventoPorID_Error(t *testing.T) {
	mockRepo := new(mocks.EventoRepository)
	mockRepo.On("BuscarEventoPorID", 1).Return(nil, errors.New("not found"))

	result, err := mockRepo.BuscarEventoPorID(1)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestInsertarEvento_OK(t *testing.T) {
	mockRepo := new(mocks.EventoRepository)
	req := evento.EventoRequest{
		Nombre:      "Nuevo Evento",
		Descripcion: "Descripción",
		Fecha:       "2025-06-30",
		HoraInicio:  "08:00:00",
		HoraFin:     "10:00:00",
		Aforo:       100,
		Invitados:   20,
		Precio:      15.0,
		IdEspacio:   2,
	}
	mockRepo.On("InsertarEvento", req).Return(5, nil)

	id, err := mockRepo.InsertarEvento(req)

	assert.NoError(t, err)
	assert.Equal(t, 5, id)
	mockRepo.AssertExpectations(t)
}

func TestInsertarEvento_Error(t *testing.T) {
	mockRepo := new(mocks.EventoRepository)
	req := evento.EventoRequest{
		Nombre:      "Evento fallido",
		Descripcion: "X",
		Fecha:       "2025-13-40", // fecha inválida
		HoraInicio:  "25:00:00",   // hora inválida
		HoraFin:     "xx:yy:zz",
	}
	mockRepo.On("InsertarEvento", req).Return(-1, errors.New("formato inválido"))

	id, err := mockRepo.InsertarEvento(req)

	assert.Error(t, err)
	assert.Equal(t, -1, id)
	mockRepo.AssertExpectations(t)
}

func TestModificarEvento_OK(t *testing.T) {
	mockRepo := new(mocks.EventoRepository)
	req := evento.EventoRequest{
		IdEvento:    1,
		Nombre:      "Actualizado",
		Descripcion: "Evento actualizado",
		Aforo:       120,
		Invitados:   15,
		Precio:      25.5,
		Estado:      1,
	}
	mockRepo.On("ModificarEvento", req).Return(nil)

	err := mockRepo.ModificarEvento(req)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestModificarEvento_Error(t *testing.T) {
	mockRepo := new(mocks.EventoRepository)
	req := evento.EventoRequest{IdEvento: -1}
	mockRepo.On("ModificarEvento", req).Return(errors.New("id inválido"))

	err := mockRepo.ModificarEvento(req)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCancelarEvento_OK(t *testing.T) {
	mockRepo := new(mocks.EventoRepository)
	mockRepo.On("CancelarEvento", 10).Return(nil)

	err := mockRepo.CancelarEvento(10)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCancelarEvento_Error(t *testing.T) {
	mockRepo := new(mocks.EventoRepository)
	mockRepo.On("CancelarEvento", 10).Return(errors.New("error al cancelar"))

	err := mockRepo.CancelarEvento(10)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestEliminarEvento_OK(t *testing.T) {
	mockRepo := new(mocks.EventoRepository)
	mockRepo.On("EliminarEvento", 7).Return(nil)

	err := mockRepo.EliminarEvento(7)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestEliminarEvento_Error(t *testing.T) {
	mockRepo := new(mocks.EventoRepository)
	mockRepo.On("EliminarEvento", 7).Return(errors.New("fallo al eliminar"))

	err := mockRepo.EliminarEvento(7)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
