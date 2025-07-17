package pago

import (
	"fmt"

	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

func InsertarPagoPorMembresia(idMembresia int, concepto, metodo string) (int, error) {
	var idPago int

	logs.Logger.Println("Ejecutando CALL CrearPagoDesdeCuota con:", "idMembresia=", idMembresia, "concepto=", concepto, "metodo=", metodo)

	// Llamar al procedimiento almacenado
	_, err := database.DB.Exec("CALL CrearPagoDesdeCuota(?, ?, ?, @p_idPago)", idMembresia, concepto, metodo)
	if err != nil {
		logs.Logger.Println("Error al ejecutar el procedimiento InsertarPagoPorMembresia:", err)
		return -1, err
	}

	logs.Logger.Println("CALL ejecutado correctamente, consultando @p_idPago...")

	// Obtener el ID del pago insertado
	err = database.DB.QueryRow("SELECT @p_idPago").Scan(&idPago)
	if err != nil {
		logs.Logger.Println("Error al obtener el ID del pago después del CALL:", err)
		return -1, err
	}

	logs.Logger.Println("ID del pago obtenido:", idPago)

	if idPago <= 0 {
		logs.Logger.Println("idPago inválido recibido:", idPago)
		return -1, err
	}

	return idPago, nil
}

func PagarTodosPagosPendientes(idTitular int, metodoPago string) ([]int, error) {
	var idsPagos []int

	logs.Logger.Printf("Ejecutando CALL PagarTodosPendientesPorTitular con: idTitular=%d, metodoPago=%s", idTitular, metodoPago)

	// Llamar al procedimiento almacenado que paga todos los pendientes
	rows, err := database.DB.Query("CALL PagarTodosPendientesPorTitular(?, ?)", idTitular, metodoPago)
	if err != nil {
		logs.Logger.Println("Error al ejecutar el procedimiento PagarTodosPendientesPorTitular:", err)
		return nil, err
	}
	defer rows.Close()

	// Recorrer los resultados para obtener todos los IDs de pagos procesados
	for rows.Next() {
		var idPago int
		if err := rows.Scan(&idPago); err != nil {
			logs.Logger.Println("Error al escanear ID de pago:", err)
			return nil, err
		}
		if idPago > 0 {
			idsPagos = append(idsPagos, idPago)
		}
	}

	if err := rows.Err(); err != nil {
		logs.Logger.Println("Error al procesar resultados:", err)
		return nil, err
	}

	logs.Logger.Printf("Pagos procesados exitosamente. IDs: %v", idsPagos)

	if len(idsPagos) == 0 {
		logs.Logger.Println("No se encontraron pagos pendientes para procesar")
		return nil, fmt.Errorf("no se encontraron pagos pendientes para el titular con ID %d", idTitular)
	}

	return idsPagos, nil
}
