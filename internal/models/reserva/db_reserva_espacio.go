package reserva

import (
	"context"
	"fmt"

	"github.com/MervanXD/backend_ClubYa/database"
	detalledisponibilidad "github.com/MervanXD/backend_ClubYa/internal/models/detalle_disponibilidad"
	"github.com/MervanXD/backend_ClubYa/internal/models/pago"
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	servicios "github.com/MervanXD/backend_ClubYa/internal/services"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type reservaRepositoryDB struct{}

func NewReservaRepositoryDB() ReservaRepository {
	return &reservaRepositoryDB{}
}

func (r *reservaRepositoryDB) ReservarEspacio(ctx context.Context, re ReservaEspacio) error {
	// 1) Pre‐validaciones
	bloques := re.HorarioDia.BloquesTiempo
	if len(bloques) == 0 {
		return fmt.Errorf("no time blocks")
	}
	cantidad := len(bloques)
	montoTotal := re.Espacio.Costo * float64(cantidad)
	if (re.Pago != pago.Pago{}) && re.Pago.Monto != montoTotal {
		return fmt.Errorf("expected payment %v, got %v", montoTotal, re.Pago.Monto)
	}
	intervalos := utils.ConvertBloquesToIntervals(bloques)
	horaIni, horaFin, err := utils.ValidateAndSortIntervals(intervalos)
	if err != nil {
		return fmt.Errorf("invalid time blocks: %w", err)
	}

	if re.HoraInicio != horaIni || re.HoraFin != horaFin {
		return fmt.Errorf("expected time range %s-%s, got %s-%s", horaIni, horaFin, re.HoraInicio, re.HoraFin)
	}

	// 2) Start tx + defer rollback
	tx, err := database.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer func() {
		_ = tx.Rollback()
	}()

	// 3) Actualiza disponibilidad
	repo2 := detalledisponibilidad.NewDetalleDisponibilidadRepositoryDB()
	idHorario := re.HorarioDia.IdHorarioDia
	for _, b := range bloques {
		detalle := detalledisponibilidad.DetalleRequestActualizar{
			IdHorarioDia:         idHorario,
			IdBloqueTiempo:       b.IdBloqueTiempo,
			EstadoDisponibilidad: tipos.Reservado,
			Fecha:                re.Fecha,
			Dia:                  re.HorarioDia.Dia,
			Id_Espacio:           re.Espacio.Id,
		}
		nuevoID, err := repo2.ActualizarDisponibilidadSegunReservaTx(tx, detalle)
		if err != nil {
			return fmt.Errorf("update dispo: %w", err)
		}
		if idHorario < 0 {
			idHorario = nuevoID
		} else if nuevoID != idHorario {
			return fmt.Errorf("horario mismatch %d != %d", idHorario, nuevoID)
		}
	}

	// 4) Reserva espacio
	if err := utils.ExecSPWithOut(ctx, tx, "ReservarEspacio", "id_Reserva", &re.Id,
		re.IdSocio, re.Espacio.Id, re.Fecha, re.HoraInicio, re.HoraFin, idHorario,
	); err != nil {
		return err
	}
	logs.Logger.Println("Datos pago:", re.Pago.Metodo, re.Pago.Monto)

	// 5) Enlaza bloques con reserva
	for _, b := range bloques {
		if err := utils.ExecSP(ctx, tx, "EnlazarBloqueReservaEspacio",
			idHorario, b.IdBloqueTiempo, re.Id,
		); err != nil {
			return fmt.Errorf("link block to reservation: %w", err)
		}
	}
	// 6) Crea pago si corresponde
	if (re.Pago != pago.Pago{}) {
		if err := utils.ExecSPWithOut(ctx, tx, "CrearPagoReservaEspacio", "p_idPago", &re.Pago.IdPago,
			re.IdSocio, re.Id, "Pago de reserva de espacio",
			re.Pago.Metodo, re.Pago.Monto,
		); err != nil {
			return err
		}
	}

	// 7) Commit
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}
	return nil
}

func (r *reservaRepositoryDB) AnulacionReservaEspacio(idReserva int, idEspacio int, idHorarioDia int, motivo string) error {
	query := "call ingesoft.AnularReservaEspacio(?, ?, ?, ?)"
	_, err := database.DB.Exec(query, idReserva, idEspacio, idHorarioDia, motivo)
	if err != nil {
		logs.Logger.Println("Error al cancelar la reserva del espacio", err)
		return err
	}
	return nil
}

