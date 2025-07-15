package academia

type ReporteAcademiaDTO struct {
	NombreAcademia string  `json:"nombre_academia"`
	Deporte        string  `json:"deporte"`
	Entrenador     string  `json:"entrenador"`
	Inscritos      int     `json:"inscritos"`
	FechaInicio    string  `json:"fecha_inicio"`
	FechaFin       string  `json:"fecha_fin"`
	IngresoTotal   float64 `json:"ingreso_total"`
}
