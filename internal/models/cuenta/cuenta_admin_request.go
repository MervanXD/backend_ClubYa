package cuenta

import "github.com/MervanXD/backend_ClubYa/internal/models/tipos"

type CuentaAdminRequest struct {
	IdCuenta        int                          `json:"id_cuenta"`
	IdPersona       int                          `json:"id_persona"`
	Username        string                       `json:"username"`
	Email           string                       `json:"email"`
	Rol             tipos.Rol                    `json:"rol"`
	Nombre          string                       `json:"nombre"`
	Apellidos       string                       `json:"apellidos"`
	Sexo            tipos.Sexo                   `json:"sexo"`
	TipoDocumento   tipos.TipoDocumentoIdentidad `json:"tipo_documento"`
	NroDocumento    string                       `json:"nro_documento"`
	FechaNacimiento string                       `json:"fecha_nacimiento"`
	Telefono        string                       `json:"telefono"`
	Pais            string                       `json:"pais"`
	Provincia       string                       `json:"provincia"`
	Distrito        string                       `json:"distrito"`
	TipoVia         tipos.TipoVia                `json:"tipo_via"`
	Direccion       string                       `json:"direccion"`
	Referencia      string                       `json:"referencia"`
	Ciudad          string                       `json:"ciudad"`
	CodigoPostal    string                       `json:"codigo_postal"`
}
