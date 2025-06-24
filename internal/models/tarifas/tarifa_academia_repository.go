package tarifas

type TarifaAcademiaRepository interface {
	ObtenerTarifasAcademiaPorId(idGrupoAcademia int) ([]TarifaAcademia, error)
	InsertarTarifaAcademia(tarifa *TarifaAcademia) (int64, error)
}
