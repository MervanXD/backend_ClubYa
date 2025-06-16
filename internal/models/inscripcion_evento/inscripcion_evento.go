package inscripcion_evento

import (
	"time"

	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type InscripcionEvento struct {
	IdInscripcionEvento int          `json:"id_inscripcion_evento"`
	IdEvento            int          `json:"id_evento"`
	IdSocio             int          `json:"id_socio"`
	FechaInscripcion    time.Time    `json:"fecha_inscripcion"`
	HoraInscripcion     time.Time    `json:"hora_inscripcion"`
	Estado              tipos.Estado `json:"estado"`
	CantidadInvitados   int          `json:"cantidad_invitados"`
}
