package sesiones

import (
	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type sesionRepositoryDB struct{}

func NewSesionRepositoryDB() SesionRepository {
	return &sesionRepositoryDB{}
}

func (r *sesionRepositoryDB) ObtenerSesionesGrupo(idGrupo int) ([]Sesion, error) {
	sesionesQuery := "CALL ListarSesionPorGruposAcademia(?)"
	sesionesRows, err := database.DB.Query(sesionesQuery, idGrupo)
	if err != nil {
		logs.Logger.Println("Error al obtener las sesiones del grupo de la academia:", err)
		return nil, err
	}
	defer sesionesRows.Close()
	var sesiones []Sesion
	for sesionesRows.Next() {
		var sesion Sesion
		if err := sesionesRows.Scan(&sesion.IDsesion, &sesion.Dia, &sesion.HoraInicio, &sesion.HoraFin); err != nil {
			sesionesRows.Close()
			logs.Logger.Println("Error al escanear la sesion del grupo:", err)
			return nil, err
		}
		// Convertimos la hora de formato "2006-01-02T15:04:05Z07:00" a "15:04:05"
		if len(sesion.HoraInicio) >= 19 {
			sesion.HoraInicio = sesion.HoraInicio[11:19]
		}
		if len(sesion.HoraFin) >= 19 {
			sesion.HoraFin = sesion.HoraFin[11:19]
		}
		sesiones = append(sesiones, sesion)
	}
	return sesiones, nil
}

func (r *sesionRepositoryDB) InsertarSesion(sesion *Sesion) (int64, error) {
	query := "CALL InsertarSesion(?,?,?,?)"
	result, err := database.DB.Exec(query, sesion.IdGrupo, sesion.Dia, sesion.HoraInicio, sesion.HoraFin)
	if err != nil {
		logs.Logger.Println("Error al insertar la sesión:", err)
		return 0, err
	}
	idSesion, err := result.LastInsertId()
	if err != nil {
		logs.Logger.Println("Error al obtener el ID de la sesión insertada:", err)
		return 0, err
	}
	return idSesion, nil
}
