package academia

import (
	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type AcademiaDTO struct {
	ID          int64         `json:"id"`
	Nombre      string        `json:"nombre"`
	Descripcion string        `json:"descripcion"`
	Deporte     tipos.Deporte `json:"deporte"`
	Imagen      []byte        `json:"imagen"`
	Monto       float64       `json:"monto"`
	EdadMinima  int64         `json:"edad_minima"`
}

func ObtenerAcademias() ([]AcademiaDTO, error) {
	query := "call ingesoft.ListarAcademias()"
	rows, err := database.DB.Query(query)
	if err != nil {
		logs.Logger.Println("Error al obtener la informacion de las academias: ", err)
		return nil, err
	}
	defer rows.Close()
	var academias []AcademiaDTO
	for rows.Next() {
		var academia AcademiaDTO
		if err := rows.Scan(&academia.ID, &academia.Nombre, &academia.Descripcion, &academia.Deporte, &academia.Imagen,
			&academia.Monto, &academia.EdadMinima); err != nil {
			logs.Logger.Println("Error al escanear la academia deportiva: ", err)
			return nil, err
		}

		academias = append(academias, academia)
	}
	return academias, nil
}
