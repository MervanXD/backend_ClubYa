package persona

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
)

type Titular struct {
	Persona
	IngresoPromedio  float64           `json:"ingreso_promedio"`
	Ocupacion        utils.NullString  `json:"ocupacion"`
	NombreEmpresa    utils.NullString  `json:"nombre_empresa"`
	DireccionEmpresa utils.NullString  `json:"direccion_empresa"`
	EsPostulante     bool              `json:"es_postulante"`
	TipoTrabajo      tipos.TipoTrabajo `json:"tipo_trabajo"`
}
