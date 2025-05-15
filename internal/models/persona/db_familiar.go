package persona

import (
	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type FamiliarResquest struct {
	IdTitular  int        `json:"id_titular"`
	Familiares []Familiar `json:"familiares"`
}

func InsertarFamiliar(f Familiar, idTitular int) error {
	query := "call ingesoft.InsertarFamiliar(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	_, err := database.DB.Exec(query, f.Nombre, f.Apellidos, f.Sexo.String(),
		f.Dni, f.FechaNacimiento, f.Telefono, f.Pais, f.Provincia, f.Distrito, f.TipoVia.String(),
		 f.Direccion, f.Referencia, f.EsConyuge, idTitular, f.MismaDireccionPostulante,f.Ciudad,f.CodigoPostal,f.TipoFamiliar.String())
	if err != nil {
		logs.Logger.Fatal("Error al insertar Familiar: ", err)
		return err
	}
	return nil
}

func RegistrarFamiliares(req FamiliarResquest) error {
	for _, familiar := range req.Familiares {
		err := InsertarFamiliar(familiar, req.IdTitular)
		if err != nil {
			logs.Logger.Println("Error al insertar familiar: ", err)
			return err
		}
	}
	return nil
}
