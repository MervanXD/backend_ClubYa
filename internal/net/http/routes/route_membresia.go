package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasMembresia(api fiber.Router) {
	api.Get("/membresia/solicitud/:id", handlers.ObtenerMembresiaPorSolicitud)
}