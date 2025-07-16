package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasTarifasMembresia(api fiber.Router) {
	api.Get("/tarifa-membresia/:id", handlers.ObtenerTarifaMembresia)
}
