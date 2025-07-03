package espacio

import (
	"fmt"
	"strings"

	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type canchaRespositoryDB struct{}

func NewCanchaRepositoryDB() CanchaRepository {
	return &canchaRespositoryDB{}
}

func (r *canchaRespositoryDB) ObtenerCanchasHorarios() ([]CanchaHorarioDTO, error) {
	query := "call ingesoft.listarCanchasHorarios()"
	rows, err := database.DB.Query(query)
	if err != nil {
		logs.Logger.Println("Error al obtener los horarios de las lozas deportivas: ", err)
		return nil, err
	}
	defer rows.Close()
	var canchasHorarios []CanchaHorarioDTO
	for rows.Next() {
		var es CanchaHorarioDTO
		if err := rows.Scan(&es.Espacio.Id, &es.Espacio.Codigo, &es.Espacio.Nombre, &es.Espacio.Imagen, &es.Espacio.Deporte,
			&es.Espacio.Ubicacion, &es.Espacio.Capacidad, &es.Espacio.Costo, &es.Fecha, &es.HoraInicio, &es.HoraFinal,
			&es.Estado, &es.IdHorario, &es.IdBloque); err != nil {
			logs.Logger.Println("Error al escanear el horario de la loza deportiva: ", err)
			return nil, err
		}
		canchasHorarios = append(canchasHorarios, es)
	}
	return canchasHorarios, nil
}

func (r *canchaRespositoryDB) InsertarCancha(c Cancha) error {
	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	query := "CALL ingesoft.InsertarCancha(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = tx.Exec(query,
		c.Codigo,
		c.Nombre,
		c.Ubicacion.String(),
		c.Capacidad,
		c.Costo,
		c.Imagen,
		c.Reglamento,
		1, // EstadoEspacio activo por defecto
		c.DuracionBloque,
		c.Deporte.String(),
	)
	if err != nil {
		logs.Logger.Println("Error al insertar espacio social: ", err)
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}


func (r *canchaRespositoryDB) ActualizarParcial(id int, dto CanchaUpdateDTO) error {
	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	// ---------- Actualiza tabla Espacio ----------
	espacioSet := []string{}
	args := []interface{}{}

	if dto.Nombre != nil {
		espacioSet = append(espacioSet, "nombre = ?")
		args = append(args, *dto.Nombre)
	}
	if dto.Ubicacion != nil {
		espacioSet = append(espacioSet, "ubicacion = ?")
		args = append(args, dto.Ubicacion.String())
	}
	if dto.Capacidad != nil {
		espacioSet = append(espacioSet, "capacidad = ?")
		args = append(args, *dto.Capacidad)
	}
	if dto.Costo != nil {
		espacioSet = append(espacioSet, "costo = ?")
		args = append(args, *dto.Costo)
	}
	if dto.Codigo != nil {
		espacioSet = append(espacioSet, "codigo = ?")
		args = append(args, *dto.Codigo)
	}
	if dto.Reglamento != nil {
		espacioSet = append(espacioSet, "reglamento = ?")
		args = append(args, *dto.Reglamento)
	}
	if dto.Imagen != nil {
		espacioSet = append(espacioSet, "imagen = ?")
		args = append(args, *dto.Imagen)
	}
	if dto.DuracionBloque != nil {
		espacioSet = append(espacioSet, "duracion_bloque = ?")
		args = append(args, *dto.DuracionBloque)
	}
	if dto.EstadoEspacio != nil {
		espacioSet = append(espacioSet, "estadoEspacio = ?")
		args = append(args, *dto.EstadoEspacio)
	}

	if len(espacioSet) > 0 {
		query := fmt.Sprintf("UPDATE Espacio SET %s WHERE idEspacio = ?", strings.Join(espacioSet, ", "))
		args = append(args, id)
		if _, err := tx.Exec(query, args...); err != nil {
			tx.Rollback()
			return fmt.Errorf("error actualizando espacio: %w", err)
		}
	}

	// ---------- Actualiza tabla Canchas ----------
	if dto.Deporte != nil {
		query := "UPDATE Canchas SET deporte = ? WHERE fid_Espacio = ?"
		if _, err := tx.Exec(query, dto.Deporte.String(), id); err != nil {
			tx.Rollback()
			return fmt.Errorf("error actualizando deporte: %w", err)
		}
	}

	// ---------- Commit final ----------
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error al hacer commit: %w", err)
	}

	return nil
}



func (r *canchaRespositoryDB) ObtenerCanchasConfiguracion() ([]Cancha, error) {
	query := "call ingesoft.ListarCanchasConfiguracion()"
	rows, err := database.DB.Query(query)
	if err != nil {
		logs.Logger.Println("Error al obtener las canchas: ", err)
		return nil, err
	}
	defer rows.Close()
	var canchas []Cancha
	for rows.Next() {
		var cancha Cancha
		//var actividad string
		if err := rows.Scan(&cancha.Id, &cancha.Codigo, &cancha.Nombre, &cancha.Ubicacion, &cancha.Capacidad,
			&cancha.Costo, &cancha.Reglamento, &cancha.Imagen, &cancha.EstadoEspacio, &cancha.DuracionBloque,
			&cancha.Deporte); err != nil {
			logs.Logger.Println("Error al escanear la cancha: ", err)
			return nil, err
		}
		canchas = append(canchas, cancha)
	}
	return canchas, nil

}
