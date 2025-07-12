package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func TitularRoutes(api fiber.Router) {
	//app.Post("/titular/postulacion", handlers.RegistrarPostulante)
	api.Post("/titular/crear/:idCuenta", handlers.CrearTitular)
	api.Get("/titular/:idTitular", handlers.ObtenerTitularPorID)
	api.Get("/buscar-titulares", handlers.BuscarTitulares)
	//app.Get("/persona", handlers.ListarPersonas)
	//app.Get("/persona/:id", handlers.ObtenerPersonaPorID)
	//app.Put("/persona/:id", handlers.ActualizarPersona)
	//app.Delete("/persona/:id", handlers.EliminarPersona)
}
