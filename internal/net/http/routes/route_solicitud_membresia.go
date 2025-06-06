package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasSolicitudMembresia(api fiber.Router) {
	//app.Post("/solicitud-membresia", handlers.ListarSolicitudMembresia)
	api.Get("/solicitud-membresia", handlers.ListarSolicitudMembresia)
	api.Put("/solicitud-membresia/:id/estado_solicitud", handlers.ActualizarEstadoSolicitud)
	api.Get("/solicitud-membresia/:id/detalleSolicitud", handlers.DatosSolicitudId)
	api.Get("/solicitud-membresia/:id/detalleFamiliares", handlers.ListarFamiliaresSolicitudId)
	api.Get("/solicitud-membresia/:id/detallePersona", handlers.ObtenerPersonaPorSolicitudId)
	api.Get("/solicitud-membresia/:id/ObtenerEstado", handlers.ObtenerEstadoSolicitud)
	//app.Put("/solicitud-membresia/:id", handlers.)
	//app.Delete("/solicitud-membresia/:id", handlers.)
}
