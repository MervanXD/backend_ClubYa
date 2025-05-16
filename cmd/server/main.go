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
	routes.RutasEspacioSocial(app)
	routes.TitularRoutes(app)
	routes.RutasSolicitudMembresia(app)
	routes.RutasMembresia(app)
	routes.RutasPago(app)
	routes.RutasCuota(app)
	routes.RutasInscripcionEvento(app)

	routes.RutasFamiliar(app)

	if err := app.Listen(":4000"); err != nil {
		logs.Logger.Fatal("Error al iniciar el servidor: ", err)
		panic(err)
	}

}
