package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func TitularRoutes(app *fiber.App) {
	//app.Post("/titular/postulacion", handlers.RegistrarPostulante)
	app.Post("/titular/crear/:idCuenta", handlers.CrearTitular)
	//app.Get("/persona", handlers.ListarPersonas)
	//app.Get("/persona/:id", handlers.ObtenerPersonaPorID)
	//app.Put("/persona/:id", handlers.ActualizarPersona)
	//app.Delete("/persona/:id", handlers.EliminarPersona)
}
