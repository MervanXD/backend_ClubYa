package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasEvento(app *fiber.App) {
	//app.Post("/evento", handlers.CrearEvento)
	app.Get("/eventos", handlers.ListarEventos)
	app.Get("/eventos/:id", handlers.ObtenerEventoPorId)
	//app.Put("/eventos/:id", handlers.ActualizarEvento)
	//app.Delete("/eventos/:id", handlers.EliminarEvento)
}
