package main

import (
	"fmt"

	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/internal/app"
	"github.com/MervanXD/backend_ClubYa/logs"
)

func main() {
	docker := logs.InitLogger()
	if docker == 1 {
		defer logs.CloseLogger()
	}

	logs.Logger.Println("🚀 Aplicación iniciada")

	// inicializamos la base de datos
	logs.Logger.Println("📊 Inicializando base de datos...")
	database.InitDB()
	defer database.CloseDB()
	fmt.Println("✅ Base de datos inicializada")

	logs.Logger.Println("🔍 Verificando salud de la base de datos...")
	if prueba := database.IsHealthy(); !prueba {
		logs.Logger.Fatal("❌ Error al conectar a la base de datos")
		panic("Error al conectar a la base de datos")
	}
	logs.Logger.Println("✅ Base de datos conectada correctamente")

	// esto es para crear una app
	logs.Logger.Println("🌐 Configurando aplicación Fiber...")
	app := app.SetupApp()
	logs.Logger.Println("✅ Aplicación Fiber configurada")

	logs.Logger.Println("🌐 Servidor iniciando en puerto 4000...")
	if err := app.Listen(":4000"); err != nil {
		logs.Logger.Fatal("❌ Error al iniciar el servidor: ", err)
		panic(err)
	}
}
