package tarifas

import (
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

func (r *tarifaAcademiaRepositoryDB) InsertarTarifaAcademia(tarifa *TarifaAcademia) (int64, error) {
	query := "call ingesoft.InsertarTarifaAcademia(?,?,?,?,?,?,?)"
	result, err := database.DB.Exec(query, tarifa.UnidadFrecuenciaSem, tarifa.CantidadFrecuencia,
		tarifa.TipoSocio, tarifa.Monto, true, tarifa.IDGrupo)
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
