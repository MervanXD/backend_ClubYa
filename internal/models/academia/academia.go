package academia

import (
	grupoacademia "github.com/MervanXD/backend_ClubYa/internal/models/grupo_academia"
	"github.com/MervanXD/backend_ClubYa/internal/models/tarifas"
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type Academia struct {
	ID             int                           `json:"id"`
	Nombre         string                        `json:"nombre"`
	Descripcion    string                        `json:"descripcion"`
	Deporte        tipos.Deporte                 `json:"deporte"`
	Entrenador     string                        `json:"entrenador"`
	CostoUniforme  float64                       `json:"costo_uniforme"`
	CostoMatricula float64                       `json:"costo_matricula"`
	Reglamento     []byte                        `json:"reglamento"`
	Imagen         []byte                        `json:"imagen"`
	Indicaciones   string                        `json:"indicaciones"`
	Grupos         []grupoacademia.GrupoAcademia `json:"grupos"`
	Tarifas        []tarifas.TarifaAcademia      `json:"tarifas"`
}
