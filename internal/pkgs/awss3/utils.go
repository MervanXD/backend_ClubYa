package awss3

import (
	"net/http"
	"path/filepath"
	"strings"
)

func DetectContentType(fileData []byte, fileName string) string {
	// Intentar detectar por contenido del archivo
	contentType := http.DetectContentType(fileData)

	// Si no se detecta bien, usar extensión del archivo
	if contentType == "application/octet-stream" {
		ext := strings.ToLower(filepath.Ext(fileName))
		switch ext {
		case ".jpg", ".jpeg":
			return "image/jpeg"
		case ".png":
			return "image/png"
		case ".gif":
			return "image/gif"
		case ".pdf":
			return "application/pdf"
		case ".doc":
			return "application/msword"
		case ".docx":
			return "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
		case ".xls":
			return "application/vnd.ms-excel"
		case ".xlsx":
			return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
		case ".txt":
			return "text/plain"
		default:
			return "application/octet-stream"
		}
	}

	return contentType
}

// Función mejorada que detecta automáticamente el content-type
func UploadFromBytesAuto(fileData []byte, fileName string) (string, error) {
	contentType := DetectContentType(fileData, fileName)
	return UploadFromBytes(fileData, fileName, contentType)
}

// Extrae el nombre del archivo de una URL de S3
func ExtractFileNameFromURL(url string) string {
    // URL típica: https://clubya-archivos.s3.amazonaws.com/uploads/1234567890_documento.pdf
    parts := strings.Split(url, "/")
    if len(parts) > 0 {
        fullFileName := parts[len(parts)-1]
        // Remover el timestamp del nombre: 1234567890_documento.pdf -> documento.pdf
        if underscoreIndex := strings.Index(fullFileName, "_"); underscoreIndex != -1 {
            return fullFileName[underscoreIndex+1:]
        }
        return fullFileName
    }
    return ""
}
