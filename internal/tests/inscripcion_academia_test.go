package tests

import (
	"testing"

	inscripcionacademia "github.com/MervanXD/backend_ClubYa/internal/models/inscripcion_academia"
	"github.com/MervanXD/backend_ClubYa/internal/models/inscripcion_academia/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRegistrarInscripcionAcademia_OK(t *testing.T) {
	mockRepo := new(mocks.InscripcionAcademiaRepository)
	mockRepo.On("RegistrarInscripcionAcademia", 1, 2, 3, 1, 100.0).Return(nil)

	err := mockRepo.RegistrarInscripcionAcademia(1, 2, 3, 1, 100.0)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRegistrarInscripcionAcademia_Error(t *testing.T) {
	mockRepo := new(mocks.InscripcionAcademiaRepository)
	mockRepo.On("RegistrarInscripcionAcademia", 1, 2, 3, 1, 100.0).Return(assert.AnError)

	err := mockRepo.RegistrarInscripcionAcademia(1, 2, 3, 1, 100.0)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestObtenerFamiliaresInscritosAcademia_OK(t *testing.T) {
	mockRepo := new(mocks.InscripcionAcademiaRepository)
	expected := []inscripcionacademia.InscritoAcademiaDTO{
		{IdPersona: 1, NombrePersona: "Juan", ApellidoPersona: "Perez", NombreAcademia: "Fútbol", IdGrupo: 10},
	}
	mockRepo.On("ObtenerFamiliaresInscritosAcademia", 1).Return(expected, nil)

	result, err := mockRepo.ObtenerFamiliaresInscritosAcademia(1)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestObtenerFamiliaresInscritosAcademia_Error(t *testing.T) {
	mockRepo := new(mocks.InscripcionAcademiaRepository)
	mockRepo.On("ObtenerFamiliaresInscritosAcademia", 1).Return(nil, assert.AnError)

	result, err := mockRepo.ObtenerFamiliaresInscritosAcademia(1)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
