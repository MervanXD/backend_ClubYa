package persona

import "github.com/MervanXD/backend_ClubYa/internal/models/tipos"

type Titular struct {
	Persona
	IngresoPromedio  float64           `json:"ingreso_promedio"`
	Ocupacion        string            `json:"ocupacion"`
	NombreEmpresa    string            `json:"nombre_empresa"`
	DireccionEmpresa string            `json:"direccion_empresa"`
	EsPostulante     bool              `json:"es_postulante"`
	TipoTrabajo      tipos.TipoTrabajo `json:"tipo_trabajo"`
}
