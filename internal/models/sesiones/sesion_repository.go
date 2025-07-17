package sesiones

import "database/sql"

type SesionRepository interface {
	ObtenerSesionesGrupo(idGrupo int) ([]Sesion, error)
	InsertarSesionTx(tx *sql.Tx, sesion *Sesion) (int64, error)
	ActualizarSesionParcial(sesion *SesionUpdate) error
}
