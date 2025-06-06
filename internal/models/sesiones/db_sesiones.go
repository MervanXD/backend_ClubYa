package sesiones

import (
	"time"

	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

func ObtenerSesionesGrupo(idGrupo int) ([]Sesion, error) {
	sesionesQuery := "CALL ListarSesionPorGruposAcademia(?)"
	sesionesRows, err := database.DB.Query(sesionesQuery, idGrupo)
	if err != nil {
		logs.Logger.Println("Error al obtener las sesiones del grupo de la academia:", err)
		return nil, err
	}
	defer sesionesRows.Close()
	var sesiones []Sesion
	var inicio string
	var fin string
	for sesionesRows.Next() {
		var sesion Sesion
		if err := sesionesRows.Scan(&sesion.IDsesion, &sesion.Dia, &inicio, &fin); err != nil {
			sesionesRows.Close()
			logs.Logger.Println("Error al escanear la sesion del grupo:", err)
			return nil, err
		}

		horaInicio, err := time.Parse("15:04:05", inicio)
		if err != nil {
			logs.Logger.Println("Error al parsear horaInicio: ", err)
			return nil, err
		}
		horaFin, err := time.Parse("15:04:05", fin)
		if err != nil {
			logs.Logger.Println("Error al parsear horaFin: ", err)
			return nil, err
		}
		sesion.HoraInicio = horaInicio
		sesion.HoraFin = horaFin
		sesiones = append(sesiones, sesion)
	}
	return sesiones, nil
}
