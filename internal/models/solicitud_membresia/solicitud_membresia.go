package solicitud

import (
	"time"

	"github.com/MervanXD/backend_ClubYa/internal/models/persona"
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type SolicitudMembresia struct {
	Titular persona.Titular
	Id      int                   `json:"id"`
	Fecha   time.Time             `json:"fecha"`
	Estado  tipos.EstadoSolicitud `json:"estado_solicitud"`
}
