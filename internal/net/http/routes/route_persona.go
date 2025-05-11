package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func PersonaRoutes(app *fiber.App) {
	app.Post("/persona", handlers.CrearPersona)
	//app.Get("/persona", handlers.ListarPersonas)
	//app.Get("/persona/:id", handlers.ObtenerPersonaPorID)
	//app.Put("/persona/:id", handlers.ActualizarPersona)
	//app.Delete("/persona/:id", handlers.EliminarPersona)
}
