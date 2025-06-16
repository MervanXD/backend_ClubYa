package horario

import "github.com/MervanXD/backend_ClubYa/internal/models/tipos"

type HorarioDiaDTO struct {
	IdHorarioDia int       `json:"id_horario_dia"`
	Fecha        string    `json:"fecha"`
	Dia          tipos.Dia `json:"dia"`       // Representa el día de la semana (Lunes, Martes, etc.)
	Inscritos    int       `json:"inscritos"` // Número de inscritos para el día
}
