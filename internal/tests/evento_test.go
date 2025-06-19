package tests

import (
	"testing"
	"time"

	"github.com/MervanXD/backend_ClubYa/internal/models/evento"
	"github.com/MervanXD/backend_ClubYa/internal/models/evento/mocks"
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
	"github.com/stretchr/testify/assert"
)

func TestListarEventos_OK(t *testing.T){
	mockRepo := new(mocks.EventoRepository)
	expected := []evento.Evento{
		{
			IdEvento: 1,
			Nombre:      "Torneo"
			Descripcion: "Torneo anual"
			Fecha: 		  time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
			Aforo:        100,
			Invitados:    10,
			Precio:       50.0,
			Imagen:       []byte{},
			HoraInicio:   time.Date(0, 1, 1, 10, 0, 0, 0, time.UTC),
			HoraFin:      time.Date(0, 1, 1, 12, 0, 0, 0, time.UTC),
			NroInscritos: 20,
		}
		
	}
	mockRepo.On("ListarEventos").Return(expected, nil)

	result, err :=mockRepo.ListarEventos()

	assert.NoError(t, err)
	assert.Equal(t,expected,result)
	mockRepo.AssertExpectations(t)
}