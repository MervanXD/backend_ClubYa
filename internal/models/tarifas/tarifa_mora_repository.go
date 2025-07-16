package tarifas

type TarifaMoraRepository interface {
	InsertarTarifaMora(tarifa TarifaMora) error
	ObtenerTarifaMoraPorId(id int64) (TarifaMora, error)
	ListarTarifaMora() ([]TarifaMora, error)
	ModificarTarifaMora(tarifa TarifaMora) error
}
