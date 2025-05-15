package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasMembresia(app *fiber.App) {
	app.Get("/membresia/solicitud/:id", handlers.ObtenerMembresiaPorSolicitud)
}