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
	var id int64
	err = database.DB.QueryRow("SELECT @p_id_grupo").Scan(&id)
	if err != nil {
		logs.Logger.Println("Error al obtener el ID del nuevo grupo:", err)
		return 0, err
	}

	return id, nil
}
