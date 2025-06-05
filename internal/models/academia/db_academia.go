package academia

import (
	"time"

	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

func ObtenerAcademias() ([]AcademiaDTO, error) {
	query := "call ingesoft.ListarAcademias()"
	rows, err := database.DB.Query(query)
	if err != nil {
		logs.Logger.Println("Error al obtener la informacion de las academias: ", err)
		return nil, err
	}
	defer rows.Close()
	var academias []AcademiaDTO
	for rows.Next() {
		var academia AcademiaDTO
		if err := rows.Scan(&academia.ID, &academia.Nombre, &academia.Descripcion, &academia.Deporte, &academia.Imagen,
			&academia.Monto, &academia.EdadMinima, &academia.Inscritos); err != nil {
			logs.Logger.Println("Error al escanear la academia deportiva: ", err)
			return nil, err
		}

		academias = append(academias, academia)
	}
	return academias, nil
}

func ObtenerAcademiaPorId(idAcademia int) (*Academia, error) {
	query := "CALL ObtenerAcademiaPorId(?)"
	row := database.DB.QueryRow(query, idAcademia)

	var academia Academia
	err := row.Scan(&academia.ID, &academia.Nombre, &academia.Descripcion, &academia.Deporte,
		&academia.Entrenador, &academia.CostoUniforme, &academia.CostoMatricula, &academia.Reglamento,
		&academia.Imagen, &academia.Indicaciones)

	if err != nil {
		logs.Logger.Println("Error al obtener la academia:", err)
		return nil, err
	}
	//leemos las tarifas
	tarifaQuery := "call ingesoft.listarTarifasAcademia(?)"
	rows, err := database.DB.Query(tarifaQuery, idAcademia)
	if err != nil {
		logs.Logger.Println("Error al obtener las tarifas de la academia:", err)
		return nil, err
	}
	defer rows.Close()
	var tarifas []TarifaAcademia
	for rows.Next() {
		var tarifa TarifaAcademia
		if err := rows.Scan(&tarifa.ID, &tarifa.Periodo, &tarifa.UnidadFrecuenciaSem, &tarifa.CantidadFrecuencia, &tarifa.TipoSocio,
			&tarifa.Monto); err != nil {
			logs.Logger.Println("Error al escanear tarifa:", err)
			return nil, err
		}
		tarifas = append(tarifas, tarifa)
	}
	academia.Tarifas = tarifas
	//ahora leemos los grupos con sus sesiones
	gruposQuery := "CALL ListarGruposAcademiaPorId(?)"
	gruposRows, err := database.DB.Query(gruposQuery, idAcademia)
	if err != nil {
		logs.Logger.Println("Error al obtener los grupos de la academia:", err)
		return nil, err
	}
	defer gruposRows.Close()
	var inicio string
	var fin string
	for gruposRows.Next() {
		var grupo GrupoAcademia
		if err := gruposRows.Scan(&grupo.ID, &grupo.Nombre, &grupo.Vacantes, &grupo.EdadMinima, &grupo.EdadMaxima,
			&grupo.Espacio.Id, &grupo.Espacio.Nombre, &grupo.Espacio.Ubicacion); err != nil {
			logs.Logger.Println("Error al escanear el grupo:", err)
			return nil, err
		}

		// Obtenemos las sesiones de cada grupo de Academia
		sesionesQuery := "CALL ListarSesionPorGruposAcademia(?)"
		sesionesRows, err := database.DB.Query(sesionesQuery, grupo.ID)
		if err != nil {
			logs.Logger.Println("Error al obtener las sesiones del grupo de la academia:", err)
			return nil, err
		}

		for sesionesRows.Next() {
			var sesion Sesion
			if err := sesionesRows.Scan(&sesion.IDsesion, &sesion.Dia, &inicio, &fin); err != nil {
				sesionesRows.Close()
				logs.Logger.Println("Error al escanear la sesion del grupo:", err)
				return nil, err
			}

			horaInicio, err := time.Parse("15:04:05", inicio)
			if err != nil {
				logs.Logger.Println("Error al parsear horaInicio: ", err)
				return nil, err
			}
			horaFin, err := time.Parse("15:04:05", fin)
			if err != nil {
				logs.Logger.Println("Error al parsear horaFin: ", err)
				return nil, err
			}
			sesion.HoraInicio = horaInicio
			sesion.HoraFin = horaFin
			grupo.Sesiones = append(grupo.Sesiones, sesion)
		}
		sesionesRows.Close()
		academia.Grupos = append(academia.Grupos, grupo)
	}

	return &academia, nil
}
