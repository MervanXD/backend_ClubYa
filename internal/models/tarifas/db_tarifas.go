package tarifas

import (
	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

func ObtenerTarifasAcademiaPorId(idAcademia int) ([]TarifaAcademia, error) {
	tarifaQuery := "call ingesoft.listarTarifasAcademia(?)"
	rows, err := database.DB.Query(tarifaQuery, idAcademia)
	if err != nil {
		logs.Logger.Println("Error al obtener las tarifas de la academia:", err)
		return nil, err
	}
	defer rows.Close()
	var tarifas []TarifaAcademia
	for rows.Next() {
		var tarifa TarifaAcademia
		if err := rows.Scan(&tarifa.ID, &tarifa.Periodo, &tarifa.UnidadFrecuenciaSem, &tarifa.CantidadFrecuencia, &tarifa.TipoSocio,
			&tarifa.Monto); err != nil {
			logs.Logger.Println("Error al escanear tarifa:", err)
			return nil, err
		}
		tarifas = append(tarifas, tarifa)
	}
	return tarifas, nil
}
