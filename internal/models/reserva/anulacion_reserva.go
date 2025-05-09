package reserva

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type AnulacionReserva struct {
	Id        int        `json:"id"`
	ReservaId int        `json:"reserva_id"`
	Fecha     string     `json:"fecha"`
	Motivo    string     `json:"motivo"`
	Tipo      tipos.Tipo `json:"tipo_reserva"`
}
