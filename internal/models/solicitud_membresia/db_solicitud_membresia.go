package solicitud

import (
	"strings"

	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)


func ObtenerSolicitudesMembresia() ([]SolicitudDTO, error) {
	query := "call ingesoft.obtenerSolicitudesConTitulares()"
	rows, err := database.DB.Query(query)
	if err != nil {
		logs.Logger.Fatal("Error al obtener las solicitudes de membresia: ", err)
		return nil, err
	}
	defer rows.Close()
	var solicitudes []SolicitudDTO
	for rows.Next() {
		var solicitud SolicitudDTO
		if err := rows.Scan(&solicitud.Id, &solicitud.Fecha, &solicitud.Estado, &solicitud.IdTitular, &solicitud.Nombres, &solicitud.Apellidos); err != nil {
			logs.Logger.Fatal("Error al escanear solicitud de membresia: ", err)
			return nil, err
		}
		solicitudes = append(solicitudes, solicitud)
	}
	return solicitudes, nil
}

func ActualizarEstadoSolicitud(id int, nuevoEstado string) error {
	estado := strings.Title(strings.ToLower(nuevoEstado))

	query := `CALL ActualizarEstadoSolicitud(?, ?)`
	result, err := database.DB.Exec(query, id, estado)
	if err != nil {
		logs.Logger.Println("Error al ejecutar el procedimiento:", err)
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