func (r *reservaRepositoryDB) ObtenerReservasEspaciosSocialesSocio(idSocio int) ([]ReservaEspacioSocialRequest, error) {
	query := "call ingesoft.ListarReservasEspaciosSocialesSocio(?)"
	rows, err := database.DB.Query(query, idSocio)
	if err != nil {
		logs.Logger.Println("Error al obtener los horarios de los espacios sociales del socio: ", err)
		return nil, err
	}
	defer rows.Close()
	var reservasSocio []ReservaEspacioSocialRequest
	for rows.Next() {
		var reserva ReservaEspacioSocialRequest
		if err := rows.Scan(
			&reserva.Id, &reserva.FechaReserva, &reserva.Fecha, &reserva.HoraInicio, &reserva.HoraFin,
			&reserva.Estado, &reserva.IdHorarioDia, &reserva.Espacio.Id,
			&reserva.Espacio.Codigo, &reserva.Espacio.Nombre, &reserva.Espacio.Ubicacion,
			&reserva.Espacio.Capacidad, &reserva.Espacio.Costo, &reserva.Espacio.Actividad); err != nil {
			logs.Logger.Println("Error al escanear el horario de la loza deportiva del socio: ", err)
			return nil, err
		}
		reservasSocio = append(reservasSocio, reserva)
	}
	return reservasSocio, nil
}

func (r *reservaRepositoryDB) ObtenerReservasCanchasSocio(idSocio int) ([]ReservaCanchaRequest, error) {
	query := "call ingesoft.ListarReservasCanchasSocio(?)"
	rows, err := database.DB.Query(query, idSocio)
	if err != nil {
		logs.Logger.Println("Error al obtener los horarios de las lozas deportivas del socio: ", err)
		return nil, err
	}
	defer rows.Close()
	var reservasSocio []ReservaCanchaRequest
	for rows.Next() {
		var reserva ReservaCanchaRequest
		if err := rows.Scan(
			&reserva.Id, &reserva.FechaReserva, &reserva.Fecha, &reserva.HoraInicio, &reserva.HoraFin,
			&reserva.Estado, &reserva.IdHorarioDia, &reserva.Espacio.Id,
			&reserva.Espacio.Codigo, &reserva.Espacio.Nombre, &reserva.Espacio.Ubicacion,
			&reserva.Espacio.Capacidad, &reserva.Espacio.Costo, &reserva.Espacio.Deporte); err != nil {
			logs.Logger.Println("Error al escanear el horario de la loza deportiva del socio: ", err)
			return nil, err
		}

		reservasSocio = append(reservasSocio, reserva)
	}
	return reservasSocio, nil
}

func (r *reservaRepositoryDB) ObtenerReservasPorEspacio(idEspacio int) ([]ReservaRequest, error) {
	query := "call ingesoft.ListarReservasPorEspacio(?)"
	rows, err := database.DB.Query(query, idEspacio)
	if err != nil {
		logs.Logger.Println("Error al obtener las reservas del espacio: ", err)
		return nil, err
	}
	defer rows.Close()

	var reservas []ReservaRequest
	for rows.Next() {
		var reserva ReservaRequest
		var anulacion AnulacionReserva
		err := rows.Scan(
			&reserva.IdReserva,
			&reserva.NombreSocio,
			&reserva.FechaReserva,
			&reserva.HoraInicio,
			&reserva.HoraFin,
			&reserva.Estado,
			&anulacion.Id, &anulacion.Fecha, &anulacion.Motivo, &anulacion.Devolucion,
		)
		if err != nil {
			logs.Logger.Println("Error al escanear reserva: ", err)
			return nil, err
		}
		// Si hay anulación asociada, la agregas
		if anulacion.Id != 0 {
			reserva.AnulacionReserva = &anulacion
		}
		reservas = append(reservas, reserva)
	}
	return reservas, nil
}

func (r *reservaRepositoryDB) AceptarDevolucionAnulacionReserva(anulacion AnulacionReservaRequest) error {
	query := "call ingesoft.AceptarDevolucionAnulacionReserva(?, ?, ?, @p_correo)"
	_, err := database.DB.Exec(query, anulacion.IdReserva, anulacion.IdAnulacion, anulacion.PorcentajeDevolucion)
	if err != nil {
		logs.Logger.Println("Error al aceptar la devolución de la anulación de reserva: ", err)
		return err
	}
	// Recupera el valor del parámetro de salida
	var correo string
	row := database.DB.QueryRow("SELECT @p_correo")
	if err := row.Scan(&correo); err != nil {
		logs.Logger.Println("Error al obtener el correo de salida: ", err)
		return err
	}
	//Mandamos un correo de confirmacion al socio
	if correo != "" {
		err = servicios.EnviarCorreo([]string{correo}, // destinatario
			"Confirmación de devolución de anulación de reserva", // asunto
			// cuerpo del mensaje
			"Su solicitud de devolución por anulación de reserva ha sido aceptada. El porcentaje de devolución es: "+fmt.Sprintf("%.2f", anulacion.PorcentajeDevolucion)+"%.")
		if err != nil {
			logs.Logger.Println("Error al enviar correo de confirmación de devolución: ", err)
			return err
		}
	}

	return nil
}
