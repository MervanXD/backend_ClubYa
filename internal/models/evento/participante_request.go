package evento

type ParticipanteRequest struct {
	IdPersona         int    `json:"idPersona"`
	Nombres           string `json:"nombres"`
	Apellidos         string `json:"apellidos"`
	Dni               string `json:"dni"`
	FechaInscripcion  string `json:"fechaInscripcion"`
	HoraInscripcion   int    `json:"horaInscripcion"`
	EstadoInscripcion string `json:"estadoInscripcion"`
	CantidadInvitados int    `json:"cantidadInvitados"`
}
