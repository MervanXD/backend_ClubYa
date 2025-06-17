package horario

import (
	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type horarioDiaRepositoryDB struct{}

func NewHorarioDiaRepositoryDB() HorarioDiaRepository {
	return &horarioDiaRepositoryDB{}
}

func (r *horarioDiaRepositoryDB) ObtenerInscritosEspacioFecha(idEspacio int, fecha string) ([]HorarioDiaDTO, error) {
	query := "call ingesoft.ListarInscritosHorarioDia(?,?)"
	rows, err := database.DB.Query(query, idEspacio, fecha)
	if err != nil {
		logs.Logger.Println("Error al obtener los horarios dias del espacio : ", err)
		return nil, err
	}
	defer rows.Close()
	var horarioDia []HorarioDiaDTO
	for rows.Next() {
		var hd HorarioDiaDTO
		if err := rows.Scan(&hd.IdHorarioDia, &hd.Fecha, &hd.Dia, &hd.Inscritos); err != nil {
			logs.Logger.Println("Error al escanear el bloque: ", err)
			return nil, err
		}
		horarioDia = append(horarioDia, hd)
	}
	return horarioDia, nil
}
