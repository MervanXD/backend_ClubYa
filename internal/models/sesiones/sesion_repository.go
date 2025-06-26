package sesiones

type SesionRepository interface {
	ObtenerSesionesGrupo(idGrupo int) ([]Sesion, error)
	InsertarSesion(sesion *Sesion) (int64, error)
	ActualizarSesionParcial(sesion *SesionUpdate) error
}
