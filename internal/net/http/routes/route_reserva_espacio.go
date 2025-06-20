package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasReservaEspacio(api fiber.Router) {
	api.Put("/reserva-espacio", handlers.ReservarEspacio)
	api.Post("/reserva-espacio-social", handlers.ReservarEspacioSocial)
	api.Put("/anulacion-reserva-espacio-social", handlers.AnularReservarEspacioSocial)
	api.Get("/reservas-espacios-social-socio/:idSocio", handlers.ListarEspaciosSocialesSocio)
	api.Get("/reserva-canchas/:idSocio",handlers.ListarCanchasSocio)
	api.Get("/reservas-espacios/:idEspacio", handlers.ListarReservasPorEspacio)
	api.Put("/anulacion-reserva/aceptar", handlers.AceptarDevolucionAnulacionReserva)
}
