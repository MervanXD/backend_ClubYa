package reserva

import (
	"time"

	"github.com/MervanXD/backend_ClubYa/internal/models/espacio"
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type ReservaEspacio struct {
	Id           int             `json:"id"`
	Espacio      espacio.Espacio `json:"espacio"`
	HoraInicio   time.Time       `json:"horaInicio"`
	HoraFin      time.Time       `json:"horaFin"`
	Fecha        time.Time       `json:"fecha"`
	FechaReserva time.Time       `json:"fechaReserva"`
	Estado       tipos.Estado    `json:"estado"`
	//AnulacionReserva *AnulacionReserva `json:"anulacionReserva,omitempty"`
}
