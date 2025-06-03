package academia

import "github.com/MervanXD/backend_ClubYa/internal/models/tipos"

type Academia struct {
	ID             int64            `json:"id"`
	Nombre         string           `json:"nombre"`
	Descripcion    string           `json:"descripcion"`
	Deporte        tipos.Deporte    `json:"deporte"`
	Grupos         []GrupoAcademia  `json:"grupos"`
	PrecioSocio    float64          `json:"precio_socio"`
	PrecioExterno  float64          `json:"precio_externo"`
	Entrenador     string           `json:"entrenador"`
	CostoUniforme  float64          `json:"costo_uniforme"`
	CostoMatricula float64          `json:"costo_matricula"`
	Reglamento     []byte           `json:"reglamento"`
	Imagen         []byte           `json:"imagen"`
	Indicaciones   string           `json:"indicaciones"`
	Tarifas        []TarifaAcademia `json:"tarifas"`
}
