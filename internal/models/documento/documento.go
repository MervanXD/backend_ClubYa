package documento

import "github.com/MervanXD/backend_ClubYa/internal/models/tipos"

type Documento struct {
	Id            int                 `json:"id"`
	Contenido     []byte              `json:"contenido"`
	TipoDocumento tipos.TipoDocumento `json:"tipo_documento"`
	IdPersona     int                 `json:"id_persona"`
}
