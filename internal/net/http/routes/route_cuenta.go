package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RouteCuenta(app *fiber.App) {
	app.Post("/cuenta/crear", handlers.CrearCuenta)
	app.Post("/cuenta/login", handlers.LogIn)
}
