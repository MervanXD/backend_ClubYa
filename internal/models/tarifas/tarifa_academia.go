package tarifas

import "github.com/MervanXD/backend_ClubYa/internal/models/tipos"

type TarifaAcademia struct {
	ID                  int64                     `json:"id"`
	Periodo             int64                     `json:"periodo"`
	UnidadFrecuenciaSem tipos.UnidadFrecuenciaSem `json:"unidad_frecuencia_sem"`
	CantidadFrecuencia  int64                     `json:"cantidad_frecuencia"`
	TipoSocio           tipos.TipoSocio           `json:"tipo_socio"`
	Monto               float64                   `json:"monto"`
	EsActiva            bool                      `json:"es_activa"`
}
