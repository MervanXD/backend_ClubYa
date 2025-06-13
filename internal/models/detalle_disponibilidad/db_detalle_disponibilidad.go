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

func ActualizarEstadoDetalleDisponibilidad(idHorarioDia int, idBloqueTiempo int, estado string) (err error) {
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

func ObtenerDisponibilidadEspacioSocialPorId(ctx context.Context, idEspacio int, idHorarioDia int, idBloqueTiempo int) (*DisponibilidadEspacioResponse, error) {
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

func ObtenerRangosInicioDisponibles(idEspacio int, fecha string) ([]string, error) {
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
