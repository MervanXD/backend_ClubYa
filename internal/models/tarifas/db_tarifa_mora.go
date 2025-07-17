package tarifas

import (
	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type tarifaMoraRepositoryDB struct{}

func NewTarifaMoraRespositoryDB() TarifaMoraRepository {
	return &tarifaMoraRepositoryDB{}
}

func (r *tarifaMoraRepositoryDB) InsertarTarifaMora(tarifa TarifaMora) error {
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

	query := "CALL ingesoft.InsertarTarifaMora(?, ?, ?)"
	_, err = tx.Exec(query, tarifa.NombreMora, tarifa.CostoMora, tarifa.FrecuenciaDias)

	if err != nil {
		logs.Logger.Printf("Error al insertar tarifa mora (nombre: %s, costo: %.2f, frecuencia: %d): %v\n",
			tarifa.NombreMora, tarifa.CostoMora, tarifa.FrecuenciaDias, err)
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		logs.Logger.Println("Error al hacer commit de la transacción:", err)
		return err
	}

	return nil
}

func (r *tarifaMoraRepositoryDB) ObtenerTarifaMoraPorMembresia(id int64) (TarifaMora, error) {
	query := "CALL ObtenerTarifaMoraPorMembresia(?)"
	row := database.DB.QueryRow(query, id)
	var tarifa TarifaMora
	if err := row.Scan(&tarifa.IdTarifaMora, &tarifa.FidTarifaMembresia, &tarifa.NombreMora,
		&tarifa.CostoMora, &tarifa.Estado, &tarifa.FrecuenciaDias); err != nil {
		logs.Logger.Println("Error al escanear la mora:", err)
		return TarifaMora{}, err
	}
	return tarifa, nil
}

func (r *tarifaMoraRepositoryDB) ListarTarifaMora() ([]TarifaMora, error) {
	query := "CALL ListarTarifasMora()"
	rows, err := database.DB.Query(query)
	if err != nil {
		logs.Logger.Println("Error al listar las moras: ", err)
		return nil, err
	}
	defer rows.Close()

	var moras []TarifaMora
	for rows.Next() {
		var mora TarifaMora
		if err := rows.Scan(&mora.IdTarifaMora, &mora.FidTarifaMembresia, &mora.NombreMora,
			&mora.CostoMora, &mora.Estado, &mora.FrecuenciaDias); err != nil {
			logs.Logger.Println("Error al escanear la mora:", err)
			return nil, err
		}
		moras = append(moras, mora)
	}
	return moras, nil
}

func (r *tarifaMoraRepositoryDB) ModificarTarifaMora(tarifa TarifaMora) error {
	query := "CALL ActualizarTarifaMora(?, ?,?,?)"
	_, err := database.DB.Exec(query, tarifa.IdTarifaMora, tarifa.NombreMora, tarifa.CostoMora, tarifa.FrecuenciaDias)
	if err != nil {
		logs.Logger.Printf("Error al actualizar la mora con ID %d: %v\n", tarifa.IdTarifaMora, err)
		return err
	}
	return nil
}
