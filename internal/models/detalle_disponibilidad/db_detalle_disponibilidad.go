package detalledisponibilidad

import (
	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

func ActualizarEstadoDetalleDisponibilidad(idHorarioDia int, idBloqueTiempo int, estado string) (err error) {
	stmt := "call ActualizarEstadoDetalleDisponibilidad(?,?,?)"
	result, err := database.DB.Exec(stmt, idHorarioDia, idBloqueTiempo, estado)
	if err != nil {
		logs.Logger.Println("Error al ActualizarEstadoDetalleDisponibilidad: ", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logs.Logger.Println("Error al obtener rowsAffected:", err)
		return err
	}

	if rowsAffected == 0 {
		return err
	}

	return nil
}
