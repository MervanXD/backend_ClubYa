package espacio

type CanchaRepository interface {
	ObtenerCanchasHorarios() ([]CanchaHorarioDTO, error)
	ObtenerCanchasConfiguracion() ([]Cancha, error)
}
