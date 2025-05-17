package pago

import (
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

