package persona

import (
	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

func InsertarTitular(p Titular) (int, error) {
	query := "call ingesoft.InsertarTitular(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,@p_idTitular)"
	_, err := database.DB.Exec(query, p.Nombre, p.Apellidos, p.Sexo.String(),
		p.Dni, p.FechaNacimiento, p.Telefono, p.Pais, p.Provincia, p.Distrito, p.TipoVia.String(),
		 p.Direccion, p.Referencia, p.Ocupacion, p.NombreEmpresa, p.DireccionEmpresa, p.IngresoPromedio,p.EsPostulante,p.Ciudad,p.CodigoPostal)
	if err != nil {
		logs.Logger.Fatal("Error al insertar Persona: ", err)
		return -1, err
	}
	var idTitular int
	err = database.DB.QueryRow("SELECT @p_idTitular").Scan(&idTitular)
	if err != nil {
		logs.Logger.Fatal("Error al obtener idTitular: ", err)
		return -1, err
	}

	return idTitular, nil
}
