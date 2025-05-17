package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasPago(app *fiber.App) {
	app.Post("/pago/registrar", handlers.RegistrarPago)
}
