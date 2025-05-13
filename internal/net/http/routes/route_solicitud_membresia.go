package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasSolicitudMembresia(app *fiber.App) {
	//app.Post("/solicitud-membresia", handlers.ListarSolicitudMembresia)
	app.Get("/solicitud-membresia", handlers.ListarSolicitudMembresia)
	//app.Get("/solicitud-membresia/:id", handlers.)
	//app.Put("/solicitud-membresia/:id", handlers.)
	//app.Delete("/solicitud-membresia/:id", handlers.)
}
