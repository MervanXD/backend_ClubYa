package detalledisponibilidad

import "github.com/MervanXD/backend_ClubYa/internal/models/tipos"

type DetalleRequestActualizar struct {
	IdHorarioDia         int                        `json:"id_horario_dia"`
	IdBloqueTiempo       int                        `json:"id_bloque_tiempo"`
	EstadoDisponibilidad tipos.EstadoDisponibilidad `json:"estado_disponibilidad"`
	Fecha                string                     `json:"fecha"`
	Dia                  tipos.Dia                  `json:"dia"`
	Id_Espacio           int                        `json:"id_espacio"`
}
