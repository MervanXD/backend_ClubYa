package espacio

type ReporteCanchaDTO struct {
    // Métricas generales
    Metricas ReporteCanchaMetricas `json:"metricas"`
    // Datos por socio
    DatosSocios []ReporteCanchaSocioDTO `json:"datos_socios"`
}

type ReporteCanchaMetricas struct {
    IngresosTotales      float64 `json:"ingresos_totales"`
    HorasReservadas      int     `json:"horas_reservadas"`
    ReservasConfirmadas  int     `json:"reservas_confirmadas"`
    ReservasAnuladas     int     `json:"reservas_anuladas"`
}

type ReporteCanchaSocioDTO struct {
    NombreSocio         string  `json:"nombre_socio"`
    ReservasConfirmadas int     `json:"reservas_confirmadas"`
    ReservasAnuladas    int     `json:"reservas_anuladas"`
    HorasReservadas     int     `json:"horas_reservadas"`
    IngresosPorSocio    float64 `json:"ingresos_por_socio"`
}