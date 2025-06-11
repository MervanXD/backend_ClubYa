package inscripcion_evento

type InscripcionEventoRepository interface {
	RegistrarInscripcion(fidPersona, idEvento, cantidadInvitados int) error
	ObtenerEventosSocio(idSocio int) ([]InscripcionSocioDTO, error)
}
