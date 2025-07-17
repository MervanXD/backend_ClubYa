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

func ListarMembresias() ([]MembresiaDTO, error) {
	query := "CALL ListarMembresias()"
	rows, err := database.DB.Query(query)
	if err != nil {
		logs.Logger.Println("Error al ejecutar el procedure ListarMembresias: ", err)
		return nil, err
	}
	defer rows.Close()

	var membresias []MembresiaDTO
	for rows.Next() {
		var m MembresiaDTO
		err = rows.Scan(&m.Membresia.Id, &m.Membresia.FechaInicio, &m.Membresia.FechaFin, &m.Membresia.Tipo, &m.Membresia.Estado, &m.NombreTitular, &m.CuotaBase)
		if err != nil {
			logs.Logger.Println("Error al escanear la membresia: ", err)
			return nil, err
		}
		membresias = append(membresias, m)
	}

	return membresias, nil
}
