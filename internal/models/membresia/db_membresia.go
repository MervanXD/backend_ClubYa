package membresia

import (
	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

func BuscarMembresiaPorSolicitud(idSolicitud int) (int, error) {
	query := "CALL BuscarMembresiaPorSolicitud(?, @p_idMembresia)"
	_, err := database.DB.Exec(query, idSolicitud)
	if err != nil {
		logs.Logger.Println("Error al ejecutar el procedure BuscarMembresiaPorSolicitud: ", err)
		return 0, err
	}

	var idMembresia int
	err = database.DB.QueryRow("SELECT @p_idMembresia").Scan(&idMembresia)
	if err != nil {
		logs.Logger.Println("Error al obtener idMembresia desde variable OUT: ", err)
		return 0, err
	}

	return idMembresia, nil
}
