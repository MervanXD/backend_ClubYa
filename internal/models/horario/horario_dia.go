package horario

import (
	"time"

	bloque "github.com/MervanXD/backend_ClubYa/internal/models/bloque_tiempo"
)

type HorarioDia struct {
	IdHorarioDia  int                   `json:"id_horario_dia"`
	Fecha         time.Time             `json:"fecha"`
	BloquesTiempo []bloque.BloqueTiempo `json:"bloques_tiempo"`
}
