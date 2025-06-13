package espacio

type CanchaRepository interface {
	ObtenerCanchasHorarios() ([]CanchaHorarioDTO, error)
}
