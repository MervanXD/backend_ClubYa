package cuenta

import "github.com/MervanXD/backend_ClubYa/internal/models/tipos"

type CuentaSocioRequest struct {
	IDCuenta             int64                 `json:"idCuenta"`
	Nombre               string                `json:"nombre"`
	Apellidos            string                `json:"apellidos"`
	Email                string                `json:"email"`
	FechaInicioMembresia string                `json:"fecha_inicio_membresia"`
	Username             string                `json:"username"`
	EstadoMembresia      tipos.EstadoMembresia `json:"estado_membresia"`
	Activo               bool                  `json:"activo"`
}
