package servicios

import (
	"fmt"
	"net/smtp"

	"github.com/MervanXD/backend_ClubYa/logs"
)

// Configuración del servidor SMTP y credenciales
const (
	smtpHost    = "smtp.gmail.com"
	smtpPort    = "587"
	senderEmail = "codehere33@gmail.com" // Tu correo
	appPassword = "telg yafv ecjr wsuw "            // Contraseña de aplicación de Gmail
)

// EnviarCorreo envía un correo simple a una o más direcciones
func EnviarCorreo(to []string, subject, body string) error {
	auth := smtp.PlainAuth("", senderEmail, appPassword, smtpHost)

	// Preparar el mensaje
	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s\r\n",
		to[0], subject, body)) // solo muestra el primer destinatario en el encabezado

	addr := smtpHost + ":" + smtpPort
	err := smtp.SendMail(addr, auth, senderEmail, to, msg)
	if err != nil {
		logs.Logger.Println("Error al enviar correo: ", err)
		return fmt.Errorf("error al enviar correo: %w", err)
	}
	return nil
}
