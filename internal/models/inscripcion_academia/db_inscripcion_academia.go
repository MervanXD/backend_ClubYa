package inscripcionacademia

import (
	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/internal/models/sesiones"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type inscripcionAcademiaRepositoryDB struct{}

func NewInscripcionAcademiaRepositoryDB() InscripcionAcademiaRepository {
	return &inscripcionAcademiaRepositoryDB{}
}

func (r *inscripcionAcademiaRepositoryDB) RegistrarInscripcionAcademia(idPersonaint int, idGrupo int, idTarifa int, uniforme int, costo_total float64, idTitular int) error {
	query := "CALL RegistrarInscripcionAcademia(?, ?, ?,?,?,?)"
	_, err := database.DB.Exec(query, idPersonaint, idGrupo, idTarifa, uniforme, costo_total, idTitular)
	if err != nil {
		logs.Logger.Println("Error al registrar la inscripción al evento: ", err)
		return err
	}
	return nil
}

func (r *inscripcionAcademiaRepositoryDB) ObtenerFamiliaresInscritosAcademia(idSocio int) ([]InscritoAcademiaDTO, error) {
	inscritosQuery := "CALL ListarInscripcionesFamiliares(?)"
	inscritosRows, err := database.DB.Query(inscritosQuery, idSocio)
	if err != nil {
		logs.Logger.Println("Error al obtener las personas inscritas:", err)
		return nil, err
	}
	defer inscritosRows.Close()
	var inscritos []InscritoAcademiaDTO
	for inscritosRows.Next() {
		var inscrito InscritoAcademiaDTO
		if err := inscritosRows.Scan(&inscrito.IdPersona, &inscrito.NombrePersona, &inscrito.ApellidoPersona, &inscrito.NombreGrupo,
			&inscrito.IdGrupo, &inscrito.NombreAcademia, &inscrito.FechaInicio, &inscrito.FechaFin, &inscrito.EstadoInscripcion); err != nil {
			logs.Logger.Println("Error al escanear a la persona inscrita:", err)
			return nil, err
		}
		//leemos sus sesiones
		sesiones, err := sesiones.ObtenerSesionesGrupo(inscrito.IdGrupo)
		if err != nil {
			logs.Logger.Println("Error al obtener las sesiones del grupo :", err)
			return nil, err
		}
		inscrito.Sesiones = sesiones
		inscritos = append(inscritos, inscrito)
	}

	return inscritos, nil
}
