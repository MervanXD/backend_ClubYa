package detalledisponibilidad

import (
	"context"
	"database/sql"
)

type DetalleDisponibilidadRepository interface {
	ActualizarEstadoDetalleDisponibilidad(etalle DetalleRequestActualizar) (err error)
	ObtenerDisponibilidadEspacioSocialPorId(ctx context.Context, idEspacio int, idHorarioDia int, idBloqueTiempo int) (*DisponibilidadEspacioResponse, error)
	ObtenerDetalleDisponibilidadEspacioFechaId(idEspacio int, fecha string) ([]DetalleDisponibilidadDto, error)
	ObtenerRangosInicioDisponibles(idEspacio int, fecha string) ([]string, error)
	ActualizarDisponibilidadSegunReservaTx(tx *sql.Tx, detalle DetalleRequestActualizar) (int, error)
	ActualizarDisponibilidadPorSesion(sesionID int64, dia string, horaInicio string, horaFin string, idGrupo int) error
}
