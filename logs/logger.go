package logs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	LogFile *os.File
	Logger  *log.Logger
)

// Busca hacia arriba hasta encontrar la carpeta backend_ClubYa
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

func InitLogger() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error al obtener el directorio de trabajo: %v", err)
	}

	projectRoot := findProjectRoot(wd, "backend_ClubYa")
	if projectRoot == "" {
		log.Fatalf("No se encontró la raíz del proyecto backend_ClubYa partiendo de: %s", wd)
	}

	logDir := filepath.Join(projectRoot, "logs")

	logFilename := filepath.Join(logDir, "app-"+time.Now().Format("2006-01-02")+".log")
	LogFile, err := os.OpenFile(logFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(logDir)
		log.Fatalf("No se pudo abrir el archivo de logs: %v\nDir: %s", err, logFilename)
	}

	Logger = log.New(LogFile, "APP_LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Siempre cerrar al salir
func CloseLogger() {
	if LogFile != nil {
		LogFile.Close()
	}
}
