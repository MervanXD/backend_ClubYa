package academia

import (
	"fmt"
	"strings"

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
	query := "CALL InsertarAcademia(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,@c_id_academia)"
	_, err := database.DB.Exec(query, academia.Nombre, academia.Descripcion, academia.Deporte.String(),
		academia.Entrenador, academia.CostoUniforme, academia.CostoMatricula, academia.Reglamento,
		academia.Imagen, academia.Indicaciones, academia.FechaInicio, academia.FechaFin)
	if err != nil {
		logs.Logger.Println("Error al insertar la academia:", err)
		return err
	}
	// Obtenemos el ID de la nueva academia
	var id int64
	err = database.DB.QueryRow("SELECT @c_id_academia").Scan(&id)
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

func (r *academiaRespositoryDB) ListarAcademiasGenerales() ([]AcademiaListarRequest, error) {
	query := "CALL ListarAcademiasAdmin()"
	rows, err := database.DB.Query(query)
	if err != nil {
		logs.Logger.Println("Error al listar las academias generales:", err)
		return nil, err
	}
	defer rows.Close()

	var academias []AcademiaListarRequest
	for rows.Next() {
		var academia AcademiaListarRequest
		if err := rows.Scan(&academia.ID, &academia.Nombre, &academia.Deporte, &academia.Entrenador,
			&academia.Imagen, &academia.FechaInicio, &academia.FechaFin, &academia.EdadMinima, &academia.EdadMaxima,
			&academia.Vacantes, &academia.Inscritos); err != nil {
			logs.Logger.Println("Error al escanear la academia general:", err)
			return nil, err
		}
		academias = append(academias, academia)
	}
	return academias, nil
}

func (r *academiaRespositoryDB) ActualizarParcialAcademia(id int, dto AcademiaUpdateDTO) error {
	setClauses := []string{}
	args := []interface{}{}

	if dto.Nombre != nil {
		setClauses = append(setClauses, "nombre = ?")
		args = append(args, *dto.Nombre)
	}
	if dto.Descripcion != nil {
		setClauses = append(setClauses, "descripcion = ?")
		args = append(args, *dto.Descripcion)
	}
	if dto.Deporte != nil {
		setClauses = append(setClauses, "deporte = ?")
		args = append(args, dto.Deporte.String())
	}
	if dto.Entrenador != nil {
		setClauses = append(setClauses, "entrenador = ?")
		args = append(args, *dto.Entrenador)
	}
	if dto.CostoUniforme != nil {
		setClauses = append(setClauses, "costoUniforme = ?")
		args = append(args, *dto.CostoUniforme)
	}
	if dto.CostoMatricula != nil {
		setClauses = append(setClauses, "costoMatricula = ?")
		args = append(args, *dto.CostoMatricula)
	}
	if dto.Reglamento != nil {
		setClauses = append(setClauses, "reglamento = ?")
		args = append(args, *dto.Reglamento)
	}
	if dto.Imagen != nil {
		setClauses = append(setClauses, "imagen = ?")
		args = append(args, *dto.Imagen)
	}
	if dto.Indicaciones != nil {
		setClauses = append(setClauses, "indicaciones = ?")
		args = append(args, *dto.Indicaciones)
	}
	if dto.FechaInicio != nil {
		setClauses = append(setClauses, "fechaInicio = ?")
		args = append(args, *dto.FechaInicio)
	}
	if dto.FechaFin != nil {
		setClauses = append(setClauses, "fechaFin = ?")
		args = append(args, *dto.FechaFin)
	}

	if len(setClauses) > 0 {
		query := fmt.Sprintf("UPDATE Academias SET %s WHERE idAcademia = ?", strings.Join(setClauses, ", "))
		args = append(args, id)
		_, err := database.DB.Exec(query, args...)
		if err != nil {
			return fmt.Errorf("error actualizando academia: %w", err)
		}
	}

	return nil
}
