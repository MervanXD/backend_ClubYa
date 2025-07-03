package membresia

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
)

type Membresia struct {
	Id          int                   `json:"id"`
	FechaInicio string                `json:"fecha_inicio"`
	FechaFin    utils.NullString      `json:"fecha_fin"`
	Tipo        tipos.TipoMembresia   `json:"tipo_membresia"`
	Estado      tipos.EstadoMembresia `json:"estado_membresia"`
}

type MembresiaDTO struct {
	Membresia     Membresia `json:"membresia"`
	NombreTitular string    `json:"nombre_titular"`
	CuotaBase     float64   `json:"cuota_base"`
}
