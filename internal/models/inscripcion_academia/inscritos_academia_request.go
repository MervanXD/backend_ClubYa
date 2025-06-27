package inscripcionacademia

import "github.com/MervanXD/backend_ClubYa/internal/models/tipos"

type InscritosAcademiaRequest struct {
	IdPersona             int                          `json:"id_persona"`
	NombrePersona         string                       `json:"nombre_persona"`
	ApellidoPersona       string                       `json:"apellido_persona"`
	TipoSocio             string                       `json:"tipo_socio"`
	NombreGrupo           string                       `json:"nombre_grupo"`
	IdGrupo               int                          `json:"id_grupo"`
	EdadPersona           int                          `json:"edad_persona"`
	FechaInscripcion      string                       `json:"fecha_inscripcion"`
	Monto                 float64                      `json:"monto_pagado"`
	EstadoInscripcion     tipos.Estado                 `json:"estado_inscripcion"`
	AnulaccionInscripcion AnulacionInscripcionAcademia `json:"anulacion_inscripcion"`
}
