package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasTarifas(api fiber.Router) {
	api.Get("/tarifas-familiar", handlers.ListarTarifasMiembrosFamiliar)
	api.Put("/tarifa-familiar-actualizar", handlers.ModificarTarifaMiembroAdicional)
	api.Post("/tarifa-familiar-nueva", handlers.InsertarTarifaMiembroAdicional)
}
