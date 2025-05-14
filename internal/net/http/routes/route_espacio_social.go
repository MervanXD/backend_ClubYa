package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasEspacioSocial(app *fiber.App) {
	app.Post("/espacio-social", handlers.CrearEspacioSocial)
	app.Get("/espacio-social", handlers.ListarEspaciosSociales)
	app.Get("/espacio-social/:id", handlers.ObtenerEspacioSocialPorId)
	//app.Put("/espacio-social/:id", handlers.ActualizarEspacioSocial)
	//app.Delete("/espacio-social/:id", handlers.EliminarEspacioSocial)
}
