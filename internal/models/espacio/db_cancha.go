package espacio

import (
	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type canchaRespositoryDB struct{}

func NewCanchaRepositoryDB() CanchaRepository {
	return &canchaRespositoryDB{}
}

func(r *canchaRespositoryDB) ObtenerCanchasHorarios() ([]CanchaHorarioDTO, error) {
	query := "call ingesoft.listarCanchasHorarios()"
	rows, err := database.DB.Query(query)
	if err != nil {
		logs.Logger.Println("Error al obtener los horarios de las lozas deportivas: ", err)
		return nil, err
	}
	defer rows.Close()
	var canchasHorarios []CanchaHorarioDTO
	for rows.Next() {
		var es CanchaHorarioDTO
		if err := rows.Scan(&es.Espacio.Id, &es.Espacio.Codigo, &es.Espacio.Nombre, &es.Espacio.Imagen, &es.Espacio.Deporte,
			&es.Espacio.Ubicacion, &es.Espacio.Capacidad, &es.Espacio.Costo, &es.Fecha, &es.HoraInicio, &es.HoraFinal,
			&es.Estado, &es.IdHorario, &es.IdBloque); err != nil {
			logs.Logger.Println("Error al escanear el horario de la loza deportiva: ", err)
			return nil, err
		}
		canchasHorarios = append(canchasHorarios, es)
	}
	return canchasHorarios, nil
}
