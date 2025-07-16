package tarifas

type TarifaMoraRepository interface {
	InsertarTarifaMora(tarifa TarifaMora) error
	ObtenerTarifaMoraPorMembresia(id int64) (TarifaMora, error)
	ListarTarifaMora() ([]TarifaMora, error)
	ModificarTarifaMora(tarifa TarifaMora) error
}
