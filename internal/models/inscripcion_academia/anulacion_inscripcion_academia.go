package inscripcionacademia

import "github.com/MervanXD/backend_ClubYa/internal/models/tipos"

type AnulacionInscripcionAcademia struct {
	IdAnulacion   int        `json:"id_anulacion"`
	IdInscripcion int        `json:"id_inscripcion"`
	Fecha         string     `json:"fecha"`
	Motivo        string     `json:"motivo"`
	Tipo          tipos.Tipo `json:"tipo_reserva"`
	Devolucion    float64    `json:"devolucion"`
}
