package reserva

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/espacio"
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type ReservaEspacio struct {
	Id             int             `json:"id"`
	Espacio        espacio.Espacio `json:"espacio"`
	HoraInicio     string          `json:"horaInicio"`
	HoraFin        string          `json:"horaFin"`
	Fecha          string          `json:"fecha"`
	FechaReserva   string          `json:"fechaReserva"`
	Estado         tipos.Estado    `json:"estado"`
	IdHorarioDia   int             `json:"id_horario_dia"`
	IdBloqueTiempo int             `json:"id_bloque_tiempo"`
	IdSocio        int             `json:"id_persona"`
	//AnulacionReserva *AnulacionReserva `json:"anulacionReserva,omitempty"`
}
