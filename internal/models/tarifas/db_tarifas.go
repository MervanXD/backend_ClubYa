package tarifas

import (
	"database/sql"
	"strings"

	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type tarifaAcademiaRepositoryDB struct{}

func NewTarifaAcademiaRepositoryDB() TarifaAcademiaRepository {
	return &tarifaAcademiaRepositoryDB{}
}

func (r *tarifaAcademiaRepositoryDB) ObtenerTarifasAcademiaPorId(idGrupoAcademia int) ([]TarifaAcademia, error) {
	tarifaQuery := "call ingesoft.listarTarifasAcademia(?)"
	rows, err := database.DB.Query(tarifaQuery, idGrupoAcademia)
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

func (r *tarifaAcademiaRepositoryDB) InsertarTarifaAcademiaTx(tx *sql.Tx, tarifa *TarifaAcademia) (int64, error) {
	query := "call ingesoft.InsertarTarifaAcademia(?,?,?,?,?,?)"
	result, err := tx.Exec(query, tarifa.UnidadFrecuenciaSem.String(), tarifa.CantidadFrecuencia,
		tarifa.TipoSocio.String(), tarifa.Monto, true, tarifa.IDGrupo)
	if err != nil {
		logs.Logger.Println("Error al insertar la tarifa de la academia:", err)
		return 0, err
	}
	idTarifa, err := result.LastInsertId()
	if err != nil {
		logs.Logger.Println("Error al obtener el ID de la tarifa insertada:", err)
		return 0, err
	}
	return idTarifa, nil
}

func (r *tarifaAcademiaRepositoryDB) ActualizarTarifaAcademia(tarifa *TarifaAcademiaUpdate) error {
	setClauses := []string{}
	args := []interface{}{}

	if tarifa.Periodo != nil {
		setClauses = append(setClauses, "periodo = ?")
		args = append(args, *tarifa.Periodo)
	}
	if tarifa.UnidadFrecuenciaSem != nil {
		setClauses = append(setClauses, "unidadFrecuenciaSem = ?")
		args = append(args, *tarifa.UnidadFrecuenciaSem)
	}
	if tarifa.CantidadFrecuencia != nil {
		setClauses = append(setClauses, "cantidadFrecuencia = ?")
		args = append(args, *tarifa.CantidadFrecuencia)
	}
	if tarifa.TipoSocio != nil {
		setClauses = append(setClauses, "tipoSocio = ?")
		args = append(args, *tarifa.TipoSocio)
	}
	if tarifa.Monto != nil {
		setClauses = append(setClauses, "monto = ?")
		args = append(args, *tarifa.Monto)
	}
	if tarifa.EsActiva != nil {
		setClauses = append(setClauses, "esActiva = ?")
		args = append(args, *tarifa.EsActiva)
	}

	if len(setClauses) == 0 {
		return nil // No fields to update
	}

	query := "UPDATE TarifaAcademia SET " + strings.Join(setClauses, ", ") + " WHERE idTarifaAcademia = ?"
	args = append(args, tarifa.ID)

	if _, err := database.DB.Exec(query, args...); err != nil {
		logs.Logger.Println("Error al actualizar la tarifa de la academia:", err)
		return err
	}
	return nil
}
