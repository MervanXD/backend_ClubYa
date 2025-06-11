package reserva

type ReservaRepository interface {
	ReservarEspacio(idEspacio int, idHorarioDia int, idBloqueTiempo int) error
	ReservarEspacioSocial(reserva ReservaEspacio) error
	AnulacionReservaEspacioSocial(idReserva int, idEspacio int, idHorarioDia int, idBloque int, motivo string) error
	ObtenerReservasEspaciosSocialesSocio(idSocio int) ([]ReservaEspacioSocialRequest, error)
	ObtenerReservasCanchasSocio(idSocio int) ([]ReservaCanchaRequest, error)
}
