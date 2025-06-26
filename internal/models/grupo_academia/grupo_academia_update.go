package grupoacademia

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/sesiones"
	"github.com/MervanXD/backend_ClubYa/internal/models/tarifas"
)

type GrupoAcademiaUpdate struct {
	ID         *int                            `json:"id"`
	IdAcademia *int                            `json:"id_academia"`
	Nombre     *string                         `json:"nombre"`
	Vacantes   *int                            `json:"vacantes"`
	EdadMinima *int                            `json:"edad_minima"`
	EdadMaxima *int                            `json:"edad_maxima"`
	Espacio    *int                            `json:"espacio"` // ID del espacio
	Inscritos  *int                            `json:"cantidad_inscritos"`
	EsVigente  *bool                           `json:"es_vigente"` // Indica si el grupo está vigente
	Sesiones   *[]sesiones.SesionUpdate        `json:"sesiones"`   // IDs de las sesiones
	Tarifas    *[]tarifas.TarifaAcademiaUpdate `json:"tarifas"`    // IDs de las tarifas
}
