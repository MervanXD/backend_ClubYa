package persona

import (
	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

func InsertarPersona(p Persona) error{
	query := "call ingesoft.InsertarPersona(?,?,?,?,?,?,?,?,?,?,?,?)"
	_, err := database.DB.Exec(query, p.Nombre,p.Apellidos,p.Sexo.String(),
	p.Dni,p.FechaNacimiento,p.Telefono,p.Pais,p.Provincia,p.Distrito,p.TipoVia.String(),p.Direccion,p.Referencia)
	if err != nil {
		logs.Logger.Fatal("Error al insertar Persona: ", err)
		return err
	}
	return nil
}