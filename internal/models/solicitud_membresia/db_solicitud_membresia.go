package solicitud

import (
	"errors"
	"strings"

	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

func ActualizarEstadoSolicitud(id int, nuevoEstado string) error {
	estado := strings.Title(strings.ToLower(nuevoEstado))
	if estado != "Aceptada" && estado != "Rechazada" && estado != "Pendiente" {
		return errors.New("estado no válido")
	}

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
		return errors.New("no se actualizó ninguna fila")
	}

	return nil
}

