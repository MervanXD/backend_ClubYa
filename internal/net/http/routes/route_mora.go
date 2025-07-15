package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasMoras(api fiber.Router) {
	api.Get("/tarifas-mora", handlers.ListarTarifasMoras)
	api.Put("/tarifa-mora-actualizar", handlers.ModificarMora)
	api.Post("/tarifa-mora-nueva", handlers.InsertarMora)
}
