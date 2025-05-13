package persona

type Socio struct {
	IdMembresia int  `json:"id_membresia"`
	IdPersona   int  `json:"id_persona"`
	EsTitular   bool `json:"es_titular"`
	EsVitalicio bool `json:"es_vitalicio"`
}
