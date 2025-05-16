package detalledisponibilidad

import "github.com/MervanXD/backend_ClubYa/internal/models/tipos"

type DetalleDisponibilidad struct {
	IdHorarioDia         int                        `json:"id_horario_dia"`
	IdBloqueTiempo       int                        `json:"id_bloque_tiempo"`
	EstadoDisponibilidad tipos.EstadoDisponibilidad `json:"estado_disponibilidad"`
}
