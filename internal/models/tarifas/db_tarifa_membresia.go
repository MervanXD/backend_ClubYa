package tarifas

import (
	"context"
	"database/sql"

	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type tarifaMembresiaRepositoryDB struct{}

func NewTarifaMembresiaRespositoryDB() TarifaMembresiaRepository {
	return &tarifaMembresiaRepositoryDB{}
}

func (r *tarifaMembresiaRepositoryDB) ObtenerTarifasMembresiaPorId(id int64) (TarifaMembresia, error) {
	query := "CALL ObtenerTarifasMembresiaPorId(?)"
	rows, err := database.DB.Query(query, id)
	if err != nil {
		logs.Logger.Println("Error al obtener las tarifas de membresía:", err)
		return TarifaMembresia{}, err
	}
	defer rows.Close()

	var tarifa TarifaMembresia
	err = rows.Scan(&tarifa.IdTarifaMembresia, &tarifa.FechaInicio, &tarifa.CuotaBase, &tarifa.DiasPlazoPago, &tarifa.MetodoPago, &tarifa.MontoPagoInicial)
	if err != nil {
		logs.Logger.Println("Error al escanear la tarifa de membresía:", err)
		return TarifaMembresia{}, err
	}
	return tarifa, nil
}

func (r *tarifaMembresiaRepositoryDB) InsertarTarifaMembresiaTx(ctx context.Context, tx *sql.Tx, tarifa *TarifaMembresia) (int64, error) {
	var idTarifa int64
	err := utils.ExecSPWithOut(ctx, tx, "InsertarTarifaMembresia", "idTarifaMembresia", &idTarifa,
		tarifa.FechaInicio, tarifa.CuotaBase, tarifa.DiasPlazoPago, "Tarjeta", tarifa.MontoPagoInicial)
	if err != nil {
		logs.Logger.Println("Error al insertar la tarifa de membresía:", err)
		return 0, err
	}

	return idTarifa, nil
}

func (r *tarifaMembresiaRepositoryDB) ActualizarTarifaMembresia(ctx context.Context, tx *sql.Tx, tarifa *TarifaMembresia) error {
	err := utils.ExecSP(ctx, tx, "ActualizarTarifaMembresia",
		tarifa.IdTarifaMembresia, tarifa.FechaFin, tarifa.CuotaBase, tarifa.DiasPlazoPago, tarifa.MontoPagoInicial)
	if err != nil {
		logs.Logger.Println("Error al actualizar la tarifa de membresía:", err)
		return err
	}
	return nil
}
