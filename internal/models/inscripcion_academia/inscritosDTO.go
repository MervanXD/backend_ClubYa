package inscripcionacademia

import "github.com/MervanXD/backend_ClubYa/internal/models/sesiones"

type InscritoAcademiaDTO struct {
	IdPersona       int               `json:"id_persona"`
	NombrePersona   string            `json:"nombre_persona"`
	ApellidoPersona string            `json:"apellido_persona"`
	NombreAcademia  string            `json:"nombre_grupo"`
	IdGrupo         int               `json:"id_grupo"`
	Sesiones        []sesiones.Sesion `json:"sesiones"`
}
