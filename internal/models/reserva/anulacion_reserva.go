package reserva

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
)

type AnulacionReserva struct {
	Id         int              `json:"id"`
	ReservaId  int              `json:"reserva_id"`
	Fecha      utils.NullString `json:"fecha"`
	Motivo     utils.NullString `json:"motivo"`
	Tipo       tipos.Tipo       `json:"tipo_reserva"`
	Devolucion float64          `json:"devolucion,omitempty"` // Monto a devolver al socio, si aplica
}
