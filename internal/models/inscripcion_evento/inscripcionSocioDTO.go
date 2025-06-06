package inscripcion_evento

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/evento"
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type InscripcionSocioDTO struct {
	IdInscripcionEvento int           `json:"id_inscripcion_evento"`
	FechaInscripcion    string        `json:"fecha_inscripcion"`
	Estado              tipos.Estado  `json:"estado"`
	Evento              evento.Evento `json:"evento"`
}
