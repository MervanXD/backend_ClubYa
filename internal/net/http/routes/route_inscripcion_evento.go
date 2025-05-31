package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasInscripcionEvento(app *fiber.App) {
	app.Post("/evento/inscripcion", handlers.RegistrarInscripcionEvento)
}
