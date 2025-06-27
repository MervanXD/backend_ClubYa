package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasAcademia(api fiber.Router) {
	api.Post("/academias", handlers.InsertarAcademia)
	api.Get("/academias", handlers.ListarAcademiasGenerales)
	api.Get("/academias/:id", handlers.ObtenerAcademiaPorId)
	api.Put("/academias/:id", handlers.ActualizarAcademia)
}
