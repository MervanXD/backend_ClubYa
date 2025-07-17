package tarifas

import "github.com/MervanXD/backend_ClubYa/internal/models/tipos"

type TarifaMiembroAdicional struct {
	IdTarifaMiembroAdicional int                   `json:"id_tarifa_miembro_adicional"`
	FidTarifaMembresia       int                   `json:"id_tarifa_membresia"`
	TipoMiembro              tipos.MiembroFamiliar `json:"tipo_miembro"`
	CostoAdicional           float64               `json:"costo_adicional"`
	Estado                   bool                  `json:"estado"`
}
