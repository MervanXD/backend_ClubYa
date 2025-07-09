package cuenta

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/persona"
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type CuentaSocioDTO struct {
	IdCuenta             int                   `json:"id_cuenta"`
	Username             string                `json:"username"`
	Email                string                `json:"email"`
	Rol                  tipos.Rol             `json:"rol"`
	Persona              persona.Titular       `json:"datosPersonales"`
	EstadoMembresia      tipos.EstadoMembresia `json:"estado_membresia"`
	FechaInicioMembresia string                `json:"fecha_inicio_membresia"`
}
