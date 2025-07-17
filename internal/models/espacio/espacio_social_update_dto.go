package espacio

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
)

type EspacioSocialUpdateDTO struct {
	Nombre         *string           `json:"nombre,omitempty"`
	Ubicacion      *tipos.Ubicacion  `json:"ubicacion,omitempty"`
	Capacidad      *int              `json:"capacidad,omitempty"`
	Costo          *float64          `json:"costo,omitempty"`
	Codigo         *string           `json:"codigo,omitempty"`
	Reglamento     *utils.NullString `json:"reglamento,omitempty"`
	Imagen         *utils.NullString `json:"imagen,omitempty"`
	DuracionBloque *int              `json:"duracion_bloque,omitempty"`
	EstadoEspacio  *int              `json:"estado_espacio,omitempty"`
	Actividad      *tipos.Actividad  `json:"actividad,omitempty"`
}
