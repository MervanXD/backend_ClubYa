package espacio

type ReporteCanchaRequest struct {
	FechaInicio  string `json:"fecha_inicio" validate:"required"`
	FechaFin     string `json:"fecha_fin" validate:"required"`
	CanchaID     int    `json:"cancha_id" validate:"required"`
	OrdenIngreso bool   `json:"orden_ingreso" validate:"required"`//true es descendente, false ascendente
	OrdenHoras   bool   `json:"orden_horas" validate:"required"`//true es descendente, false ascendente
}
