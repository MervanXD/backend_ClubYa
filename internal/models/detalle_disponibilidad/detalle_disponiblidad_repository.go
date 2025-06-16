package detalledisponibilidad

import "context"

type DetalleDisponibilidadRepository interface {
	ActualizarEstadoDetalleDisponibilidad(idHorarioDia int, idBloqueTiempo int, estado string) (err error)
	ObtenerDisponibilidadEspacioSocialPorId(ctx context.Context, idEspacio int, idHorarioDia int, idBloqueTiempo int) (*DisponibilidadEspacioResponse, error)
	ObtenerDetalleDisponibilidadEspacioSocialFechaId(idEspacio int, fecha string) ([]DetalleDisponibilidad, error)
}
