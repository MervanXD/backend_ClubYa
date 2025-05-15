package persona

import "github.com/MervanXD/backend_ClubYa/internal/models/tipos"

type Familiar struct {
	Persona
	MismaDireccionPostulante bool               `json:"misma_direccion_postulante"`
	EsConyuge                bool               `json:"es_conyuge"`
	TipoFamiliar             tipos.TipoFamiliar `json:"tipo_familiar"`
}
