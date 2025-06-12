package academia

type AcademiaRepository interface {
	ObtenerAcademias() ([]AcademiaDTO, error)
	ObtenerAcademiaPorId(idAcademia int) (*Academia, error)
}