package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasEspacioSocial(api fiber.Router) {
	api.Post("/espacio-social", handlers.CrearEspacioSocial)
	api.Get("/espacio-social", handlers.ListarEspaciosSociales)
	api.Get("/espacio-social/:id", handlers.ObtenerEspacioSocialPorId)
	api.Get("/espacio-social-horarios", handlers.ListarEspaciosSocialesHorarios)
	api.Put("/espacio-social/:id", handlers.ActualizarEspacioSocial)
	api.Get("/espacio-social-configuracion", handlers.ListarEspaciosSocialesConfiguracion)
	api.Post("/espacio-social-disponibilidad-configuracion", handlers.ListarDisponibilidadEspacio)
	//app.Delete("/espacio-social/:id", handlers.EliminarEspacioSocial)
}
