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

func findProjectRoot(startDir, projectName string) string {
	dir := startDir
	for {
		if filepath.Base(dir) == projectName {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break // Llegó a la raíz del sistema
		}
		dir = parent
	}
	return ""
}

func isRunningInContainer() bool {
    // Detectar si estamos en un contenedor Docker
    if _, err := os.Stat("/.dockerenv"); err == nil {
        return true
    }
    // También verificar variables de entorno típicas de Docker
    if os.Getenv("DOCKER_CONTAINER") != "" || os.Getenv("ENVIRONMENT") == "production" {
        return true
    }
    return false
}

func LoadConfig() string {
    var projectRoot string
    
    if isRunningInContainer() {
        // Estamos en Docker/contenedor
        projectRoot = "/root"
        logs.Logger.Println("🐳 Ejecutándose en contenedor Docker")
        
        // Asegurar que existan las carpetas necesarias
        os.MkdirAll("/root/logs", 0755)
        os.MkdirAll("/root/config", 0755)
    } else {
        // Estamos en desarrollo local
        wd, err := os.Getwd()
        if err != nil {
            logs.Logger.Fatalf("Error al obtener el directorio de trabajo: %v", err)
        }
        
        projectRoot = findProjectRoot(wd, "backend_ClubYa")
        if projectRoot == "" {
            logs.Logger.Fatalf("No se encontró la raíz del proyecto backend_ClubYa partiendo de: %s", wd)
        }
        logs.Logger.Println("💻 Ejecutándose en desarrollo local, PROJECT_ROOT:", projectRoot)
    }

    // Intentar cargar archivos .env en orden de prioridad
    envFiles := []string{".env"}
    var err error
    var loadedFile string
    
    for _, envFile := range envFiles {
        envPath := filepath.Join(projectRoot, envFile)
        if _, fileErr := os.Stat(envPath); fileErr == nil {
            err = godotenv.Load(envPath)
            if err == nil {
                loadedFile = envPath
                logs.Logger.Println("📄 Archivo de configuración cargado:", loadedFile)
                break
            }
        }
    }
    
    if err != nil || loadedFile == "" {
        logs.Logger.Fatal("❌ No se pudo cargar ningún archivo de configuración (.env.local o .env)")
    }

    // Obtener variables encriptadas del entorno
    encryptedPort := os.Getenv("DB_PORT")
    encryptedPassword := os.Getenv("DB_PASS")
    encryptedHost := os.Getenv("DB_HOST")
    encryptedUser := os.Getenv("DB_USER")
    encryptedName := os.Getenv("DB_NAME")


    // Desencriptar las variables de la base de datos
    DBPassword, err = security.DecryptEnv(encryptedPassword)
    if err != nil {
        logs.Logger.Fatal("Error al desencriptar DB_PASS: ", err)
    }
    
    DBHost, err = security.DecryptEnv(encryptedHost)
    if err != nil {
        logs.Logger.Fatal("Error al desencriptar DB_HOST: ", err)
    }
    
    DBPort, err = security.DecryptEnv(encryptedPort)
    if err != nil {
        logs.Logger.Fatal("Error al desencriptar DB_PORT: ", err)
    }
    
    DBUser, err = security.DecryptEnv(encryptedUser)
    if err != nil {
        logs.Logger.Fatal("Error al desencriptar DB_USER: ", err)
    }
    
    DBName, err = security.DecryptEnv(encryptedName)
    if err != nil {
        logs.Logger.Fatal("Error al desencriptar DB_NAME: ", err)
    }

    // Construir DSN para la conexión a la base de datos
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
        DBUser, DBPassword, DBHost, DBPort, DBName, "parseTime=true")
    
    logs.Logger.Println("✅ Configuración de base de datos cargada exitosamente")
    return dsn
}
