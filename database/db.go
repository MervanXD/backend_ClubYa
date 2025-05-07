package database

import (
	"database/sql"
	"log"

	"github.com/MervanXD/backend_ClubYa/config"
	"github.com/MervanXD/backend_ClubYa/logs"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	// sacamos el .env
	dsn := config.LoadConfig()

	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error abriendo la conexion: %v", err)
	}
	//hacemos ping para ver si la conexion esta bien
	if err := DB.Ping(); err != nil {
		logs.Logger.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}
}

func CloseDB() {
	// basicamente cerramos la base de datos
	if DB != nil {
		if err := DB.Close(); err != nil {
			logs.Logger.Print("error al cerrar la DB: ", err)
		}
	}

}

func IsHealthy() bool {
	// para saber si la conexion sigue activa
	return DB.Ping() == nil
}
