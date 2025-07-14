package sesiones

import (
	"database/sql"
	"fmt"

	"github.com/MervanXD/backend_ClubYa/database"
	detalledisponibilidad "github.com/MervanXD/backend_ClubYa/internal/models/detalle_disponibilidad"
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

func (r *sesionRepositoryDB) InsertarSesionTx(tx *sql.Tx, sesion *Sesion) (int64, error) {
	query := "CALL InsertarSesion(?,?,?,?)"
	result, err := tx.Exec(query, sesion.IdGrupo, sesion.Dia.String(), sesion.HoraInicio, sesion.HoraFin)
	fmt.Println(sesion.Dia)
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

func (r *sesionRepositoryDB) ActualizarSesionParcial(sesion *SesionUpdate) error {
	setClauses := []string{}
	args := []interface{}{}
	if sesion.Dia != nil {
		setClauses = append(setClauses, "dia = ?")
		args = append(args, *sesion.Dia)
	}
	if sesion.HoraFin != nil {
		setClauses = append(setClauses, "horaFin = ?")
		args = append(args, *sesion.HoraFin)
	}
	if sesion.HoraInicio != nil {
		setClauses = append(setClauses, "horaInicio = ?")
		args = append(args, *sesion.HoraInicio)
	}
	query := "UPDATE Sesion SET " + setClauses[0]
	for i := 1; i < len(setClauses); i++ {
		query += ", " + setClauses[i]
	}
	query += " WHERE id = ?"
	args = append(args, sesion.IDSesion)
	_, err := database.DB.Exec(query, args...)
	if err != nil {
		logs.Logger.Println("Error al actualizar la sesion:", err)
		return err
	}
	//actualizar los detalle disponibilidad en caso se haya modificado el dia o hora inicio y hora fin
	if sesion.Dia != nil || sesion.HoraInicio != nil || sesion.HoraFin != nil {
		repoDisponibilidad := detalledisponibilidad.NewDetalleDisponibilidadRepositoryDB()
		var dia, horaInicio, horaFin string
		if sesion.Dia != nil {
			dia = *sesion.Dia
		}
		if sesion.HoraInicio != nil {
			horaInicio = *sesion.HoraInicio
		}
		if sesion.HoraFin != nil {
			horaFin = *sesion.HoraFin
		}
		if err := repoDisponibilidad.ActualizarDisponibilidadPorSesion(int64(*sesion.IDSesion), dia, horaInicio, horaFin, *sesion.IdGrupo); err != nil {
			logs.Logger.Println("Error al actualizar la disponibilidad de la sesión:", err)
			return err
		}
	}

	return nil

}
