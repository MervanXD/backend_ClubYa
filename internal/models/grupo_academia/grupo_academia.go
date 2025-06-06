package grupoacademia

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/espacio"
	"github.com/MervanXD/backend_ClubYa/internal/models/sesiones"
)

type GrupoAcademia struct {
	ID         int               `json:"id"`
	Nombre     string            `json:"nombre"`
	Vacantes   int               `json:"vacantes"`
	EdadMinima int               `json:"edad_minima"`
	EdadMaxima int               `json:"edad_maxima"`
	Espacio    espacio.Espacio   `json:"espacio"`
	Sesiones   []sesiones.Sesion `json:"sesiones"`
}
