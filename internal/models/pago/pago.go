package pago

import (
	"time"

	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type Pago struct {
	IdPago   int              `json:"id_pago"`
	Fecha    time.Time        `json:"fecha"`
	Concepto string           `json:"concepto"`
	Metodo   tipos.MetodoPago `json:"metodo"`
	Monto    float64          `json:"monto"`
}
