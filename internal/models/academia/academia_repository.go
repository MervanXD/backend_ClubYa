package academia

type AcademiaRepository interface {
	ObtenerAcademias() ([]AcademiaDTO, error)
	ObtenerAcademiaPorId(idAcademia int) (*Academia, error)
	InsertarAcademia(academia *Academia) error
	ListarAcademiasGenerales() ([]AcademiaListarRequest, error)
	ActualizarParcialAcademia(id int, dto AcademiaUpdateDTO) error
}
