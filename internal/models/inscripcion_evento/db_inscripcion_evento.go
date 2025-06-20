package inscripcion_evento

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type inscripcionEventoRepositoryDB struct{}

func NewInscripcionEventoRepositoryDB() InscripcionEventoRepository {
	return &inscripcionEventoRepositoryDB{}
}

// RegistrarInscripcion permite registrar a un socio en un evento.
// Ejecuta el procedimiento almacenado 'RegistrarInscripcionEvento'.
func (r *inscripcionEventoRepositoryDB) RegistrarInscripcion(fidPersona, idEvento, cantidadInvitados int) error {
	query := "CALL RegistrarInscripcionEvento(?, ?, ?)"
	_, err := database.DB.Exec(query, fidPersona, idEvento, cantidadInvitados)
	if err != nil {
		logs.Logger.Println("Error al registrar la inscripción al evento: ", err)
		return err
	}
	return nil
}

func (r *inscripcionEventoRepositoryDB) ObtenerEventosSocio(idSocio int) ([]InscripcionSocioDTO, error) {
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

func (r *inscripcionEventoRepositoryDB) AnularInscripcion(idInscripcionEvento int, motivo string) error {
	query := "CALL AnularInscripcionEvento(?, ?, @errorMsg)"
	_, err := database.DB.Exec(query, idInscripcionEvento, motivo)
	if err != nil {
		logs.Logger.Println("Error al ejecutar el procedimiento almacenado: ", err)
		return fmt.Errorf("error al ejecutar el procedimiento: %w", err)
	}

	var errorMsg sql.NullString
	err = database.DB.QueryRow("SELECT @errorMsg").Scan(&errorMsg)
	if err != nil {
		logs.Logger.Println("Error al leer la variable @errorMsg: ", err)
		return fmt.Errorf("error al leer variable de salida: %w", err)
	}

	if errorMsg.Valid {
		logs.Logger.Println("Error de negocio desde el procedimiento: ", errorMsg.String)
		return errors.New(errorMsg.String)
	}

	return nil
}

func (r *inscripcionEventoRepositoryDB) PagarInscripcion(fidPersona int, idEvento int, concepto string, metodoPago string, monto float64) (int, error) {
	query := "CALL CrearPagoEvento(?, ?, ?, ?, ?, @idPago)"
	_, err := database.DB.Exec(query, fidPersona, idEvento, concepto, metodoPago, monto)
	if err != nil {
		logs.Logger.Println("Error al registrar el pago de la inscripción: ", err)
		return -1, err
	}

	var idPago int
	err = database.DB.QueryRow("SELECT @idPago").Scan(&idPago)
	if err != nil {
		logs.Logger.Println("Error al leer la variable @idPago: ", err)
		return -1, fmt.Errorf("error al leer variable de salida: %w", err)
	}

	if idPago <= 0 {
		logs.Logger.Println("idPago inválido recibido:", idPago)
		return -1, err
	}

	return idPago, nil
}
