package bloque_tiempo

import (
	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type bloqueTiempoRespositoryDB struct{}

func NewBloqueTiempoRepositoryDB() BloqueTiempoRepository {
	return &bloqueTiempoRespositoryDB{}
}

func (r *bloqueTiempoRespositoryDB) ObtenerBloquesTiempoEspacio(idEspacio int) ([]BloqueTiempo, error) {
	query := "call ingesoft.ListarBloquesTiempo(?)"
	rows, err := database.DB.Query(query, idEspacio)
	if err != nil {
		logs.Logger.Println("Error al obtener los bloques de tiempo: ", err)
		return nil, err
	}
	defer rows.Close()
	var bloques []BloqueTiempo
	for rows.Next() {
		var bloque BloqueTiempo
		if err := rows.Scan(&bloque.IdBloqueTiempo, &bloque.RangoInicio, &bloque.RangoFin); err != nil {
			logs.Logger.Println("Error al escanear el bloque: ", err)
			return nil, err
		}
		bloques = append(bloques, bloque)
	}
	return bloques, nil
}

func (r *bloqueTiempoRespositoryDB) ListarBloquesTiempoEstandar() ([]BloqueTiempo, error) {
	query := "call ingesoft.ListarBloquesTiempoEstandar()"
	rows, err := database.DB.Query(query)
	if err != nil {
		logs.Logger.Println("Error al obtener los bloques de tiempo estandar: ", err)
		return nil, err
	}
	defer rows.Close()
	var bloques []BloqueTiempo
	for rows.Next() {
		var bloque BloqueTiempo
		if err := rows.Scan(&bloque.IdBloqueTiempo, &bloque.RangoInicio, &bloque.RangoFin); err != nil {
			logs.Logger.Println("Error al escanear el bloque estandar: ", err)
			return nil, err
		}
		bloques = append(bloques, bloque)
	}
	return bloques, nil
}
