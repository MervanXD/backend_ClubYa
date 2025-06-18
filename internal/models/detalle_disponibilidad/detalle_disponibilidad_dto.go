package detalledisponibilidad

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type DetalleDisponibilidadDto struct {
	IdHorarioDia         int                        `json:"id_horario_dia"`
	IdBloqueTiempo       int                        `json:"id_bloque_tiempo"`
	EstadoDisponibilidad tipos.EstadoDisponibilidad `json:"estado_disponibilidad"`
	IdPersona            *int                       `json:"id_persona"`     //
	NombrePersona        *string                    `json:"nombre_persona"` //
}
