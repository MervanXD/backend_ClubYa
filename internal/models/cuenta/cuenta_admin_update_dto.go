package cuenta

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/persona"
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type CuentaAdminUpdateDTO struct {
	Username   *string                   `json:"username"`
	Email      *string                   `json:"email"`
	Contrasena *string                   `json:"contrasena"`
	Rol        *tipos.Rol                `json:"rol"`
	Titular    *persona.TitularUpdateDTO `json:"datosPersonales"`
}
