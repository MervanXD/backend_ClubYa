package configuracion_disponibilidad

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type configuracionDisponibilidadRepositoryDB struct{}

func NewConfiguracionDisponibilidadRepositoryDB() ConfiguracionDisponibilidadRepository {
	return &configuracionDisponibilidadRepositoryDB{}
}

func (r *configuracionDisponibilidadRepositoryDB) InsertarConfiguracionDisponibilidad(ctx context.Context, tx *sql.Tx, configuracion ConfiguracionDisponibilidad) error {
	err := utils.ExecSP(ctx, tx, "InsertarConfiguracionDisponibilidad", configuracion.IdEspacio, configuracion.IdBloqueTiempo, configuracion.Dia.String())
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
	rows, err := utils.QuerySP(ctx, queryer, "ObtenerConfiguracionDeEspacio", idEspacio)
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
	err := utils.ExecSP(ctx, tx, "EliminarConfiguracionDisponibilidad", configuracion.IdEspacio, configuracion.IdBloqueTiempo, configuracion.Dia.String())
	if err != nil {
		logs.Logger.Println("Error al eliminar la configuracion de disponibilidad: ", err)
		return err
	}
	return nil
}

type ConfiguracionDisponibilidadDTO struct {
	ConfiguracionDisponibilidad
	Eliminada bool `json:"eliminada"`
}

func (r *configuracionDisponibilidadRepositoryDB) ActualizarConfiguracionDisponibilidad(ctx context.Context, configuraciones []ConfiguracionDisponibilidadDTO) error {
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

	cruces, err := r.ObtenerCrucesEspacio(ctx, idEspacio)
	if err != nil {
		logs.Logger.Println("Error al obtener los cruces del espacio: ", err)
		return err
	}

	crucesMap := make(map[string]bool)
	for _, cruce := range cruces {
		cruceKey := fmt.Sprintf("%s_%d", cruce.Dia, cruce.IdBloqueTiempo)
		crucesMap[cruceKey] = true
	}

	var paraInsertar, paraEliminar []ConfiguracionDisponibilidad
	for _, d := range configuraciones {
		cfg := ConfiguracionDisponibilidad{
			IdEspacio:      d.IdEspacio,
			Dia:            d.Dia,
			IdBloqueTiempo: d.IdBloqueTiempo,
		}
		if d.Eliminada {
			cruceKey := fmt.Sprintf("%s_%d", d.Dia.String(), d.IdBloqueTiempo)
			if crucesMap[cruceKey] {
				logs.Logger.Println("No se puede eliminar un bloque de tiempo que ya tiene inscritos")
				return errors.New("no se puede eliminar un bloque de tiempo que ya tiene inscritos")
			}
			paraEliminar = append(paraEliminar, cfg)
		} else {
			paraInsertar = append(paraInsertar, cfg)
		}
	}

	for _, cfg := range paraEliminar {
		err := r.EliminarConfiguracionDisponibilidad(ctx, tx, cfg)
		if err != nil {
			logs.Logger.Println("Error al eliminar configuración: ", err)
			return err
		}
	}

	for _, cfg := range paraInsertar {
		err := r.InsertarConfiguracionDisponibilidad(ctx, tx, cfg)
		if err != nil {
			logs.Logger.Println("Error al insertar configuración: ", err)
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		logs.Logger.Println("Error al confirmar transacción: ", err)
		return err
	}

	return nil
}

type NroCruces struct {
	Dia            string `json:"dia"`
	IdBloqueTiempo int    `json:"id_bloque_tiempo"`
	Cantidad       int    `json:"cantidad"`
}

func (r *configuracionDisponibilidadRepositoryDB) ObtenerCrucesEspacio(ctx context.Context, idEspacio int) ([]NroCruces, error) {
	rows, err := utils.QuerySP(ctx, database.DB, "ObtenerNroCrucesPorDiaBloque", idEspacio)
	if err != nil {
		logs.Logger.Println("Error al obtener los inscritos del espacio: ", err)
		return nil, err
	}
	defer rows.Close()

	var inscritos []NroCruces
	for rows.Next() {
		var inscrito NroCruces
		err = rows.Scan(&inscrito.Dia, &inscrito.IdBloqueTiempo, &inscrito.Cantidad)
		if err != nil {
			return nil, err
		}
		inscritos = append(inscritos, inscrito)
	}
	return inscritos, nil
}
