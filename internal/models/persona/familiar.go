package persona

type Familiar struct {
	Persona
	MismaDireccionPostulante bool   `json:"misma_direccion_postulante"`
	EsConyuge                bool   `json:"es_conyuge"`
	TipoFamiliar             string `json:"tipo_familiar"`
}
