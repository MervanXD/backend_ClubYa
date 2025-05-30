package espacio

import (
	"time"

	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type CanchaHorarioDTO struct {
	Espacio    Cancha              	      `json:"espacio"`
	HoraInicio string                     `json:"hora_inicio"`
	HoraFinal  string                     `json:"hora_fin"`
	Fecha      time.Time                  `json:"fecha"`
	Estado     tipos.EstadoDisponibilidad `json:"disponibilidad"`
	IdHorario  int                        `json:"id_horario"`
	IdBloque   int                        `json:"id_bloque"`
}
