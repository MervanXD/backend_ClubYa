package detalledisponibilidad

import "context"

type DetalleDisponibilidadRepository interface {
	ActualizarEstadoDetalleDisponibilidad(etalle DetalleRequestActualizar) (err error)
	ObtenerDisponibilidadEspacioSocialPorId(ctx context.Context, idEspacio int, idHorarioDia int, idBloqueTiempo int) (*DisponibilidadEspacioResponse, error)
	ObtenerDetalleDisponibilidadEspacioFechaId(idEspacio int, fecha string) ([]DetalleDisponibilidadDto, error)
	ObtenerRangosInicioDisponibles(idEspacio int, fecha string) ([]string, error)
}
