package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasCancha(app *fiber.App) {

	app.Get("/canchas-horarios", handlers.ListarCanchasHorarios)
	
}