package academia

import (
	"github.com/MervanXD/backend_ClubYa/database"
	grupoacademia "github.com/MervanXD/backend_ClubYa/internal/models/grupo_academia"
	"github.com/MervanXD/backend_ClubYa/internal/models/tarifas"
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
			&academia.Monto, &academia.EdadMinima, &academia.Inscritos); err != nil {
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
		&academia.Imagen, &academia.Indicaciones)

	if err != nil {
		logs.Logger.Println("Error al obtener la academia:", err)
		return nil, err
	}
	//leemos las tarifas
	tarifas, err := tarifas.ObtenerTarifasAcademiaPorId(idAcademia)
	if err != nil {
		logs.Logger.Println("Error al obtener las tarifas de la academia:", err)
		return nil, err
	}
	academia.Tarifas = tarifas
	//leemos los grupos
	grupos, err := grupoacademia.ObtenerGruposAcademiaPorId(idAcademia)
	if err != nil {
		logs.Logger.Println("Error al obtener los grupos:", err)
		return nil, err
	}
	academia.Grupos = grupos
	return &academia, nil
}
