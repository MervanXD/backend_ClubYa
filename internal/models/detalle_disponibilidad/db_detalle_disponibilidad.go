package detalledisponibilidad

import (
	"context"
	"errors"
	"time"

	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/internal/models/espacio"
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type detalleDisponibilidadRepositoryDB struct{}

func NewDetalleDisponibilidadRepositoryDB() DetalleDisponibilidadRepository {
	return &detalleDisponibilidadRepositoryDB{}
}

func (r *detalleDisponibilidadRepositoryDB) ActualizarEstadoDetalleDisponibilidad(detalle DetalleRequestActualizar) (err error) {
	stmt := "call ActualizarEstadoDetalleDisponibilidad(?,?,?,?,?,?)"
	result, err := database.DB.Exec(stmt,
		detalle.IdHorarioDia,
		detalle.IdBloqueTiempo,
		detalle.EstadoDisponibilidad.String(),
		detalle.Id_Espacio,
		detalle.Fecha,
		detalle.Dia.String())
	if err != nil {
		logs.Logger.Println("Error al ActualizarEstadoDetalleDisponibilidad: ", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logs.Logger.Println("Error al obtener rowsAffected:", err)
		return err
	}

	if rowsAffected == 0 {
		return errors.New("columnas no fueron afectadas")
	}

	return nil
}

// a esta funcion no le voy a poner test pq es lo mismo que hacer lo de arriba solo que varias veces
func ActualizarDetalleGrupo(detalles []DetalleRequestActualizar) error {
	repo := NewDetalleDisponibilidadRepositoryDB()
	for _, detalle := range detalles {
		err := repo.ActualizarEstadoDetalleDisponibilidad(detalle)
		if err != nil {
			logs.Logger.Println("Error al actualizar el detalle de disponibilidad: ", err)
			return err
		}
	}
	return nil
}

type DisponibilidadEspacioResponse struct {
	Espacio        espacio.EspacioSocial      `json:"espacio"`
	HoraInicio     string                     `json:"hora_inicio"`
	HoraFin        string                     `json:"hora_fin"`
	Fecha          time.Time                  `json:"fecha"`
	Disponibilidad tipos.EstadoDisponibilidad `json:"disponibilidad"`
}

func (r *detalleDisponibilidadRepositoryDB) ObtenerDisponibilidadEspacioSocialPorId(ctx context.Context, idEspacio int, idHorarioDia int, idBloqueTiempo int) (*DisponibilidadEspacioResponse, error) {
	query := "call ingesoft.ObtenerDetalleDisponibilidadEspacioSocial(?,?,?)"

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var res DisponibilidadEspacioResponse
	row := database.DB.QueryRowContext(ctx, query, idEspacio, idHorarioDia, idBloqueTiempo)
	err := row.Scan(&res.Espacio.Id, &res.Espacio.Nombre, &res.Espacio.Codigo, &res.Espacio.Ubicacion, &res.Espacio.Capacidad, &res.Espacio.Costo, &res.Espacio.Imagen, &res.Espacio.Reglamento, &res.Espacio.Actividad,
		&res.Fecha, &res.HoraInicio, &res.HoraFin, &res.Disponibilidad)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *detalleDisponibilidadRepositoryDB) ObtenerDetalleDisponibilidadEspacioFechaId(idEspacio int, fecha string) ([]DetalleDisponibilidadDto, error) {
	query := "call ingesoft.ListarDetalleDisponibilidad(?,?)"
	rows, err := database.DB.Query(query, idEspacio, fecha)
	if err != nil {
		logs.Logger.Println("Error al obtener los de detalles tiempo: ", err)
		return nil, err
	}
	defer rows.Close()
	var detalles []DetalleDisponibilidadDto
	for rows.Next() {
		var dp DetalleDisponibilidadDto
		if err := rows.Scan(&dp.IdHorarioDia, &dp.IdBloqueTiempo, &dp.EstadoDisponibilidad, &dp.IdPersona, &dp.NombrePersona); err != nil {
			logs.Logger.Println("Error al escanear el detalle: ", err)
			return nil, err
		}

		detalles = append(detalles, dp)
	}

	return detalles, nil
}

func (r *detalleDisponibilidadRepositoryDB) ObtenerRangosInicioDisponibles(idEspacio int, fecha string) ([]string, error) {
	query := "CALL ListarHorariosDisponiblesPorEspacioYFecha(?, ?)"
	rows, err := database.DB.Query(query, idEspacio, fecha)
	if err != nil {
		logs.Logger.Println("Error al ejecutar el procedimiento: ", err)
		return nil, err
	}
	defer rows.Close()

	var rangos []string

	for rows.Next() {
		var rango string
		if err := rows.Scan(&rango); err != nil {
			logs.Logger.Println("Error al escanear fila: ", err)
			return nil, err
		}
		rangos = append(rangos, rango)
	}

	return rangos, nil
}

func (r *detalleDisponibilidadRepositoryDB) ActualizarDisponibilidadSegunReserva(detalle DetalleRequestActualizar) (int, error) {
	stmt := "call ActualizaDisponibilidadSegunReserva(?,?,?,?,?,?, @p_nuevoIdHorarioDia)"
	result, err := database.DB.Exec(stmt,
		detalle.IdHorarioDia,
		detalle.IdBloqueTiempo,
		detalle.EstadoDisponibilidad.String(),
		detalle.Id_Espacio,
		detalle.Fecha,
		detalle.Dia.String(),
	)
	if err != nil {
		logs.Logger.Println("Error al ActualizarEstadoSegunReserva: ", err)
		return -1, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logs.Logger.Println("Error al obtener rowsAffected:", err)
		return -1, err
	}

	if rowsAffected == 0 {
		return -1, errors.New("columnas no fueron afectadas")
	}

	// Recupera el valor del parámetro de salida
	var nuevoIdHorarioDia int
	row := database.DB.QueryRow("SELECT @p_nuevoIdHorarioDia")
	if err := row.Scan(&nuevoIdHorarioDia); err != nil {
		logs.Logger.Println("Error al obtener el nuevoIdHorarioDia: ", err)
		return -1, err
	}

	if nuevoIdHorarioDia == -1 {
		logs.Logger.Println("No se pudo obtener el nuevoIdHorarioDia")
		return -1, errors.New("no se pudo obtener el nuevoIdHorarioDia")
	}

	return nuevoIdHorarioDia, nil
}
