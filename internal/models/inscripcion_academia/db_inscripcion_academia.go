package inscripcionacademia

import (
	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

func RegistrarInscripcionAcademia(idPersonaint int, idGrupo int, idTarifa int, uniforme int) error {
	query := "CALL RegistrarInscripcionAcademia(?, ?, ?,?)"
	_, err := database.DB.Exec(query, idPersonaint, idGrupo, idTarifa, uniforme)
	if err != nil {
		logs.Logger.Println("Error al registrar la inscripción al evento: ", err)
		return err
	}
	return nil
}
