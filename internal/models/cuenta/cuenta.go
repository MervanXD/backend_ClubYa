package cuenta

import "github.com/MervanXD/backend_ClubYa/internal/models/tipos"

type Cuenta struct {
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Contrasena string    `json:"contrasena"`
	Email      string    `json:"email"`
	Activo     bool      `json:"activo"`
	Rol        tipos.Rol `json:"rol"`
}
