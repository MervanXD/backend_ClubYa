package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasPrueba(api fiber.Router) {
	api.Get("/pruebas", handlers.ListarPruebas)
	api.Post("/prueba", handlers.CrearPrueba)
	api.Put("/prueba", handlers.ModificarPrueba)
	api.Delete("/prueba/:id", handlers.EliminarPrueba)
}
