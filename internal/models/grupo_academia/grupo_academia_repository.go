package grupoacademia

import "database/sql"

type GrupoAcademiaRepository interface {
	ObtenerGruposAcademiaPorId(idAcademia int) ([]GrupoAcademia, error)
	InsertarGrupoAcademiaTx(tx *sql.Tx, grupo *GrupoAcademia) (int64, error)
	ActualizarGrupoAcademiaParcial(grupo *GrupoAcademiaUpdate) error
}
