package reserva

import "context"

type ReservaRepository interface {
	ReservarEspacio(ctx context.Context, reservaEspacio ReservaEspacio) error
	AnulacionReservaEspacio(idReserva int, idEspacio int, idHorarioDia int, motivo string, correo string) error
	ObtenerReservasEspaciosSocialesSocio(idSocio int) ([]ReservaEspacioSocialRequest, error)
	ObtenerReservasCanchasSocio(idSocio int) ([]ReservaCanchaRequest, error)
	ObtenerReservasPorEspacio(idEspacio int) ([]ReservaRequest, error)
	AceptarDevolucionAnulacionReserva(AnulacionReservaRequest) error
}
