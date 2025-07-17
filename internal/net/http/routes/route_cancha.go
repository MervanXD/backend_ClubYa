package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasCancha(api fiber.Router) {

	api.Get("/canchas-horarios", handlers.ListarCanchasHorarios)
	api.Post("/canchas", handlers.InsertarCancha)
	api.Patch("/canchas/:id", handlers.ActualizarCancha)
	api.Get("/canchas-configuracion", handlers.ListarCanchasConfiguracion)
	api.Get("/canchas-disponibilidad-configuracion", handlers.ListarDisponibilidadEspacio)
}

