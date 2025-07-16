package pago

import "github.com/MervanXD/backend_ClubYa/internal/models/tipos"

type PagoMultipleRequest struct {
	IdTitular  int              `json:"id_titular" validate:"required"`
	MetodoPago tipos.MetodoPago `json:"metodo_pago" validate:"required"` // "Tarjeta" o "Voucher"
}
