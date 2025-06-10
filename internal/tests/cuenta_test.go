package tests

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/cuenta"
	"github.com/MervanXD/backend_ClubYa/internal/models/cuenta/mocks"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

func TestCrearCuentaSinUsername(t *testing.T) {
	logs.InitLogger()
	defer logs.CloseLogger()

	mockRepo := new(mocks.CuentaRepository)
	mockRepo.On("CrearCuenta", mock.Anything).Return(int64(123), nil)

	id, err := mockRepo.CrearCuenta(cuenta.Cuenta{
		Username:   "testuser",
		Contrasena: "1234",
		Email:      "test@correo.com",
	})

	assert.NoError(t, err)
	assert.Equal(t, int64(123), id)
	mockRepo.AssertExpectations(t)
}

func TestCrearCuentaNormal(t *testing.T) {
	mockRepo := new(mocks.CuentaRepository)
	mockRepo.On("CrearCuenta", mock.Anything).Return(int64(0), assert.AnError)

	id, err := mockRepo.CrearCuenta(cuenta.Cuenta{
		Username:   "",
		Contrasena: "1234",
		Email:      "test@correo.com",
	})

	assert.Error(t, err)
	assert.Equal(t, int64(0), id)
	mockRepo.AssertExpectations(t)
}
