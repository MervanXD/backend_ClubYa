package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RouteCuenta(api fiber.Router) {
	api.Post("/cuenta/crear", handlers.CrearCuenta)
	api.Post("/cuenta/login", handlers.LogIn)
}
