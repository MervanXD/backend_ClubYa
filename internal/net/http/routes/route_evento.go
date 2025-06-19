package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasEvento(api fiber.Router) {
	api.Post("/evento", handlers.CrearEvento)
	api.Get("/eventos", handlers.ListarEventos)
	api.Get("/eventos/:id", handlers.ObtenerEventoPorId)
	api.Put("/eventos/:id", handlers.ModificarEvento)
	api.Put("/evento/cancelar/:id", handlers.CancelarEvento)
	api.Put("/evento/eliminar/:id", handlers.EliminarEvento)
	api.Get("/eventos/participantes/:id/", handlers.ListarParticipantesEvento)
}
