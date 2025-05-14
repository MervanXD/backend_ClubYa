package espacio

import (
	"context"
	"time"

	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

func InsertarEspacioSocial(es EspacioSocial) error {
	query := "call ingesoft.InsertarEspacioSocial(?, ?, ?, ?, ?, ?,?,?)"
	_, err := database.DB.Exec(query, es.Nombre, es.Codigo, es.Ubicacion, es.Capacidad, es.Costo, es.Imagen, es.Reglamento, es.Actividad.String())
	if err != nil {
		logs.Logger.Fatal("Error al insertar espacio social: ", err)
		return err
	}
	return nil
}

func ObtenerEspaciosSociales() ([]EspacioSocial, error) {
	query := "call ingesoft.ObtenerEspaciosSociales()"
	rows, err := database.DB.Query(query)
	if err != nil {
		logs.Logger.Fatal("Error al obtener espacios sociales: ", err)
		return nil, err
	}
	defer rows.Close()
	var espacios []EspacioSocial
	for rows.Next() {
		var es EspacioSocial
		//var actividad string
		if err := rows.Scan(&es.Id, &es.Nombre, &es.Codigo, &es.Ubicacion, &es.Capacidad, &es.Costo, &es.Imagen, &es.Reglamento, &es.Actividad); err != nil {
			logs.Logger.Fatal("Error al escanear espacio social: ", err)
			return nil, err
		}
		espacios = append(espacios, es)
	}
	return espacios, nil
}

func ObtenerEspacioSocialPorID(ctx context.Context, id int) (*EspacioSocial, error) {
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
