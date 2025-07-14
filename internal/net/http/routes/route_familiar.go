package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasFamiliar(api fiber.Router) {
	api.Post("/familiar", handlers.InsertarFamiliares)
	api.Get("/familiar/titular/:id", handlers.ObtenerFamiliaresPorTitular)
	api.Get("/familiar/:id", handlers.ObtenerFamiliarPorIdPersona)
	api.Post("/familiar/registrar-con-solicitud/:id", handlers.RegistrarFamiliarConSolicitud)
	api.Post("/familiar/solicitud-retiro/:id", handlers.CrearSolicitudRetiro)
	// app.Get("/persona/:id", handlers.ObtenerPersonaPorID)
	// app.Put("/persona/:id", handlers.ActualizarPersona)
	// app.Delete("/persona/:id", handlers.EliminarPersona)
}
