package reserva

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/espacio"
	"github.com/MervanXD/backend_ClubYa/internal/models/horario"
	"github.com/MervanXD/backend_ClubYa/internal/models/pago"
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type ReservaEspacio struct {
	Id           int                `json:"id"`
	Pago         pago.Pago          `json:"pago"`
	Espacio      espacio.Espacio    `json:"espacio"`
	HorarioDia   horario.HorarioDia `json:"horario_dia"`
	HoraInicio   string             `json:"hora_inicio"`
	HoraFin      string             `json:"hora_fin"`
	FechaReserva string             `json:"fecha_reserva"`
	Fecha        string             `json:"fecha"`
	Estado       tipos.Estado       `json:"estado"`
	IdSocio      int                `json:"id_socio"`
	//AnulacionReserva *AnulacionReserva `json:"anulacionReserva,omitempty"`
}
