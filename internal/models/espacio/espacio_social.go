package espacio

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type EspacioSocial struct {
	Espacio
	Actividad tipos.Actividad `json:"actividad"`
}
