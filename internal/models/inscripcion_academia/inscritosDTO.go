package inscripcionacademia

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/sesiones"
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type InscritoAcademiaDTO struct {
	IdInscripcion     int               `json:"id_inscripcion"`
	IdPersona         int               `json:"id_persona"`
	NombrePersona     string            `json:"nombre_persona"`
	ApellidoPersona   string            `json:"apellido_persona"`
	TipoSocio         string            `json:"tipo_socio"`
	NombreGrupo       string            `json:"nombre_grupo"`
	IdGrupo           int               `json:"id_grupo"`
	EdadMinima        int               `json:"edadMinima"`
	EdadMaxima        int               `json:"edadMaxima"`
	NombreAcademia    string            `json:"nombre_academia"`
	FechaInicio       string            `json:"fecha_inicio"`
	FechaFin          string            `json:"fecha_fin"`
	EstadoInscripcion tipos.Estado      `json:"estado"`
	FechaInscripcion  string            `json:"fechaInscripcion"`
	Monto             float64           `json:"montoPagado"`
	Sesiones          []sesiones.Sesion `json:"sesiones"`
}
