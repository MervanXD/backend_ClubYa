package sesiones

type SesionUpdate struct {
	IDSesion   *int    `json:"id_sesion"`   // ID de la sesión
	IdGrupo    *int    `json:"id_grupo"`    // ID del grupo al que pertenece la sesión
	Dia        *string `json:"dia"`         // Día de la sesión
	HoraInicio *string `json:"hora_inicio"` // Hora de inicio de la sesión
	HoraFin    *string `json:"hora_fin"`    // Hora de fin de la sesión
}
