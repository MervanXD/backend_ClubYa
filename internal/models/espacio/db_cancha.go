package espacio

import (
	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type canchaRespositoryDB struct{}

func NewCanchaRepositoryDB() CanchaRepository {
	return &canchaRespositoryDB{}
}

func (r *canchaRespositoryDB) ObtenerCanchasHorarios() ([]CanchaHorarioDTO, error) {
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

func (r *canchaRespositoryDB) ObtenerCanchasConfiguracion() ([]Cancha, error) {
	query := "call ingesoft.ListarCanchasConfiguracion()"
	rows, err := database.DB.Query(query)
	if err != nil {
		logs.Logger.Println("Error al obtener las canchas: ", err)
		return nil, err
	}
	defer rows.Close()
	var canchas []Cancha
	for rows.Next() {
		var cancha Cancha
		//var actividad string
		if err := rows.Scan(&cancha.Id, &cancha.Nombre, &cancha.Codigo, &cancha.Deporte); err != nil {
			logs.Logger.Println("Error al escanear la cancha: ", err)
			return nil, err
		}
		canchas = append(canchas, cancha)
	}
	return canchas, nil
}
