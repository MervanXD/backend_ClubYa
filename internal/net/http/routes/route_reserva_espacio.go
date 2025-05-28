package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasReservaEspacio(app *fiber.App) {
	app.Put("/reserva-espacio", handlers.ReservarEspacio)
	app.Post("/reserva-espacio-social", handlers.ReservarEspacioSocial)
}
