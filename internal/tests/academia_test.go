package tests

import (
	"testing"

	"github.com/MervanXD/backend_ClubYa/internal/models/academia"
	"github.com/MervanXD/backend_ClubYa/internal/models/academia/mocks"
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestObtenerAcademias_OK(t *testing.T) {
	mockRepo := new(mocks.AcademiaRepository)
	mockAcademias := []academia.AcademiaDTO{
		{ID: 1, Nombre: "Fútbol", Descripcion: "Academia de fútbol", Deporte: tipos.Deporte(0)},
		{ID: 2, Nombre: "Natación", Descripcion: "Academia de natación", Deporte: tipos.Deporte(3)},
	}
	mockRepo.On("ObtenerAcademias", mock.Anything).Return(mockAcademias, nil)

	result, err := mockRepo.ObtenerAcademias()

	assert.NoError(t, err)
	assert.Equal(t, mockAcademias, result)
	mockRepo.AssertExpectations(t)
}

func TestObtenerAcademias_Error(t *testing.T) {
	mockRepo := new(mocks.AcademiaRepository)
	mockRepo.On("ObtenerAcademias").Return(nil, assert.AnError)

	result, err := mockRepo.ObtenerAcademias()

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestObtenerAcademiaPorId_OK(t *testing.T) {
	mockRepo := new(mocks.AcademiaRepository)
	expected := &academia.Academia{ID: 1, Nombre: "Fútbol"}
	mockRepo.On("ObtenerAcademiaPorId", 1).Return(expected, nil)

	result, err := mockRepo.ObtenerAcademiaPorId(1)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestObtenerAcademiaPorId_Error(t *testing.T) {
	mockRepo := new(mocks.AcademiaRepository)
	mockRepo.On("ObtenerAcademiaPorId", 1).Return(nil, assert.AnError)

	result, err := mockRepo.ObtenerAcademiaPorId(1)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
