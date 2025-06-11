package main

import (
	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/internal/app"
	"github.com/MervanXD/backend_ClubYa/logs"
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

	// esto es para crear una app
	app := app.SetupApp()

	if err := app.Listen(":4000"); err != nil {
		logs.Logger.Fatal("Error al iniciar el servidor: ", err)
		panic(err)
	}

}
