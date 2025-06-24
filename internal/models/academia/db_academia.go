package academia

import (
	"github.com/MervanXD/backend_ClubYa/database"
	grupoacademia "github.com/MervanXD/backend_ClubYa/internal/models/grupo_academia"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type academiaRespositoryDB struct{}

func NewAcademiaRepositoryDB() AcademiaRepository {
	return &academiaRespositoryDB{}
}

func (r *academiaRespositoryDB) ObtenerAcademias() ([]AcademiaDTO, error) {
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
			&academia.Monto, &academia.FechaInicio, &academia.FechaFin, &academia.EdadMinima, &academia.Inscritos); err != nil {
			logs.Logger.Println("Error al escanear la academia deportiva: ", err)
			return nil, err
		}

		academias = append(academias, academia)
	}
	return academias, nil
}

func (r *academiaRespositoryDB) ObtenerAcademiaPorId(idAcademia int) (*Academia, error) {
	query := "CALL ObtenerAcademiaPorId(?)"
	row := database.DB.QueryRow(query, idAcademia)

	var academia Academia
	err := row.Scan(&academia.ID, &academia.Nombre, &academia.Descripcion, &academia.Deporte,
		&academia.Entrenador, &academia.CostoUniforme, &academia.CostoMatricula, &academia.Reglamento,
		&academia.Imagen, &academia.Indicaciones, &academia.FechaInicio, &academia.FechaFin)

	if err != nil {
		logs.Logger.Println("Error al obtener la academia:", err)
		return nil, err
	}

	//leemos los grupos
	repo := grupoacademia.NewGrupoAcademiaRepositoryDB()
	grupos, err := repo.ObtenerGruposAcademiaPorId(idAcademia)
	if err != nil {
		logs.Logger.Println("Error al obtener los grupos:", err)
		return nil, err
	}
	academia.Grupos = grupos
	return &academia, nil
}

func (r *academiaRespositoryDB) InsertarAcademia(academia *Academia) error {
	query := "CALL InsertarAcademia(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := database.DB.Exec(query, academia.Nombre, academia.Descripcion, academia.Deporte.String(),
		academia.Entrenador, academia.CostoUniforme, academia.CostoMatricula, academia.Reglamento,
		academia.Imagen, academia.Indicaciones, academia.FechaInicio, academia.FechaFin)
	if err != nil {
		logs.Logger.Println("Error al insertar la academia:", err)
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		logs.Logger.Println("Error al obtener el ID de la nueva academia:", err)
		return err
	}
	repo := grupoacademia.NewGrupoAcademiaRepositoryDB()
	for _, grupo := range academia.Grupos {
		grupo.IdAcademia = id
		if _, err := repo.InsertarGrupoAcademia(&grupo); err != nil {
			logs.Logger.Println("Error al insertar el grupo de la academia:", err)
			return err
		}
	}
	return nil
}
