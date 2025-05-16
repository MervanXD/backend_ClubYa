package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasEvento(app *fiber.App) {
	//app.Post("/evento", handlers.CrearEvento)
	app.Get("/evento", handlers.ListarEventos)
	//app.Get("/evento/:id", handlers.ObtenerEventoPorId)
	//app.Put("/evento/:id", handlers.ActualizarEvento)
	//app.Delete("/evento/:id", handlers.EliminarEvento)
}
