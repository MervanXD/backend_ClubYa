package grupoacademia

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/espacio"
	"github.com/MervanXD/backend_ClubYa/internal/models/sesiones"
	"github.com/MervanXD/backend_ClubYa/internal/models/tarifas"
)

type GrupoAcademia struct {
	ID         int                      `json:"id"`
	Nombre     string                   `json:"nombre"`
	Vacantes   int                      `json:"vacantes"`
	EdadMinima int                      `json:"edad_minima"`
	EdadMaxima int                      `json:"edad_maxima"`
	Espacio    espacio.Espacio          `json:"espacio"`
	Inscritos  int                      `json:"cantidad_inscritos"`
	Sesiones   []sesiones.Sesion        `json:"sesiones"`
	Tarifas    []tarifas.TarifaAcademia `json:"tarifas"`
}
