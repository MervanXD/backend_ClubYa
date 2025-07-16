package espacio

type CanchaRepository interface {
	ObtenerCanchasHorarios() ([]CanchaHorarioDTO, error)
	InsertarCancha(c Cancha) error
	ActualizarParcial(id int, dto CanchaUpdateDTO) error
	ObtenerCanchasConfiguracion() ([]Cancha, error)
	GenerarReporteCanchas(filtros ReporteCanchaRequest) (ReporteCanchaDTO, error)
}
