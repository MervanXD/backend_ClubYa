package evento

type ReporteEventoDTO struct {
    NombreEvento   string  `json:"nombre_evento"`
    FechaEvento    string  `json:"fecha_evento"`
    Duracion       string  `json:"duracion"`
    NroInscritos   int     `json:"nro_inscritos"`
    IngresoTotal   float64 `json:"ingreso_total"`
}