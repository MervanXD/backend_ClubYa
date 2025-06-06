package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasPago(api fiber.Router) {
	api.Post("/pago/registrar", handlers.RegistrarPago)
}
