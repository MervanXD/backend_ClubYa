package cuenta

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
)

type CuentaUsuariosRequest struct {
	IdCuenta     int              `json:"idCuenta"`
	Username     string           `json:"username"`
	Email        utils.NullString `json:"email"`
	Nombre       string           `json:"nombre"`
	NroDocumento string           `json:"nro_documento"`
	Telefono     string           `json:"telefono"`
	Rol          tipos.Rol        `json:"rol"`
}
