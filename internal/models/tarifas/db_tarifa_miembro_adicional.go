package tarifas

import (
	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type tarifaMiembroAdicionalRepositoryDB struct{}

func NewTarifaMiembroAdicionalRespositoryDB() TarifaMiembroAdicionalRepository {
	return &tarifaMiembroAdicionalRepositoryDB{}
}

func (r *tarifaMiembroAdicionalRepositoryDB) ListarTarifaMiembroAdicional() ([]TarifaMiembroAdicional, error) {
	query := "CALL ListarTarifasMiembroAdicional()"
	rows, err := database.DB.Query(query)
	if err != nil {
		logs.Logger.Println("Error al listar las tarifas por miembro de familia:", err)
		return nil, err
	}
	defer rows.Close()

	var tarifas []TarifaMiembroAdicional
	for rows.Next() {
		var tarifa TarifaMiembroAdicional
		if err := rows.Scan(&tarifa.IdTarifaMiembroAdicional, &tarifa.FidTarifaMembresia, &tarifa.TipoMiembro,
			&tarifa.CostoAdicional, &tarifa.Estado); err != nil {
			logs.Logger.Println("Error al escanear la tarifa:", err)
			return nil, err
		}
		tarifas = append(tarifas, tarifa)
	}
	return tarifas, nil
}

func (r *tarifaMiembroAdicionalRepositoryDB) ModificarTarifaMiembroAdicional(idTarifa int, monto float64) error {
	query := "CALL ActualizarTarifaMiembroFamiliar(?, ?)"
	_, err := database.DB.Exec(query, monto, idTarifa)
	if err != nil {
		logs.Logger.Printf("Error al actualizar la tarifa del tipo de familiar con ID %d: %v\n", idTarifa, err)
		return err
	}
	return nil
}

func (r *tarifaMiembroAdicionalRepositoryDB) InsertarTarifaMiembroAdicional(tarifa TarifaMiembroAdicional) error {
	tx, err := database.DB.Begin()
	if err != nil {
		logs.Logger.Println("Error al iniciar la transacción:", err)
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	query := "CALL ingesoft.InsertarTarifaMiembroFamiliar(?, ?)"
	_, err = tx.Exec(query, tarifa.TipoMiembro.String(), tarifa.CostoAdicional)

	if err != nil {
		logs.Logger.Printf("Error al insertar tarifa (%s, %.2f): %v\n", tarifa.TipoMiembro, tarifa.CostoAdicional, err)
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		logs.Logger.Println("Error al hacer commit de la transacción:", err)
		return err
	}

	return nil
}
