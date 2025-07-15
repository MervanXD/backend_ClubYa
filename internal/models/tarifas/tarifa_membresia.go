package tarifas

import "github.com/MervanXD/backend_ClubYa/internal/models/tipos"

type TarifaMembresia struct {
	IdTarifaMembresia  int              `json:"id_tarifa_membresia"`
	FechaInicio        string           `json:"fecha_inicio"`
	FechaFin           string           `json:"fecha_fin"`
	CuotaBase          float64          `json:"cuota_base"`
	DiasPlazoPago      int              `json:"dias_plazo_pago"`
	MetodoPago         tipos.MetodoPago `json:"metodo_pago"`
	MontoPagoAdicional float64          `json:"monto_adicional"`
}
