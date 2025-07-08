package cuenta

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
)

type CuentaSocioDTO struct {
	IdCuenta             int                          `json:"id_cuenta"`
	IdPersona            int                          `json:"id_persona"`
	Username             string                       `json:"username"`
	Email                utils.NullString             `json:"email"`
	Rol                  tipos.Rol                    `json:"rol"`
	Nombre               string                       `json:"nombre"`
	Apellidos            string                       `json:"apellidos"`
	Sexo                 tipos.Sexo                   `json:"sexo"`
	TipoDocumento        tipos.TipoDocumentoIdentidad `json:"tipo_documento"`
	NroDocumento         string                       `json:"nro_documento"`
	FechaNacimiento      string                       `json:"fecha_nacimiento"`
	Telefono             string                       `json:"telefono"`
	Pais                 utils.NullString             `json:"pais"`
	Provincia            utils.NullString             `json:"provincia"`
	Distrito             utils.NullString             `json:"distrito"`
	TipoVia              tipos.TipoVia                `json:"tipo_via"`
	Direccion            utils.NullString             `json:"direccion"`
	Referencia           utils.NullString             `json:"referencia"`
	Ciudad               utils.NullString             `json:"ciudad"`
	CodigoPostal         utils.NullString             `json:"codigo_postal"`
	EstadoMembresia      tipos.EstadoMembresia        `json:"estado_membresia"`
	FechaInicioMembresia string                       `json:"fecha_inicio_membresia"`
}
