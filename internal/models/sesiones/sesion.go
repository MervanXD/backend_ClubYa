package sesiones

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type Sesion struct {
	IDsesion   int       `json:"id_sesion"`
	IdGrupo    int       `json:"id_grupo"`
	Dia        tipos.Dia `json:"dia"`
	HoraInicio string    `json:"hora_inicio"`
	HoraFin    string    `json:"hora_fin"`
}
