package reserva

import "context"

type ReservaRepository interface {
	ReservarEspacio(ctx context.Context, reservaEspacio ReservaEspacio) error
	AnulacionReservaEspacioSocial(idReserva int, idEspacio int, idHorarioDia int, idBloque int, motivo string) error
	ObtenerReservasEspaciosSocialesSocio(idSocio int) ([]ReservaEspacioSocialRequest, error)
	ObtenerReservasCanchasSocio(idSocio int) ([]ReservaCanchaRequest, error)
	ObtenerReservasPorEspacio(idEspacio int) ([]ReservaRequest, error)
	AceptarDevolucionAnulacionReserva(AnulacionReservaRequest) error
}
