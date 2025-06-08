package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/MervanXD/backend_ClubYa/internal/pkgs/security"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/joho/godotenv"
)

var (
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
)

func LoadConfig() string {
	// Cargar el archivo .env
	rootDir := os.Getenv("BACKEND_CLUBYA_ROOT")
	if rootDir == "" {
		logs.Logger.Fatal("No se ha definido la variable de entorno BACKEND_CLUBYA_ROOT")
	}
	envDir := filepath.Join(rootDir, ".env")
	err := godotenv.Load(envDir)
	if err != nil {
		logs.Logger.Fatal("Error cargando .env")
	}

	encryptedPort := os.Getenv("DB_PORT")
	encryptedPassword := os.Getenv("DB_PASS")
	encryptedHost := os.Getenv("DB_HOST")
	encryptedUser := os.Getenv("DB_USER")
	encryptedName := os.Getenv("DB_NAME")

	DBPassword, err = security.DecryptEnv(encryptedPassword) // aquí desencriptas
	if err != nil {
		logs.Logger.Fatal("Error al desencriptar los datos de la base de datos: ", err)
	}
	DBHost, err = security.DecryptEnv(encryptedHost) // aquí desencriptas
	if err != nil {
		logs.Logger.Fatal("Error al desencriptar los datos de la base de datos: ", err)
	}
	DBPort, err = security.DecryptEnv(encryptedPort) // aquí desencriptas
	if err != nil {
		logs.Logger.Fatal("Error al desencriptar los datos de la base de datos: ", err)
	}
	DBUser, err = security.DecryptEnv(encryptedUser) // aquí desencriptas
	if err != nil {
		logs.Logger.Fatal("Error al desencriptar los datos de la base de datos: ", err)
	}
	DBName, err = security.DecryptEnv(encryptedName) // aquí desencriptas
	if err != nil {
		logs.Logger.Fatal("Error al desencriptar los datos de la base de datos: ", err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		DBUser, DBPassword, DBHost, DBPort, DBName, "parseTime=true")
	//fmt.Println("DSN: ", dsn)
	return dsn
}
