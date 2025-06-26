package grupoacademia

import (
	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/internal/models/sesiones"
	"github.com/MervanXD/backend_ClubYa/internal/models/tarifas"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type grupoAcademiaRepositoryDB struct{}

func NewGrupoAcademiaRepositoryDB() GrupoAcademiaRepository {
	return &grupoAcademiaRepositoryDB{}
}

func (r *grupoAcademiaRepositoryDB) ObtenerGruposAcademiaPorId(idAcademia int) ([]GrupoAcademia, error) {
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
			&grupo.Espacio.Id, &grupo.Espacio.Nombre, &grupo.Espacio.Ubicacion, &grupo.Inscritos, &grupo.Espacio.EstadoEspacio); err != nil {
			logs.Logger.Println("Error al escanear el grupo:", err)
			return nil, err
		}
		//leemos sus sesiones
		repo := sesiones.NewSesionRepositoryDB()
		sesiones, err := repo.ObtenerSesionesGrupo(grupo.ID)
		if err != nil {
			logs.Logger.Println("Error al obtener las sesiones del grupo :", err)
			return nil, err
		}
		grupo.Sesiones = sesiones

		//leemos sus tarifas
		repoTarifa := tarifas.NewTarifaAcademiaRepositoryDB()
		tarifas, err := repoTarifa.ObtenerTarifasAcademiaPorId(grupo.ID)
		if err != nil {
			logs.Logger.Println("Error al obtener las tarifas del grupo :", err)
			return nil, err
		}
		grupo.Tarifas = tarifas

		grupos = append(grupos, grupo)
	}

	return grupos, nil
}

func (r *grupoAcademiaRepositoryDB) InsertarGrupoAcademia(grupo *GrupoAcademia) (int64, error) {
	query := "CALL InsertarGrupoAcademia(?,?,?,?,?,?,?,@p_id_grupo)"
	_, err := database.DB.Exec(query, grupo.Nombre, grupo.Vacantes, grupo.EdadMinima, grupo.EdadMaxima,
		grupo.Espacio.Id, 0, grupo.IdAcademia) //inscritos inicialmente es 0
	if err != nil {
		logs.Logger.Println("Error al insertar el grupo de la academia:", err)
		return 0, err
	}
	var id int
	err = database.DB.QueryRow("SELECT @p_id_grupo").Scan(&id)
	if err != nil {
		logs.Logger.Println("Error al obtener el ID del nuevo grupo:", err)
		return 0, err
	}
	// Insertar las tarifas del grupo
	repoTarifa := tarifas.NewTarifaAcademiaRepositoryDB()
	for _, tarifa := range grupo.Tarifas {
		tarifa.IDGrupo = id
		if _, err := repoTarifa.InsertarTarifaAcademia(&tarifa); err != nil {
			logs.Logger.Println("Error al insertar la tarifa del grupo:", err)
			return 0, err
		}
	}
	// Insertar las sesiones del grupo
	repoSesion := sesiones.NewSesionRepositoryDB()
	for _, sesion := range grupo.Sesiones {
		sesion.IdGrupo = id
		if _, err := repoSesion.InsertarSesion(&sesion); err != nil {
			logs.Logger.Println("Error al insertar la sesión del grupo:", err)
			return 0, err
		}
	}

	return int64(id), nil
}

func (r *grupoAcademiaRepositoryDB) ActualizarGrupoAcademiaParcial(grupo *GrupoAcademiaUpdate) error {
	setClauses := []string{}
	args := []interface{}{}

	if grupo.Nombre != nil {
		setClauses = append(setClauses, "nombre = ?")
		args = append(args, *grupo.Nombre)
	}
	if grupo.Vacantes != nil {
		setClauses = append(setClauses, "vacantes = ?")
		args = append(args, *grupo.Vacantes)
	}
	if grupo.EdadMinima != nil {
		setClauses = append(setClauses, "edadMinima = ?")
		args = append(args, *grupo.EdadMinima)
	}
	if grupo.EdadMaxima != nil {
		setClauses = append(setClauses, "edadMaxima = ?")
		args = append(args, *grupo.EdadMaxima)
	}
	if grupo.Espacio != nil {
		setClauses = append(setClauses, "fid_Espacio = ?")
		args = append(args, *grupo.Espacio)
	}
	if len(setClauses) == 0 {
		logs.Logger.Println("No se proporcionaron campos para actualizar el grupo de la academia")
		return nil // No hay nada que actualizar
	}
	query := "UPDATE grupo_academia SET " + setClauses[0]
	for i := 1; i < len(setClauses); i++ {
		query += ", " + setClauses[i]
	}
	query += " WHERE id = ?"
	args = append(args, grupo.ID)
	_, err := database.DB.Exec(query, args...)
	if err != nil {
		logs.Logger.Println("Error al actualizar el grupo de la academia:", err)
		return err
	}
	if grupo.Sesiones != nil {
		repoSesion := sesiones.NewSesionRepositoryDB()
		for _, sesion := range *grupo.Sesiones {
			sesion.IdGrupo = grupo.ID
			if err := repoSesion.ActualizarSesionParcial(&sesion); err != nil {
				logs.Logger.Println("Error al actualizar la sesión del grupo:", err)
				return err
			}
		}
	}

	return nil
}
