package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasFamiliar(app *fiber.App) {
	app.Post("/familiar", handlers.InsertarFamiliares)
	app.Get("/familiares/titular/:id", handlers.ObtenerFamiliaresPorTitular)
	//app.Get("/persona", handlers.ListarPersonas)
	//app.Get("/persona/:id", handlers.ObtenerPersonaPorID)
	//app.Put("/persona/:id", handlers.ActualizarPersona)
	//app.Delete("/persona/:id", handlers.EliminarPersona)
}