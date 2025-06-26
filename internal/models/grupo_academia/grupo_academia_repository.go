package grupoacademia

type GrupoAcademiaRepository interface {
	ObtenerGruposAcademiaPorId(idAcademia int) ([]GrupoAcademia, error)
	InsertarGrupoAcademia(grupo *GrupoAcademia) (int64, error)
	ActualizarGrupoAcademiaParcial(grupo *GrupoAcademiaUpdate) error
}
