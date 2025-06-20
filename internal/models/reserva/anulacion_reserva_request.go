package reserva

type AnulacionReservaRequest struct {
	IdReserva            int     `json:"id_reserva"`
	IdAnulacion          int     `json:"id_anulacion"`
	PorcentajeDevolucion float64 `json:"porcentaje_devolucion"`
}
