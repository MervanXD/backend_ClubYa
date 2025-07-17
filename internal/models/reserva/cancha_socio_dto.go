package reserva

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/espacio"
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type ReservaCanchaRequest struct {
	Id           int            `json:"id"`
	Espacio      espacio.Cancha `json:"espacioSocial"`
	HoraInicio   string         `json:"horaInicio"`
	HoraFin      string         `json:"horaFin"`
	Fecha        string         `json:"fecha"`
	FechaReserva string         `json:"fechaReserva"`
	Estado       tipos.Estado   `json:"estado"`
	IdHorarioDia int            `json:"id_horario_dia"`
	IdSocio      int            `json:"id_persona"`
}
