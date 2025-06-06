package reserva

import (
	"github.com/MervanXD/backend_ClubYa/database"
	detalledisponibilidad "github.com/MervanXD/backend_ClubYa/internal/models/detalle_disponibilidad"
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
	"github.com/MervanXD/backend_ClubYa/logs"
)

func ReservarEspacio(idEspacio int, idHorarioDia int, idBloqueTiempo int) (err error) {
	err = detalledisponibilidad.ActualizarEstadoDetalleDisponibilidad(idHorarioDia, idBloqueTiempo, tipos.Reservado.String())
	if err != nil {
		logs.Logger.Println("Error al ReservarEspacio: ", err)
		return err
	}
	return nil
}

func ReservarEspacioSocial(reserva ReservaEspacio) error {
	query := "call ingesoft.ReservarEspacioSocial(?, ?, ?, ?, ?, ?,?)"
	_, err := database.DB.Exec(query, reserva.IdSocio, reserva.Espacio.Id, reserva.IdHorarioDia, reserva.IdBloqueTiempo, reserva.Fecha, reserva.HoraInicio, reserva.HoraFin)

	if err != nil {
		logs.Logger.Println("Error al reservar el espacio social: ", err)
		return err
	}

	err = detalledisponibilidad.ActualizarEstadoDetalleDisponibilidad(reserva.IdHorarioDia, reserva.IdBloqueTiempo, tipos.Reservado.String())
	if err != nil {
		logs.Logger.Println("Error al ReservarEspacio: ", err)
		return err
	}
	return nil
}

func AnulacionReservaEspacioSocial(idReserva int, idEspacio int, idHorarioDia int, idBloque int, motivo string) error {
	query := "call ingesoft.AnularReservaEspacioSocial(?, ?, ?, ?, ?)"
	_, err := database.DB.Exec(query, idReserva, idEspacio, idHorarioDia, idBloque, motivo)
	if err != nil {
		logs.Logger.Println("Error al cancelar la reserva del espacio ", err)
		return err
	}
	return nil
}

func ObtenerReservasEspaciosSocialesSocio(idSocio int) ([]ReservaEspacioSocialRequest, error) {
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
		//var actividad string
		if err := rows.Scan(
			&reserva.Id, &reserva.FechaReserva, &reserva.Fecha, &reserva.HoraInicio, &reserva.HoraFin,
			&reserva.Estado, &reserva.IdBloqueTiempo, &reserva.IdHorarioDia, &reserva.Espacio.Id,
			&reserva.Espacio.Codigo, &reserva.Espacio.Nombre, &reserva.Espacio.Ubicacion,
			&reserva.Espacio.Capacidad, &reserva.Espacio.Costo, &reserva.Espacio.Actividad); err != nil {
			logs.Logger.Println("Error al escanear el horario del espacio social del socio: ", err)
			return nil, err
		}

		reservasSocio = append(reservasSocio, reserva)
	}
	return reservasSocio, nil
}

func ObtenerReservasCanchasSocio(idSocio int) ([]ReservaCanchaRequest, error) {
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
			&reserva.Estado, &reserva.IdBloqueTiempo, &reserva.IdHorarioDia, &reserva.Espacio.Id,
			&reserva.Espacio.Codigo, &reserva.Espacio.Nombre, &reserva.Espacio.Ubicacion,
			&reserva.Espacio.Capacidad, &reserva.Espacio.Costo, &reserva.Espacio.Deporte); err != nil {
			logs.Logger.Println("Error al escanear el horario de la loza deportiva del socio: ", err)
			return nil, err
		}

		reservasSocio = append(reservasSocio, reserva)
	}
	return reservasSocio, nil
}
