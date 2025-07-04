package app

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/routes"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupApp() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	api := app.Group("/api")

	// Puedes envolver cada registro de ruta en un recover para loggear errores inesperados
	defer func() {
		if r := recover(); r != nil {
			logs.Logger.Printf("Pánico al registrar rutas: %v", r)
		}
	}()
	routes.RutasEspacioSocial(api)
	routes.TitularRoutes(api)
	routes.RutasSolicitudMembresia(api)
	routes.RutasMembresia(api)
	routes.RutasPago(api)
	routes.RutasCuota(api)
	routes.RutasInscripcionEvento(api)
	routes.RutasEvento(api)
	routes.RouteCuenta(api)
	routes.RutasReservaEspacio(api)
	routes.RutasDisponibilidad(api)
	routes.RutasFamiliar(api)
	routes.RutasCancha(api)
	routes.RutasInscripcionAcademia(api)
	routes.RutasAcademia(api)
	routes.RutasReporte(api)

	routes.RutasPrueba(api)
	return app
}
