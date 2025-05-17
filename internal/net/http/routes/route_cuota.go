package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasCuota(app *fiber.App) {
	app.Get("/cuotas/membresia/:id/:limite", handlers.ObtenerCuotasPorMembresia)

}