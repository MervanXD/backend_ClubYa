package bloque_tiempo

type BloqueTiempoRepository interface {
	ObtenerBloquesTiempoEspacio(idEspacio int) ([]BloqueTiempo, error)
	ListarBloquesTiempoEstandar() ([]BloqueTiempo, error)
}
