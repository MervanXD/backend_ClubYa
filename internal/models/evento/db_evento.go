package evento

import (
	"errors"
	"time"

	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type EventoRequest struct {
	Nombre      string  `json:"nombre"`
	Descripcion string  `json:"descripcion"`
	Fecha       string  `json:"fecha"`
	Aforo       int     `json:"aforo"`
	Invitados   int     `json:"invitados"`
	Precio      float64 `json:"precio"`
	HoraInicio  string  `json:"hora_inicio"`
	HoraFin     string  `json:"hora_fin"`
	NroInscritos string `json:"nro_inscritos"`
	IdEspacio   int     `json:"id_espacio"`
}

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
		if err := rows.Scan(&e.IdEvento, &e.Nombre, &e.Descripcion, &e.Precio, &e.Fecha, &horaInicio, &horaFin, &e.Imagen, &e.NroInscritos); err != nil {
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

func InsertarEvento(req EventoRequest) (int, error) {
	fecha, err := time.Parse("2006-01-02", req.Fecha)
	if err != nil {
		return -1, errors.New("formato de fecha inválido, se esperaba YYYY-MM-DD")
	}

	horaInicio, err := time.Parse("15:04:05", req.HoraInicio)
	if err != nil {
		return -1, errors.New("formato de hora de inicio inválido, se esperaba HH:MM:SS")
	}

	horaFin, err := time.Parse("15:04:05", req.HoraFin)
	if err != nil {
		return -1, errors.New("formato de hora de fin inválido, se esperaba HH:MM:SS")
	}

	query := "CALL InsertarEvento(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,@p_idEvento)"
	_, err = database.DB.Exec(query,
		req.Nombre,
		req.Descripcion,
		fecha.Format("2006-01-02"),
		req.Aforo,
		req.Invitados,
		req.Precio,
		horaInicio.Format("15:04:05"),
		horaFin.Format("15:04:05"),
		req.NroInscritos,
		req.IdEspacio,
	)
	if err != nil {
		logs.Logger.Println("Error al ejecutar SP InsertarEvento:", err)
		return -1, err
	}

	var idEvento int
	err = database.DB.QueryRow("SELECT @p_idEvento").Scan(&idEvento)
	if err != nil {
		logs.Logger.Println("Error al obtener idEvento:", err)
		return -1, err
	}

	return idEvento, nil
}
