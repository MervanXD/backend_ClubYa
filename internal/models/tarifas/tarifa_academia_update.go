package tarifas

import "github.com/MervanXD/backend_ClubYa/internal/models/tipos"

type TarifaAcademiaUpdate struct {
	ID                  *int                       `json:"id"`                    // ID de la tarifa
	Periodo             *int                       `json:"periodo"`               // Periodo de la tarifa
	UnidadFrecuenciaSem *tipos.UnidadFrecuenciaSem `json:"unidad_frecuencia_sem"` // Unidad de frecuencia semanal
	CantidadFrecuencia  *int                       `json:"cantidad_frecuencia"`   // Cantidad de frecuencia
	TipoSocio           *tipos.TipoSocio           `json:"tipo_socio"`            // Tipo de socio
	Monto               *float64                   `json:"monto"`                 // Monto de la tarifa
	EsActiva            *bool                      `json:"es_activa"`             // Ind
	IdGrupo             *int                       `json:"id_grupo"`              // ID del grupo al que pertenece la tarifa
}
