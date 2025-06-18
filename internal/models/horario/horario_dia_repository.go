package horario

type HorarioDiaRepository interface {
	ObtenerInscritosEspacioFecha(idEspacio int, fecha string) ([]HorarioDiaDTO, error)
}
