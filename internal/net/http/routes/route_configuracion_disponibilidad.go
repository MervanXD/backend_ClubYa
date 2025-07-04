package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasConfiguracionDisponibilidad(api fiber.Router) {
	api.Get("/configuracion-disponibilidad/:id_espacio", handlers.ListarConfiguracionDisponibilidad)
	api.Post("/configuracion-disponibilidad", handlers.ActualizarConfiguracionDisponibilidad)
	api.Put("/configuracion-disponibilidad", handlers.ActualizarConfiguracionDisponibilidad)
	api.Get("/configuracion-disponibilidad/inscritos/:id_espacio", handlers.ObtenerInscritosEspacio)
}
