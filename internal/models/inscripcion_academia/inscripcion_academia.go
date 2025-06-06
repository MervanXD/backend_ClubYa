package inscripcionacademia

import (
	"time"

	"github.com/MervanXD/backend_ClubYa/internal/models/academia"
	grupoacademia "github.com/MervanXD/backend_ClubYa/internal/models/grupo_academia"
	"github.com/MervanXD/backend_ClubYa/internal/models/persona"
	"github.com/MervanXD/backend_ClubYa/internal/models/tarifas"
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type InscripcionAcademia struct {
	IdInscripcionAcademia int                         `json:"id_inscripcion_academia"`
	FechaInscripcion      time.Time                   `json:"fecha_inscripcion"`
	EstadoInscripcion     tipos.Estado                `json:"estado"`
	AcademiaInscrita      academia.Academia           `json:"academia"`
	SocioInscrito         persona.Persona             `json:"socio"`
	GrupoInscrito         grupoacademia.GrupoAcademia `json:"grupo"`
	Tarifa                tarifas.TarifaAcademia      `json:"tarifa"`
}
