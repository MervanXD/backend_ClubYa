package reserva

import (
	"github.com/MervanXD/backend_ClubYa/database"
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

func ReservarEspacioSocial(reserva ReservaEspacio) error {
	query := "call ingesoft.ReservarEspacioSocial(?, ?, ?, ?, ?, ?,?)"
	_, err := database.DB.Exec(query, reserva.IdSocio, reserva.Espacio.Id, reserva.IdHorarioDia, reserva.IdBloqueTiempo, reserva.Fecha, reserva.HoraInicio, reserva.HoraFin)

	if err != nil {
		logs.Logger.Println("Error al reservar el espacio social: ", err)
		return err
	}

	err = detalledisponibilidad.ActualizarEstadoDetalleDisponibilidad(reserva.IdHorarioDia, reserva.IdBloqueTiempo, tipos.Reservado.String())
	if err != nil {
		logs.Logger.Println("Error al ReservarEspacio: ", err)
		return err
	}
	return nil

}
