package academia

type ReporteAcademiaRequest struct {
	MesInicio      int    `json:"mes_inicio"`
	MesFin         int    `json:"mes_fin"`
	Anio           int    `json:"anio"` //2025 o 2023 etc
	Deporte        string `json:"deporte"`
	OrdenIngreso   bool   `json:"orden_ingreso"`   //true es descendente, false ascendente
	OrdenInscritos bool   `json:"orden_inscritos"` //true es descendente, false ascendente
}
