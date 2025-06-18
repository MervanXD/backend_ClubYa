package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasInscripcionEvento(api fiber.Router) {
	api.Post("/evento/inscripcion", handlers.RegistrarInscripcionEvento)
	api.Get("/inscripciones-eventos-socio/:id", handlers.ListarEventosSocioId)
	api.Post("/evento/anular-inscripcion", handlers.AnularInscripcionEvento)
}
