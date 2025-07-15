package tarifas

type TarifaMiembroAdicionalRepository interface {
	InsertarTarifaMiembroAdicional(tarifa TarifaMiembroAdicional) error
	ListarTarifaMiembroAdicional() ([]TarifaMiembroAdicional, error)
	ModificarTarifaMiembroAdicional(idTarifa int, monto float64) error
}
