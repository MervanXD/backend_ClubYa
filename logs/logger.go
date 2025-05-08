package logs

import (
	"log"
	"os"
	"time"
)

var (
	LogFile *os.File
	Logger  *log.Logger
)

func InitLogger() {

	logFilename := "../../logs/app-" + time.Now().Format("2006-01-02") + ".log"
	LogFile, err := os.OpenFile(logFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("No se pudo abrir el archivo de logs: %v", err)
	}

	Logger = log.New(LogFile, "APP_LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Siempre cerrar al salir
func CloseLogger() {
	if LogFile != nil {
		LogFile.Close()
	}
}
