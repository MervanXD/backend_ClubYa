package tarifas

import "database/sql"

type TarifaAcademiaRepository interface {
	ObtenerTarifasAcademiaPorId(idGrupoAcademia int) ([]TarifaAcademia, error)
	InsertarTarifaAcademiaTx(tx *sql.Tx, tarifa *TarifaAcademia) (int64, error)
	ActualizarTarifaAcademia(tarifa *TarifaAcademiaUpdate) error
}
