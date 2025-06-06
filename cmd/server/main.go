package main

import (
	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/internal/net/http/routes"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	logs.InitLogger()
	defer logs.CloseLogger()

	// inicializamos la base de datos
	database.InitDB()
	defer database.CloseDB()
	if prueba := database.IsHealthy(); !prueba {
		logs.Logger.Fatal("Error al conectar a la base de datos")
		panic("Error al conectar a la base de datos")
	}

	// esto es para crear una app y tener handlers y eso
	app := fiber.New()
	app.Use(cors.New())
	api:=app.Group("/api")
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

	if err := app.Listen(":4000"); err != nil {
		logs.Logger.Fatal("Error al iniciar el servidor: ", err)
		panic(err)
	}

}
