package espacio

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type espacioSocialRespositoryDB struct{}

func NewEspacioSocialRepositoryDB() EspacioSocialRepository {
	return &espacioSocialRespositoryDB{}
}

func (r *espacioSocialRespositoryDB) InsertarEspacioSocial(es EspacioSocial) error {
	query := "call ingesoft.InsertarEspacioSocial(?, ?, ?, ?, ?, ?, ?, ?, ?,?)"
	_, err := database.DB.Exec(query,
		es.Codigo,
		es.Nombre,
		es.Ubicacion.String(),
		es.Capacidad,
		es.Costo,
		es.Reglamento,
		es.Imagen,
		0,
		es.Actividad,
		es.DuracionBloque)
	if err != nil {
		logs.Logger.Println("Error al insertar espacio social: ", err)
		return err
	}
	return nil
}

func (r *espacioSocialRespositoryDB) ActualizarParcial(id int, dto EspacioSocialUpdateDTO) error {
	// ---------- Actualiza tabla Espacio ----------
	espacioSet := []string{}
	args := []interface{}{}

	if dto.Nombre != nil {
		espacioSet = append(espacioSet, "nombre = ?")
		args = append(args, *dto.Nombre)
	}
	if dto.Ubicacion != nil {
		espacioSet = append(espacioSet, "ubicacion = ?")
		args = append(args, dto.Ubicacion.String()) // si es enum con método String()
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
		espacioSet = append(espacioSet, "estado_espacio = ?")
		args = append(args, *dto.EstadoEspacio)
	}

	if len(espacioSet) > 0 {
		query := fmt.Sprintf("UPDATE Espacio SET %s WHERE idEspacio = ?", strings.Join(espacioSet, ", "))
		args = append(args, id)
		_, err := database.DB.Exec(query, args...)
		if err != nil {
			return fmt.Errorf("error actualizando espacio: %w", err)
		}
	}

	// ---------- Actualiza tabla EspacioSocial ----------
	if dto.Actividad != nil {
		query := "UPDATE EspacioSocial SET actividad = ? WHERE fid_Espacio = ?"
		_, err := database.DB.Exec(query, dto.Actividad.String(), id)
		if err != nil {
			return fmt.Errorf("error actualizando actividad: %w", err)
		}
	}

	return nil
}

func (r *espacioSocialRespositoryDB) ObtenerEspaciosSociales() ([]EspacioSocial, error) {
	query := "call ingesoft.ObtenerEspaciosSociales()"
	rows, err := database.DB.Query(query)
	if err != nil {
		logs.Logger.Println("Error al obtener espacios sociales: ", err)
		return nil, err
	}
	defer rows.Close()
	var espacios []EspacioSocial
	for rows.Next() {
		var es EspacioSocial
		//var actividad string
		if err := rows.Scan(&es.Id, &es.Nombre, &es.Codigo, &es.Ubicacion, &es.Capacidad, &es.Costo, &es.Imagen, &es.Reglamento, &es.Actividad); err != nil {
			logs.Logger.Println("Error al escanear espacio social: ", err)
			return nil, err
		}
		espacios = append(espacios, es)
	}
	return espacios, nil
}

func (r *espacioSocialRespositoryDB) ObtenerEspacioSocialPorID(ctx context.Context, id int) (*EspacioSocial, error) {
	query := "call ingesoft.ObtenerEspacioSocialPorID(?)"

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var es EspacioSocial
	err := database.DB.QueryRowContext(ctx, query, id).Scan(&es.Id, &es.Nombre, &es.Codigo, &es.Ubicacion, &es.Capacidad, &es.Costo, &es.Imagen, &es.Reglamento, &es.Actividad)
	if err != nil {
		return nil, err
	}
	return &es, nil
}

func (r *espacioSocialRespositoryDB) ObtenerEspaciosSocialesHorarios() ([]EspacioSocialHorarioDTO, error) {
	query := "call ingesoft.listarEspaciosSocialesHorarios()"
	rows, err := database.DB.Query(query)
	if err != nil {
		logs.Logger.Println("Error al obtener los horarios de los espacios sociales: ", err)
		return nil, err
	}
	defer rows.Close()
	var espaciosHorarios []EspacioSocialHorarioDTO
	for rows.Next() {
		var es EspacioSocialHorarioDTO
		//var actividad string
		if err := rows.Scan(&es.Espacio.Id, &es.Espacio.Codigo, &es.Espacio.Nombre, &es.Espacio.Actividad,
			&es.Espacio.Ubicacion, &es.Espacio.Capacidad, &es.Espacio.Costo, &es.Fecha, &es.HoraInicio, &es.HoraFinal,
			&es.Estado, &es.IdHorario, &es.IdBloque); err != nil {
			logs.Logger.Println("Error al escanear el horario del espacio social: ", err)
			return nil, err
		}
		espaciosHorarios = append(espaciosHorarios, es)
	}
	return espaciosHorarios, nil
}

func (r *espacioSocialRespositoryDB) ObtenerEspaciosSocialesConfiguracion() ([]EspacioSocial, error) {
	query := "call ingesoft.ListarEspaciosSocialesConfiguracion()"
	rows, err := database.DB.Query(query)
	if err != nil {
		logs.Logger.Println("Error al obtener espacios sociales: ", err)
		return nil, err
	}
	defer rows.Close()
	var espacios []EspacioSocial
	for rows.Next() {
		var es EspacioSocial
		//var actividad string
		if err := rows.Scan(&es.Id, &es.Nombre, &es.Codigo, &es.Actividad); err != nil {
			logs.Logger.Println("Error al escanear espacio social: ", err)
			return nil, err
		}
		espacios = append(espacios, es)
	}
	return espacios, nil
}
