package espacio

import (
	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

func InsertarEspacioSocial(es EspacioSocial) error {
	query := "call ingesoft.InsertarEspacioSocial(?, ?, ?, ?, ?, ?)"
	_, err := database.DB.Exec(query, es.Nombre, es.Codigo, es.Ubicacion, es.Capacidad, es.Costo, es.Actividad.String())
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
		if err := rows.Scan(&es.Id, &es.Nombre, &es.Codigo, &es.Ubicacion, &es.Capacidad, &es.Costo, &es.Actividad); err != nil {
			logs.Logger.Fatal("Error al escanear espacio social: ", err)
			return nil, err
		}
		/*es.Actividad, err = es.Actividad.FromString(actividad)
		if err != nil {
			logs.Logger.Fatal("Error al convertir actividad: ", err)
			return nil, err
		}*/
		espacios = append(espacios, es)
	}
	return espacios, nil
}
