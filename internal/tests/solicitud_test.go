package tests

import (
	"testing"

	"github.com/MervanXD/backend_ClubYa/internal/models/persona"
	solicitud "github.com/MervanXD/backend_ClubYa/internal/models/solicitud_membresia"
	"github.com/MervanXD/backend_ClubYa/internal/models/solicitud_membresia/mocks"
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/stretchr/testify/assert"
)

func TestObtenerSolicitudesMembresia_OK(t *testing.T) {
	mockRepo := new(mocks.SolicitudRepository)
	mockSolicitudes := []solicitud.SolicitudDTO{
		{Id: 1, Estado: tipos.EstadoSolicitud(0), Nombres: "Juan", Apellidos: "Perez"},
		{Id: 2, Estado: tipos.EstadoSolicitud(1), Nombres: "Maria", Apellidos: "Gomez"},
	}
	mockRepo.On("ObtenerSolicitudesMembresia").Return(mockSolicitudes, nil)

	solicitudes, err := mockRepo.ObtenerSolicitudesMembresia()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(solicitudes))
	assert.Equal(t, "Pendiente", solicitudes[0].Estado.String())
	mockRepo.AssertExpectations(t)
}

func TestObtenerSolicitudesMembresia_Error(t *testing.T) {
	mockRepo := new(mocks.SolicitudRepository)
	mockRepo.On("ObtenerSolicitudesMembresia").Return(nil, assert.AnError)

	solicitudes, err := mockRepo.ObtenerSolicitudesMembresia()

	assert.Error(t, err)
	assert.Nil(t, solicitudes)
	mockRepo.AssertExpectations(t)
}

func TestActualizarEstadoSolicitud_OK(t *testing.T) {
	mockRepo := new(mocks.SolicitudRepository)
	mockRepo.On("ActualizarEstadoSolicitud", 1, "Aprobada").Return(nil)

	err := mockRepo.ActualizarEstadoSolicitud(1, "Aprobada")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestActualizarEstadoSolicitud_Error(t *testing.T) {
	mockRepo := new(mocks.SolicitudRepository)
	mockRepo.On("ActualizarEstadoSolicitud", 1, "Aprobada").Return(assert.AnError)

	err := mockRepo.ActualizarEstadoSolicitud(1, "Aprobada")

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestObtenerDatosSolicitudPorId_OK(t *testing.T) {
	mockRepo := new(mocks.SolicitudRepository)
	expected := &solicitud.SolicitudMembresia{Id: 1, Estado: tipos.EstadoSolicitud(0)}
	mockRepo.On("ObtenerDatosSolicitudPorId", 1).Return(expected, nil)

	result, err := mockRepo.ObtenerDatosSolicitudPorId(1)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestObtenerDatosSolicitudPorId_Error(t *testing.T) {
	mockRepo := new(mocks.SolicitudRepository)
	mockRepo.On("ObtenerDatosSolicitudPorId", 1).Return(nil, assert.AnError)

	result, err := mockRepo.ObtenerDatosSolicitudPorId(1)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestObtenerFamiliaresPorIdSolicitud_OK(t *testing.T) {
	mockRepo := new(mocks.SolicitudRepository)
	expected := []persona.Familiar{
		{
			Persona: persona.Persona{
				Id:              1,
				Nombre:          "Pedro",
				Apellidos:       "Ramirez",
				Dni:             "87654321B",
				Sexo:            tipos.Sexo(1),
				FechaNacimiento: "1985-05-05",
			},
			TipoFamiliar:             tipos.TipoFamiliar(0),
			MismaDireccionPostulante: true,
		},
		{
			Persona: persona.Persona{
				Id:              2,
				Nombre:          "Lucia",
				Apellidos:       "Martinez",
				Dni:             "23456789C",
				Sexo:            tipos.Sexo(0),
				FechaNacimiento: "1992-02-02",
			},
			TipoFamiliar:             tipos.TipoFamiliar(1),
			MismaDireccionPostulante: false,
		},
	}
	mockRepo.On("ObtenerFamiliaresPorIdSolicitud", 1).Return(expected, nil)

	result, err := mockRepo.ObtenerFamiliaresPorIdSolicitud(1)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestObtenerFamiliaresPorIdSolicitud_Error(t *testing.T) {
	mockRepo := new(mocks.SolicitudRepository)
	mockRepo.On("ObtenerFamiliaresPorIdSolicitud", 1).Return(nil, assert.AnError)

	result, err := mockRepo.ObtenerFamiliaresPorIdSolicitud(1)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestObtenerDatosPersonaPorIdSolicitud_OK(t *testing.T) {
	mockRepo := new(mocks.SolicitudRepository)
	expected := &persona.Titular{
		Persona: persona.Persona{
			Nombre:          "Ana",
			Apellidos:       "Lopez",
			Dni:             "12345678A",
			Sexo:            tipos.Sexo(0),
			FechaNacimiento: "1990-01-01",
			Telefono:        "987654321",
			Pais:            "Peru",
		},
		Ocupacion:       utils.NullString{},
		NombreEmpresa:   utils.NullString{},
		IngresoPromedio: 3000.00,
	}
	mockRepo.On("ObtenerDatosPersonaPorIdSolicitud", 1).Return(expected, nil)

	result, err := mockRepo.ObtenerDatosPersonaPorIdSolicitud(1)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestObtenerDatosPersonaPorIdSolicitud_Error(t *testing.T) {
	mockRepo := new(mocks.SolicitudRepository)
	mockRepo.On("ObtenerDatosPersonaPorIdSolicitud", 1).Return(nil, assert.AnError)

	result, err := mockRepo.ObtenerDatosPersonaPorIdSolicitud(1)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestObtenerEstadoSolicitudPorID_OK(t *testing.T) {
	mockRepo := new(mocks.SolicitudRepository)
	mockRepo.On("ObtenerEstadoSolicitudPorID", 1).Return("Aprobada", nil)

	result, err := mockRepo.ObtenerEstadoSolicitudPorID(1)

	assert.NoError(t, err)
	assert.Equal(t, "Aprobada", result)
	mockRepo.AssertExpectations(t)
}

func TestObtenerEstadoSolicitudPorID_Error(t *testing.T) {
	mockRepo := new(mocks.SolicitudRepository)
	mockRepo.On("ObtenerEstadoSolicitudPorID", 1).Return("", assert.AnError)

	result, err := mockRepo.ObtenerEstadoSolicitudPorID(1)

	assert.Error(t, err)
	assert.Equal(t, "", result)
	mockRepo.AssertExpectations(t)
}
