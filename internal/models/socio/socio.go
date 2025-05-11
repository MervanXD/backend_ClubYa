package socio

type Socio struct {
	IdPersona       int    `json:"id_persona"`
	NumeroMembresia string `json:"numero_membresia"`
	EsTitular       bool   `json:"es_titular"`
	EsVitalicio     bool   `json:"es_vitalicio"`
}
