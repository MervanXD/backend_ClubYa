package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasSolicitudMembresia(app *fiber.App) {
	//app.Post("/solicitud-membresia", handlers.ListarSolicitudMembresia)
	app.Get("/solicitud-membresia", handlers.ListarSolicitudMembresia)
	app.Put("/solicitud-membresia/:id/estado_solicitud", handlers.ActualizarEstadoSolicitud)
	app.Get("/solicitud-membresia/:id/detalleSolicitud", handlers.DatosSolicitudId)
	app.Get("/solicitud-membresia/:id/detalleFamiliares", handlers.ListarFamiliaresSolicitudId)
	app.Get("/solicitud-membresia/:id/detallePersona", handlers.ObtenerPersonaPorSolicitudId)
	app.Get("/solicitud-membresia/:id/ObtenerEstado", handlers.ObtenerEstadoSolicitud)
	//app.Put("/solicitud-membresia/:id", handlers.)
	//app.Delete("/solicitud-membresia/:id", handlers.)
}
