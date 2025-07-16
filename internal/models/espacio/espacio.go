package espacio

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
)

type Espacio struct {
	Id             int              `json:"id"`
	Nombre         string           `json:"nombre"`
	Ubicacion      tipos.Ubicacion  `json:"ubicacion"`
	Capacidad      int              `json:"capacidad"`
	Codigo         string           `json:"codigo"`
	Costo          float64          `json:"costo"`
	Reglamento     utils.NullString `json:"reglamento,omitempty"`
	Imagen         utils.NullString `json:"imagen,omitempty"`
	EstadoEspacio  int              `json:"estado_espacio"`
	DuracionBloque int              `json:"duracion_bloque"`
}
