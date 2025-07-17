package persona

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
)

type PersonaUpdateDTO struct {
	Id              *int                          `json:"id"`
	Nombre          *string                       `json:"nombre"`
	Apellidos       *string                       `json:"apellidos"`
	Sexo            *tipos.Sexo                   `json:"sexo"`
	TipoDocumento   *tipos.TipoDocumentoIdentidad `json:"tipo_documento"`
	NroDocumento    *string                       `json:"nro_documento"`
	FechaNacimiento *string                       `json:"fecha_nacimiento"`
	Telefono        *string                       `json:"telefono"`
	Pais            *string                       `json:"pais"`
	Provincia       *string                       `json:"provincia"`
	Distrito        *string                       `json:"distrito"`
	Direccion       *string                       `json:"direccion"`
	TipoVia         *tipos.TipoVia                `json:"tipo_via"`
	Referencia      *utils.NullString             `json:"referencia"`
	Ciudad          *utils.NullString             `json:"ciudad"`
	CodigoPostal    *utils.NullString             `json:"codigo_postal"`
}
