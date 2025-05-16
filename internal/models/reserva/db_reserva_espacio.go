package reserva

import (
	detalledisponibilidad "github.com/MervanXD/backend_ClubYa/internal/models/detalle_disponibilidad"
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
	"github.com/MervanXD/backend_ClubYa/logs"
)

func ReservarEspacio(idEspacio int, idHorarioDia int, idBloqueTiempo int) (err error) {
	err = detalledisponibilidad.ActualizarEstadoDetalleDisponibilidad(idHorarioDia, idBloqueTiempo, tipos.Reservado.String())
	if err != nil {
		logs.Logger.Println("Error al ReservarEspacio: ", err)
		return err
	}
	return nil
}
