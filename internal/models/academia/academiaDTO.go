package academia

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
)

type AcademiaDTO struct {
	ID          int64            `json:"id"`
	Nombre      string           `json:"nombre"`
	Descripcion string           `json:"descripcion"`
	FechaInicio string           `json:"fecha_inicio"`
	FechaFin    string           `json:"fecha_fin"`
	Deporte     tipos.Deporte    `json:"deporte"`
	Imagen      utils.NullString `json:"imagen"`
	Monto       float64          `json:"monto"`
	EdadMinima  int64            `json:"edad_minima"`
	Inscritos   int64            `json:"cantidad_inscritos"`
}
