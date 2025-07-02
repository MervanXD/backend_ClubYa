package configuracion_disponibilidad

import (
	"context"
	"database/sql"
	"errors"
	"slices"

	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type configuracionDisponibilidadRepositoryDB struct{}

func NewConfiguracionDisponibilidadRepositoryDB() ConfiguracionDisponibilidadRepository {
	return &configuracionDisponibilidadRepositoryDB{}
}

func (r *configuracionDisponibilidadRepositoryDB) InsertarConfiguracionDisponibilidad(ctx context.Context, tx *sql.Tx, configuracion ConfiguracionDisponibilidad) error {
	err := utils.ExecSP(ctx, tx, "InsertarConfiguracionDisponibilidad", configuracion.IdEspacio, configuracion.IdBloqueTiempo, configuracion.Dia)
	if err != nil {
		logs.Logger.Println("Error al insertar la configuracion de disponibilidad: ", err)
		return err
	}
	return nil
}

func (r *configuracionDisponibilidadRepositoryDB) ObtenerConfiguracionDisponibilidad(ctx context.Context, idEspacio int, tx ...*sql.Tx) ([]ConfiguracionDisponibilidad, error) {
	var queryer interface{}
	if len(tx) > 0 && tx[0] != nil {
		queryer = tx[0]
	} else {
		queryer = database.DB
	}
	rows, err := utils.QuerySP(ctx, queryer, "ObtenerConfiguracionDisponibilidad", idEspacio)
	if err != nil {
		logs.Logger.Println("Error al obtener la configuracion de disponibilidad: ", err)
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
		logs.Logger.Println("Error al eliminar la configuracion de disponibilidad: ", err)
		return err
	}
	return nil
}

func (r *configuracionDisponibilidadRepositoryDB) CrearConfiguracionDisponibilidad(ctx context.Context, configuraciones []ConfiguracionDisponibilidad) error {
	tx, err := database.DB.BeginTx(ctx, nil)
	if err != nil {
		logs.Logger.Println("Error al iniciar la transaccion: ", err)
		return err
	}
	defer tx.Rollback()
	for _, configuracion := range configuraciones {
		err := r.InsertarConfiguracionDisponibilidad(ctx, tx, configuracion)
		if err != nil {
			logs.Logger.Println("Error al insertar la configuracion de disponibilidad: ", err)
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		logs.Logger.Println("Error al confirmar la transaccion: ", err)
		return err
	}
	return nil
}

func (r *configuracionDisponibilidadRepositoryDB) ActualizarConfiguracionDisponibilidad(ctx context.Context, configuraciones []ConfiguracionDisponibilidad) error {
	if len(configuraciones) == 0 {
		return errors.New("no se proporcionaron configuraciones")
	}

	tx, err := database.DB.BeginTx(ctx, nil)
	if err != nil {
		logs.Logger.Println("Error al iniciar la transacción: ", err)
		return err
	}
	defer tx.Rollback()

	idEspacio := configuraciones[0].IdEspacio

	configuracionesAnteriores, err := r.ObtenerConfiguracionDisponibilidad(ctx, idEspacio, tx)
	if err != nil {
		logs.Logger.Println("Error al obtener las configuraciones anteriores: ", err)
		return err
	}

	// Eliminar las configuraciones que ya no existen
	for _, anterior := range configuracionesAnteriores {
		if !slices.Contains(configuraciones, anterior) {
			err := r.EliminarConfiguracionDisponibilidad(ctx, tx, anterior)
			if err != nil {
				logs.Logger.Println("Error al eliminar configuración: ", err)
				return err
			}
		}
	}

	// Insertar las configuraciones que no existen
	for _, nueva := range configuraciones {
		if !slices.Contains(configuracionesAnteriores, nueva) {
			err := r.InsertarConfiguracionDisponibilidad(ctx, tx, nueva)
			if err != nil {
				logs.Logger.Println("Error al insertar configuración: ", err)
				return err
			}
		}
	}

	if err := tx.Commit(); err != nil {
		logs.Logger.Println("Error al confirmar transacción: ", err)
		return err
	}

	return nil
}
