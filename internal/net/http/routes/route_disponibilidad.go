package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasDisponibilidad(api fiber.Router) {
	api.Get("/disponibilidad/:id_espacio/:id_horario_dia/:id_bloque_tiempo", handlers.ObtenerDisponibilidadEspacioSocialPorId)
	api.Put("/disponibilidad/actualizar", handlers.ActualizarDetalleDisponibilidad)
	api.Get("/disponibilidad-espacio/:idEspacio/:fecha", handlers.ListarDisponibilidadEspacio)
}
