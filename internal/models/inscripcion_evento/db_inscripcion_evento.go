package inscripcion_evento

import (
	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

// RegistrarInscripcion permite registrar a un socio en un evento.
// Ejecuta el procedimiento almacenado 'RegistrarInscripcionEvento'.
func RegistrarInscripcion(fidPersona, idEvento, cantidadInvitados int) error {
	query := "CALL RegistrarInscripcionEvento(?, ?, ?)"
	_, err := database.DB.Exec(query, fidPersona, idEvento, cantidadInvitados)
	if err != nil {
		logs.Logger.Println("Error al registrar la inscripción al evento: ", err)
		return err
	}
	return nil
}
