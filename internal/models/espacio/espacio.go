package espacio

import (


)

type Espacio struct {
	Id          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Ubicacion   string `json:"ubicacion"`
	Capacidad   int    `json:"capacidad"`
	Costo       float64 `json:"costo"`
}
