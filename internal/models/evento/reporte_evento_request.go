package evento

type ReporteEventoRequest struct {
	FechaInicio  string `json:"fecha_inicio" validate:"required"`
	FechaFin     string `json:"fecha_fin" validate:"required"`
	OrdenIngreso bool   `json:"orden_ingreso" validate:"required"` // true es descendente, false ascendente
	OrdenAsistentes bool `json:"orden_asistentes" validate:"required"` // true es descendente, false ascendente
}