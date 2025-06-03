package horario

import (
	"time"

	bloque "github.com/MervanXD/backend_ClubYa/internal/models/bloque_tiempo"
	"github.com/MervanXD/backend_ClubYa/internal/models/espacio"
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type HorarioDia struct {
	IdHorarioDia  int                   `json:"id_horario_dia"`
	Fecha         time.Time             `json:"fecha"`
	BloquesTiempo []bloque.BloqueTiempo `json:"bloques_tiempo"`
	Cancha        espacio.Cancha        `json:"cancha"`
	EspacioSocial espacio.EspacioSocial `json:"espacio_social"`
	Dia           tipos.Dia             `json:"dia"` // Representa el día de la semana (Lunes, Martes, etc.)
}
