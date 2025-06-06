package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasEvento(api fiber.Router) {
	api.Post("/evento", handlers.CrearEvento)
	api.Get("/eventos", handlers.ListarEventos)
	api.Get("/eventos/:id", handlers.ObtenerEventoPorId)
	//app.Put("/eventos/:id", handlers.ActualizarEvento)
	//app.Delete("/eventos/:id", handlers.EliminarEvento)
}
