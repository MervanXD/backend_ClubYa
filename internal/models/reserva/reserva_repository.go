package reserva

import "github.com/MervanXD/backend_ClubYa/internal/models/tipos"

type ReservaRepository interface {
	ReservarEspacio(idEspacio int, idHorarioDia int, idBloqueTiempo int) error
	ReservarEspacioSocial(reserva ReservaEspacio, dia tipos.Dia) error
	AnulacionReservaEspacioSocial(idReserva int, idEspacio int, idHorarioDia int, idBloque int, motivo string) error
	ObtenerReservasEspaciosSocialesSocio(idSocio int) ([]ReservaEspacioSocialRequest, error)
	ObtenerReservasCanchasSocio(idSocio int) ([]ReservaCanchaRequest, error)
	ObtenerReservasPorEspacio(idEspacio int) ([]ReservaRequest, error)
	AceptarDevolucionAnulacionReserva(AnulacionReservaRequest) error
}
