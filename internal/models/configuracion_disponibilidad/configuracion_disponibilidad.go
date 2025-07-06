package configuracion_disponibilidad

import "github.com/MervanXD/backend_ClubYa/internal/models/tipos"

type ConfiguracionDisponibilidad struct {
	IdEspacio      int       `json:"id_espacio"`
	IdBloqueTiempo int       `json:"id_bloque_tiempo"`
	Dia            tipos.Dia `json:"dia"`
}
