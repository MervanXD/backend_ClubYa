package academia

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/espacio"
)

type GrupoAcademia struct {
	ID         int64           `json:"id"`
	Nombre     string          `json:"nombre"`
	Vacantes   int64           `json:"vacantes"`
	EdadMinima int64           `json:"edad_minima"`
	EdadMaxima int64           `json:"edad_maxima"`
	HoraInicio string          `json:"hora_inicio"`
	HoraFin    string          `json:"hora_fin"`
	Espacio    espacio.Espacio `json:"espacio"`
}
