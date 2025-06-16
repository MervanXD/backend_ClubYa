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

func (r *detalleDisponibilidadRepositoryDB) ActualizarEstadoDetalleDisponibilidad(idHorarioDia int, idBloqueTiempo int, estado string) (err error) {
	stmt := "call ActualizarEstadoDetalleDisponibilidad(?,?,?)"
	result, err := database.DB.Exec(stmt, idHorarioDia, idBloqueTiempo, estado)
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

func (r *detalleDisponibilidadRepositoryDB) ObtenerDetalleDisponibilidadEspacioSocialFechaId(idEspacio int, fecha string) ([]DetalleDisponibilidad, error) {
	query := "call ingesoft.ListarDetalleDisponibilidad(?,?)"
	rows, err := database.DB.Query(query, idEspacio, fecha)
	if err != nil {
		logs.Logger.Println("Error al obtener los de detalles tiempo: ", err)
		return nil, err
	}
	defer rows.Close()
	var detalles []DetalleDisponibilidad
	for rows.Next() {
		var dp DetalleDisponibilidad
		if err := rows.Scan(&dp.IdHorarioDia, &dp.IdBloqueTiempo, &dp.EstadoDisponibilidad); err != nil {
			logs.Logger.Println("Error al escanear el detalle: ", err)
			return nil, err
		}
		detalles = append(detalles, dp)
	}
	return detalles, nil
}
