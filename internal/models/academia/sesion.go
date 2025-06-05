package academia

import (
	"time"

	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type Sesion struct {
	IDsesion   int64         `json:"id_sesion"`
	Dia        tipos.Dia     `json:"dia"`
	HoraInicio time.Time     `json:"hora_inicio"`
	HoraFin    time.Time     `json:"hora_fin"`
	Grupos     GrupoAcademia `json:"grupo_academia"`
}
