package pago

type PagoMultipleResponse struct {
    TotalPagos     int     `json:"total_pagos"`
    IdsPagos       []int   `json:"ids_pagos"`
    MontoTotal     float64 `json:"monto_total,omitempty"` // Si el procedure lo devuelve
    PagosAfectados int     `json:"pagos_afectados"`
}