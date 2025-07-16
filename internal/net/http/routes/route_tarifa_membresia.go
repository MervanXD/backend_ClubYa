package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func TarifaMembresiaRoutes(app *fiber.App) {
	app.Get("/tarifa-membresia/:id", handlers.ObtenerTarifaMembresia)
}
