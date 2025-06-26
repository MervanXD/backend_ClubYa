package academia

import (
	grupoacademia "github.com/MervanXD/backend_ClubYa/internal/models/grupo_academia"
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type AcademiaUpdateDTO struct {
	Nombre         *string                              `json:"nombre,omitempty"`
	Descripcion    *string                              `json:"descripcion,omitempty"`
	Deporte        *tipos.Deporte                       `json:"deporte,omitempty"`
	Entrenador     *string                              `json:"entrenador,omitempty"`
	CostoUniforme  *float64                             `json:"costo_uniforme,omitempty"`
	CostoMatricula *float64                             `json:"costo_matricula,omitempty"`
	Reglamento     *string                              `json:"reglamento,omitempty"`
	Imagen         *string                              `json:"imagen,omitempty"`
	Indicaciones   *string                              `json:"indicaciones,omitempty"`
	FechaInicio    *string                              `json:"fecha_inicio,omitempty"` // Formato: "YYYY-MM-DD"
	FechaFin       *string                              `json:"fecha_fin,omitempty"`    // Formato: "YYYY-MM-DD"
	Grupos         *[]grupoacademia.GrupoAcademiaUpdate `json:"grupos,omitempty"`       // IDs de los grupos asociados a la academia
}
