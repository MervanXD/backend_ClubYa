package inscripcion_evento

type InscripcionEventoRepository interface {
	RegistrarInscripcion(fidPersona, idEvento, cantidadInvitados int) error
	ObtenerEventosSocio(idSocio int) ([]InscripcionSocioDTO, error)
	AnularInscripcion(idInscripcionEvento int, motivo string) error
	PagarInscripcion(fidPersona int, idEvento int, concepto string, metodoPago string, monto float64) (int, error)
	//RegistrarInscripcionesEventoTx(requests []InscripcionEvento) error
}
