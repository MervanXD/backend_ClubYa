package espacio

import "github.com/MervanXD/backend_ClubYa/internal/models/tipos"

type EspacioSocialUpdateDTO struct {
	Nombre         *string          `json:"nombre,omitempty"`
	Ubicacion      *tipos.Ubicacion `json:"ubicacion,omitempty"`
	Capacidad      *int             `json:"capacidad,omitempty"`
	Costo          *float64         `json:"costo,omitempty"`
	Codigo         *string          `json:"codigo,omitempty"`
	Reglamento     *[]byte          `json:"reglamento,omitempty"`
	Imagen         *[]byte          `json:"imagen,omitempty"`
	DuracionBloque *int             `json:"duracion_bloque,omitempty"` // Duración del bloqueo en minutos
	EstadoEspacio  *int             `json:"estado_espacio,omitempty"`  // Estado del espacio (0: disponible, 1: reservado, 2: no disponible)
	Actividad      *tipos.Actividad `json:"actividad,omitempty"`
}
