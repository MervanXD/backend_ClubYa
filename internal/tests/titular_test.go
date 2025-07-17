package tests

import (
    "context"
    "testing"

    "github.com/MervanXD/backend_ClubYa/internal/models/persona"
    "github.com/MervanXD/backend_ClubYa/internal/models/persona/mocks"
    "github.com/stretchr/testify/assert"
)

func TestInsertarTitular_OK(t *testing.T) {
    mockRepo := new(mocks.TitularRepository)
    titular := persona.Titular{Persona: persona.Persona{Nombre: "Carlos"}}
    mockRepo.On("InsertarTitular", titular, 1).Return(42, nil)

    id, err := mockRepo.InsertarTitular(titular, 1)

    assert.NoError(t, err)
    assert.Equal(t, 42, id)
    mockRepo.AssertExpectations(t)
}

func TestInsertarTitular_Error(t *testing.T) {
    mockRepo := new(mocks.TitularRepository)
    titular := persona.Titular{Persona: persona.Persona{Nombre: "Carlos"}}
    mockRepo.On("InsertarTitular", titular, 1).Return(-1, assert.AnError)

    id, err := mockRepo.InsertarTitular(titular, 1)

    assert.Error(t, err)
    assert.Equal(t, -1, id)
    mockRepo.AssertExpectations(t)
}

func TestObtenerTitularPorID_OK(t *testing.T) {
    mockRepo := new(mocks.TitularRepository)
    expected := &persona.Titular{Persona: persona.Persona{Nombre: "Lucia"}}
    ctx := context.Background()
    mockRepo.On("ObtenerTitularPorID", ctx, 1).Return(expected, nil)

    result, err := mockRepo.ObtenerTitularPorID(ctx, 1)

    assert.NoError(t, err)
    assert.Equal(t, expected, result)
    mockRepo.AssertExpectations(t)
}

func TestObtenerTitularPorID_Error(t *testing.T) {
    mockRepo := new(mocks.TitularRepository)
    ctx := context.Background()
    mockRepo.On("ObtenerTitularPorID", ctx, 1).Return(nil, assert.AnError)

    result, err := mockRepo.ObtenerTitularPorID(ctx, 1)

    assert.Error(t, err)
    assert.Nil(t, result)
    mockRepo.AssertExpectations(t)
}