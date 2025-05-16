package inscripcionevento

import (
	"time"

	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type InscripcionEvento struct {
	IdInscripcionEvento int          `json:"id_inscripcion_evento"`
	FechaInscripcion    time.Time    `json:"fecha_inscripcion"`
	HoraInscripcion     time.Time    `json:"hora_inscripcion"`
	Estado              tipos.Estado `json:"estado"`
}
