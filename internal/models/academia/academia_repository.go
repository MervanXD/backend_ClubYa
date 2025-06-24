package academia

type AcademiaRepository interface {
	ObtenerAcademias() ([]AcademiaDTO, error)
	ObtenerAcademiaPorId(idAcademia int) (*Academia, error)
	InsertarAcademia(academia *Academia) error
}
