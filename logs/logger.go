package logs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
	_ "time"
)

var (
	LogFile *os.File
	Logger  *log.Logger
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


func InitLogger() int {
	var logDir string

	if isRunningInContainer() {
		// Estamos en Docker/contenedor
		logDir = "/root/logs"
		fmt.Println("🐳 Logger ejecutándose en contenedor Docker")

		// Usar solo stdout para evitar problemas de archivos
		Logger = log.New(os.Stdout, "APP_LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
		fmt.Println("✅ Logger inicializado (stdout)")

		// Test inmediato
		Logger.Println("🧪 Test: Logger funcionando")
		return 0
	} else {
		// Estamos en desarrollo local
		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Error al obtener el directorio de trabajo: %v", err)
		}

		projectRoot := findProjectRoot(wd, "backend_ClubYa")
		if projectRoot == "" {
			log.Fatalf("No se encontró la raíz del proyecto backend_ClubYa partiendo de: %s", wd)
		}
		logDir = filepath.Join(projectRoot, "logs")
		fmt.Println("💻 Logger ejecutándose en desarrollo local, LOG_DIR:", logDir)
	}

	fmt.Println("DEBUG: Creando directorio final...")
	// Crear el directorio de logs si no existe
	if err := os.MkdirAll(logDir, 0755); err != nil {
		fmt.Printf("ERROR: No se pudo crear el directorio de logs: %v\n", err)
		log.Fatalf("No se pudo crear el directorio de logs: %v", err)
	}
	fmt.Println("DEBUG: Directorio creado, abriendo archivo...")

	logFilename := filepath.Join(logDir, "app-"+time.Now().Format("2006-01-02")+".log")
	fmt.Println("DEBUG: Archivo de log:", logFilename)

	LogFile, err := os.OpenFile(logFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("ERROR: No se pudo abrir el archivo de logs: %v\n", err)
		log.Fatalf("No se pudo abrir el archivo de logs: %v", err)
	}
	fmt.Println("DEBUG: Archivo abierto, creando logger...")

	Logger = log.New(LogFile, "APP_LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
	fmt.Println("DEBUG: Logger inicializado exitosamente")
	// AGREGAR ESTO: Forzar flush
	LogFile.Sync()
	fmt.Println("DEBUG: Logger probado y sincronizado")
	return 1
}

func CloseLogger() {
	if LogFile != nil {
		LogFile.Close()
	}
}
