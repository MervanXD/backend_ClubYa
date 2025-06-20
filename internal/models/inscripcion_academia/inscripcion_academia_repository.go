package inscripcionacademia

type InscripcionAcademiaRepository interface {
	RegistrarInscripcionAcademia(idPersonaint int, idGrupo int, idTarifa int, uniforme int, costo_total float64, idTitular int, metodoPago string) error
	ObtenerFamiliaresInscritosAcademia(idSocio int) ([]InscritoAcademiaDTO, error)
}
