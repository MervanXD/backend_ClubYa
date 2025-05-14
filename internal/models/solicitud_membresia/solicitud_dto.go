package solicitud

import (
	"time"

	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type SolicitudDTO struct {
	Id        int                   `json:"id"`
	Fecha     time.Time             `json:"fecha"`
	Estado    tipos.EstadoSolicitud `json:"estado_solicitud"`
	IdTitular int                   `json:"id_titular"`
	Nombres   string                `json:"nombres"`
	Apellidos string                `json:"apellidos"`
}