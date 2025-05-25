package inscripcion_evento

import (
	"time"

	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type AnulacionInscripcionEvento struct {
	IdAnulacionInscripcionEvento int        `json:"id_anulacion_inscripcion_evento"`
	IdInscripcionEvento          int        `json:"id_inscripcion_evento"`
	FechaAnulacion               time.Time  `json:"fecha_anulacion"`
	Motivo                       string     `json:"motivo"`
	Tipo                         tipos.Tipo `json:"tipo"`
}
