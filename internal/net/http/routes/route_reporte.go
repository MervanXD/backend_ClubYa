package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasReporte(api fiber.Router) {
	api.Get("/reporte/membresias", handlers.GeneraReporteMembresias)
	api.Post("/reporte/academias", handlers.GeneraReporteAcademias)
}
