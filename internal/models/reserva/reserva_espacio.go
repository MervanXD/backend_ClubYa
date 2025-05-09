package reserva

import (
	"time"

	"github.com/MervanXD/backend_ClubYa/internal/models/espacio"
)

type ReservaEspacio struct {
	Id         int             `json:"id"`
	Espacio    espacio.Espacio `json:"espacio"`
	HoraInicio time.Time       `json:"horaInicio"`
	HoraFin    time.Time       `json:"horaFin"`
	Fecha      time.Time       `json:"fecha"`
	Estado     string          `json:"estado"`
}
