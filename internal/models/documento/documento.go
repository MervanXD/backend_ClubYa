package documento

import "github.com/MervanXD/backend_ClubYa/internal/models/tipos"

type Documento struct {
	Id            int                 `json:"id"`
	contenido     []byte              `json:"contenido"`
	TipoDocumento tipos.TipoDocumento `json:"tipo_documento"`
}
