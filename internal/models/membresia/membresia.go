package membresia

import (
	"time"

	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type SolicitudMembresia struct {
	Id          int                   `json:"id"`
	FechaInicio time.Time             `json:"fecha_inicio"`
	FechaFin    time.Time             `json:"fecha_fin"`
	Tipo        tipos.TipoMembresia   `json:"tipo_membresia"`
	Estado      tipos.EstadoMembresia `json:"estado_membresia"`
}
