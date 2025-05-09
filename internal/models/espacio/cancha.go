package espacio

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type Cancha struct {
	Espacio
	Deporte tipos.Deporte `json:"deporte"`
}
