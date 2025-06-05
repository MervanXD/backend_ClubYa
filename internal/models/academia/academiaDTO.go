package academia

import "github.com/MervanXD/backend_ClubYa/internal/models/tipos"

type AcademiaDTO struct {
	ID          int64         `json:"id"`
	Nombre      string        `json:"nombre"`
	Descripcion string        `json:"descripcion"`
	Deporte     tipos.Deporte `json:"deporte"`
	Imagen      []byte        `json:"imagen"`
	Monto       float64       `json:"monto"`
	EdadMinima  int64         `json:"edad_minima"`
	Inscritos   int64         `json:"cantidad_inscritos"`
}
