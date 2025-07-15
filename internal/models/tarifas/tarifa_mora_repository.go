package tarifas

type TarifaMoraRepository interface {
	InsertarTarifaMora(tarifa TarifaMora) error
	ListarTarifaMora() ([]TarifaMora, error)
	ModificarTarifaMora(tarifa TarifaMora) error
}
