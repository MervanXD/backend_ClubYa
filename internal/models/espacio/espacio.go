package espacio

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type Espacio struct {
	Id            int             `json:"id"`
	Nombre        string          `json:"nombre"`
	Ubicacion     tipos.Ubicacion `json:"ubicacion"`
	Capacidad     int             `json:"capacidad"`
	Restricciones []string        `json:"restricciones"`
	Costo         float64         `json:"costo"`
}
