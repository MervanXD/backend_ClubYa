package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func SolicitudMembresiaRoutes(app *fiber.App) {
	app.Put("/solicitud-membresia/:id/estado_solicitud", handlers.ActualizarEstadoSolicitud)
}