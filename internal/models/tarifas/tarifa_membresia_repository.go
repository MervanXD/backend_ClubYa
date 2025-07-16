package tarifas

import (
	"context"
	"database/sql"
)

type TarifaMembresiaRepository interface {
	ObtenerTarifasMembresiaPorId(id int64) (TarifaMembresia, error)
	InsertarTarifaMembresiaTx(ctx context.Context, tx *sql.Tx, tarifa *TarifaMembresia) (int64, error)
	ActualizarTarifaMembresia(ctx context.Context, tx *sql.Tx, tarifa *TarifaMembresia) error
}
