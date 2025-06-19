package reserva

import (
	"github.com/MervanXD/backend_ClubYa/database"
	detalledisponibilidad "github.com/MervanXD/backend_ClubYa/internal/models/detalle_disponibilidad"
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type reservaRepositoryDB struct{}

func NewReservaRepositoryDB() ReservaRepository {
	return &reservaRepositoryDB{}
}

func (r *reservaRepositoryDB) ReservarEspacio(idEspacio int, idHorarioDia int, idBloqueTiempo int) (err error) {
	repo := detalledisponibilidad.NewDetalleDisponibilidadRepositoryDB()
	err = repo.ActualizarEstadoDetalleDisponibilidad(idHorarioDia, idBloqueTiempo, tipos.Reservado.String())
	if err != nil {
		logs.Logger.Println("Error al ReservarEspacio: ", err)
		return err
	}
	return nil
}

func (r *reservaRepositoryDB) ReservarEspacioSocial(reserva ReservaEspacio) error {
	query := "call ingesoft.ReservarEspacioSocial(?, ?, ?, ?, ?, ?,?)"
	_, err := database.DB.Exec(query, reserva.IdSocio, reserva.Espacio.Id, reserva.IdHorarioDia, reserva.IdBloqueTiempo, reserva.Fecha, reserva.HoraInicio, reserva.HoraFin)

	if err != nil {
		logs.Logger.Println("Error al reservar el espacio social: ", err)
		return err
	}
	repo := detalledisponibilidad.NewDetalleDisponibilidadRepositoryDB()
	err = repo.ActualizarEstadoDetalleDisponibilidad(reserva.IdHorarioDia, reserva.IdBloqueTiempo, tipos.Reservado.String())
	if err != nil {
		logs.Logger.Println("Error al ReservarEspacio: ", err)
		return err
	}
	return nil
}

func (r *reservaRepositoryDB) AnulacionReservaEspacioSocial(idReserva int, idEspacio int, idHorarioDia int, idBloque int, motivo string) error {
	query := "call ingesoft.AnularReservaEspacioSocial(?, ?, ?, ?, ?)"
	_, err := database.DB.Exec(query, idReserva, idEspacio, idHorarioDia, idBloque, motivo)
	if err != nil {
		logs.Logger.Println("Error al cancelar la reserva del espacio social ", err)
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
