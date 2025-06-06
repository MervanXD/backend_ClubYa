package grupoacademia

import (
	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/internal/models/sesiones"
	"github.com/MervanXD/backend_ClubYa/logs"
)

func ObtenerGruposAcademiaPorId(idAcademia int) ([]GrupoAcademia, error) {
	gruposQuery := "CALL ListarGruposAcademiaPorId(?)"
	gruposRows, err := database.DB.Query(gruposQuery, idAcademia)
	if err != nil {
		logs.Logger.Println("Error al obtener los grupos de la academia:", err)
		return nil, err
	}
	defer gruposRows.Close()
	var grupos []GrupoAcademia
	for gruposRows.Next() {
		var grupo GrupoAcademia
		if err := gruposRows.Scan(&grupo.ID, &grupo.Nombre, &grupo.Vacantes, &grupo.EdadMinima, &grupo.EdadMaxima,
			&grupo.Espacio.Id, &grupo.Espacio.Nombre, &grupo.Espacio.Ubicacion); err != nil {
			logs.Logger.Println("Error al escanear el grupo:", err)
			return nil, err
		}
		//leemos sus sesiones
		sesiones, err := sesiones.ObtenerSesionesGrupo(grupo.ID)
		if err != nil {
			logs.Logger.Println("Error al obtener las sesiones del grupo :", err)
			return nil, err
		}
		grupo.Sesiones = sesiones
		grupos = append(grupos, grupo)
	}

	return grupos, nil
}
