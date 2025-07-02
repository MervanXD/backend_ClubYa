package configuracion_disponibilidad

import (
	"context"
	"database/sql"

	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
)

type configuracionDisponibilidadRepositoryDB struct{}

func NewConfiguracionDisponibilidadRepositoryDB() ConfiguracionDisponibilidadRepository {
	return &configuracionDisponibilidadRepositoryDB{}
}

func (r *configuracionDisponibilidadRepositoryDB) InsertarConfiguracionDisponibilidad(ctx context.Context, tx *sql.Tx, configuracion ConfiguracionDisponibilidad) error {
	err := utils.ExecSP(ctx, tx, "InsertarConfiguracionDisponibilidad", configuracion.IdEspacio, configuracion.IdBloqueTiempo, configuracion.Dia)
	if err != nil {
		return err
	}
	return nil
}

func (r *configuracionDisponibilidadRepositoryDB) ObtenerConfiguracionDisponibilidad(ctx context.Context, tx *sql.Tx, idEspacio int) ([]ConfiguracionDisponibilidad, error) {
	rows, err := utils.QuerySP(ctx, tx, "ObtenerConfiguracionDisponibilidad", idEspacio)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var configuraciones []ConfiguracionDisponibilidad
	for rows.Next() {
		var configuracion ConfiguracionDisponibilidad
		err = rows.Scan(&configuracion.IdEspacio, &configuracion.IdBloqueTiempo, &configuracion.Dia)
		if err != nil {
			return nil, err
		}
		configuraciones = append(configuraciones, configuracion)
	}
	return configuraciones, nil
}

func (r *configuracionDisponibilidadRepositoryDB) EliminarConfiguracionDisponibilidad(ctx context.Context, tx *sql.Tx, configuracion ConfiguracionDisponibilidad) error {
	err := utils.ExecSP(ctx, tx, "EliminarConfiguracionDisponibilidad", configuracion.IdEspacio, configuracion.IdBloqueTiempo, configuracion.Dia)
	if err != nil {
		return err
	}
	return nil
}
