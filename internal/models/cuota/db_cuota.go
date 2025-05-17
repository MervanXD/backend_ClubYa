package cuota

import (
	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

func ObtenerCuotasPorMembresiaLimitado(idMembresia int, limite int) ([]Cuota, error) {
	query := "CALL ObtenerCuotasPorMembresia(?, ?)"
	rows, err := database.DB.Query(query, idMembresia, limite)
	if err != nil {
		logs.Logger.Println("Error al obtener cuotas: ", err)
		return nil, err
	}
	defer rows.Close()

	var cuotas []Cuota
	for rows.Next() {
		var c Cuota
		err := rows.Scan(&c.IdCuota, &c.Periodo, &c.MontoTotal, &c.FechaEmision, &c.FechaVencimiento, &c.Estado)
		if err != nil {
			logs.Logger.Println("Error al escanear cuota:", err)
			return nil, err
		}
		cuotas = append(cuotas, c)
	}

	return cuotas, nil
}
