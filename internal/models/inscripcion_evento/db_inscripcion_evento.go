package inscripcion_evento

import (
	"time"

	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

// RegistrarInscripcion permite registrar a un socio en un evento.
// Ejecuta el procedimiento almacenado 'RegistrarInscripcionEvento'.
func RegistrarInscripcion(fidPersona, idEvento, cantidadInvitados int) error {
	query := "CALL RegistrarInscripcionEvento(?, ?, ?)"
	_, err := database.DB.Exec(query, fidPersona, idEvento, cantidadInvitados)
	if err != nil {
		logs.Logger.Println("Error al registrar la inscripción al evento: ", err)
		return err
	}
	return nil
}

func ObtenerEventosSocio(idSocio int) ([]InscripcionSocioDTO, error) {
	query := "call ingesoft.ListarEventosSocio(?)"
	rows, err := database.DB.Query(query, idSocio)
	if err != nil {
		logs.Logger.Println("Error al obtener la informacion de los eventos del socio: ", err)
		return nil, err
	}
	defer rows.Close()
	var eventos []InscripcionSocioDTO
	var inicio string
	var fin string
	for rows.Next() {
		var evento InscripcionSocioDTO
		if err := rows.Scan(&evento.IdInscripcionEvento, &evento.FechaInscripcion, &evento.Estado,
			&evento.Evento.IdEvento, &evento.Evento.Nombre, &evento.Evento.Descripcion,
			&evento.Evento.Fecha, &evento.Evento.Precio, &inicio, &fin); err != nil {
			logs.Logger.Println("Error al escanear la informacion del evento: ", err)
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
		evento.Evento.HoraInicio = horaInicio
		evento.Evento.HoraFin = horaFin
		eventos = append(eventos, evento)
	}
	return eventos, nil
}
