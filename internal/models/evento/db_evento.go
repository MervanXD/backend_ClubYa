package evento

import (
	"time"

	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

func ListarEventos() ([]Evento, error) {
	query := "call ingesoft.ListarEventos()"
	rows, err := database.DB.Query(query)
	if err != nil {
		logs.Logger.Println("Error al listar eventos: ", err)
		return nil, err
	}
	defer rows.Close()
	var eventos []Evento
	for rows.Next() {
		var e Evento
		var horaInicio string
		var horaFin string
		if err := rows.Scan(
			&e.IdEvento,
			&e.Nombre,
			&e.Descripcion,
			&e.Fecha,
			&e.Aforo,
			&e.Invitados,
			&e.Precio,
			&e.Imagen,
			&horaInicio,
			&horaFin,
			&e.NroInscritos); err != nil {
			logs.Logger.Println("Error al escanear evento: ", err)
			return nil, err
		}
		e.HoraInicio, err = time.Parse("15:04:05", horaInicio)
		if err != nil {
			logs.Logger.Println("Error al parsear hora de inicio: ", err)
			return nil, err
		}
		e.HoraFin, err = time.Parse("15:04:05", horaFin)
		if err != nil {
			logs.Logger.Println("Error al parsear hora de fin: ", err)
			return nil, err
		}

		eventos = append(eventos, e)
	}
	return eventos, nil
}

func BuscarEventoPorID(id int) (*Evento, error) {
	query := "CALL ObtenerEventoPorId(?)"
	row := database.DB.QueryRow(query, id)

	var e Evento
	var horaInicioStr, horaFinStr string

	err := row.Scan(
		&e.IdEvento,
		&e.Nombre,
		&e.Descripcion,
		&e.Fecha,
		&e.Aforo,
		&e.Invitados,
		&e.Precio,
		&e.Imagen,
		&e.Reglamento,
		&horaInicioStr,
		&horaFinStr,
		&e.NroInscritos,
	)
	if err != nil {
		logs.Logger.Println("Error al buscar evento por ID: ", err)
		return nil, err
	}

	// Parsear horas
	e.HoraInicio, err = time.Parse("15:04:05", horaInicioStr)
	if err != nil {
		logs.Logger.Println("Error al parsear horaInicio: ", err)
		return nil, err
	}
	e.HoraFin, err = time.Parse("15:04:05", horaFinStr)
	if err != nil {
		logs.Logger.Println("Error al parsear horaFin: ", err)
		return nil, err
	}

	return &e, nil
}
