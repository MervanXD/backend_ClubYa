package espacio

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type Espacio struct {
	Id            int             `json:"id"`
	Nombre        string          `json:"nombre"`
	Ubicacion     tipos.Ubicacion `json:"ubicacion"`
	Capacidad     int             `json:"capacidad"`
	Codigo        string          `json:"codigo"`
	Costo         float64         `json:"costo"`
	Reglamento    []byte          `json:"reglamento"`
	Imagen        []byte          `json:"imagen"`
	EstadoEspacio int             `json:"estado_espacio"`
}
