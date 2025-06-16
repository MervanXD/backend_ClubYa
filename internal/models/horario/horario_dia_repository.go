package horario

type HorarioDiaRepository interface {
	ObtenerInscritosEspacioSocialFecha(idEspacio int, fecha string) ([]HorarioDiaDTO, error)
}
