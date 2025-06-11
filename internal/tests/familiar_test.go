package tests

import (
	"testing"

	"github.com/MervanXD/backend_ClubYa/internal/models/persona"
	"github.com/MervanXD/backend_ClubYa/internal/models/persona/mocks"
	"github.com/stretchr/testify/assert"
)

func TestInsertarFamiliar_OK(t *testing.T) {
	mockRepo := new(mocks.FamiliarRepository)
	fam := persona.Familiar{Persona: persona.Persona{Nombre: "Juan"}}
	mockRepo.On("InsertarFamiliar", fam, 1).Return(nil)

	err := mockRepo.InsertarFamiliar(fam, 1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestInsertarFamiliar_Error(t *testing.T) {
	mockRepo := new(mocks.FamiliarRepository)
	fam := persona.Familiar{Persona: persona.Persona{Nombre: "Juan"}}
	mockRepo.On("InsertarFamiliar", fam, 1).Return(assert.AnError)

	err := mockRepo.InsertarFamiliar(fam, 1)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRegistrarFamiliares_OK(t *testing.T) {
	mockRepo := new(mocks.FamiliarRepository)
	req := persona.FamiliarResquest{IdTitular: 1}
	mockRepo.On("RegistrarFamiliares", req).Return(10, nil)

	id, err := mockRepo.RegistrarFamiliares(req)

	assert.NoError(t, err)
	assert.Equal(t, 10, id)
	mockRepo.AssertExpectations(t)
}

func TestRegistrarFamiliares_Error(t *testing.T) {
	mockRepo := new(mocks.FamiliarRepository)
	req := persona.FamiliarResquest{IdTitular: 1}
	mockRepo.On("RegistrarFamiliares", req).Return(-1, assert.AnError)

	id, err := mockRepo.RegistrarFamiliares(req)

	assert.Error(t, err)
	assert.Equal(t, -1, id)
	mockRepo.AssertExpectations(t)
}

func TestObtenerIdsFamiliaresPorTitular_OK(t *testing.T) {
	mockRepo := new(mocks.FamiliarRepository)
	expected := []persona.Familiar{
		{Persona: persona.Persona{Nombre: "Pedro"}},
	}
	mockRepo.On("ObtenerIdsFamiliaresPorTitular", 1).Return(expected, nil)

	result, err := mockRepo.ObtenerIdsFamiliaresPorTitular(1)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestObtenerIdsFamiliaresPorTitular_Error(t *testing.T) {
	mockRepo := new(mocks.FamiliarRepository)
	mockRepo.On("ObtenerIdsFamiliaresPorTitular", 1).Return(nil, assert.AnError)

	result, err := mockRepo.ObtenerIdsFamiliaresPorTitular(1)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestObtenerFamiliarPorIDPersona_OK(t *testing.T) {
	mockRepo := new(mocks.FamiliarRepository)
	expected := &persona.Familiar{Persona: persona.Persona{Nombre: "Lucia"}}
	mockRepo.On("ObtenerFamiliarPorIDPersona", 1).Return(expected, nil)

	result, err := mockRepo.ObtenerFamiliarPorIDPersona(1)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestObtenerFamiliarPorIDPersona_Error(t *testing.T) {
	mockRepo := new(mocks.FamiliarRepository)
	mockRepo.On("ObtenerFamiliarPorIDPersona", 1).Return(nil, assert.AnError)

	result, err := mockRepo.ObtenerFamiliarPorIDPersona(1)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
