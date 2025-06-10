package solicitud

import "github.com/MervanXD/backend_ClubYa/internal/models/persona"

type SolicitudRepository interface {
	ObtenerSolicitudesMembresia() ([]SolicitudDTO, error)
	ActualizarEstadoSolicitud(id int, nuevoEstado string) error 
	ObtenerDatosSolicitudPorId(idSolicitud int) (*SolicitudMembresia, error)
	ObtenerFamiliaresPorIdSolicitud(idSolicitud int) ([]persona.Familiar, error) 
	ObtenerDatosPersonaPorIdSolicitud(idSolicitud int) (*persona.Titular, error)
	ObtenerEstadoSolicitudPorID(idSolicitud int) (string, error) 
}
