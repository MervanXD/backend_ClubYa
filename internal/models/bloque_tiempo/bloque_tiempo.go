package bloque_tiempo

import "time"

type BloqueTiempo struct {
	IdBloqueTiempo int       `json:"id_bloque_tiempo"`
	RangoInicio    time.Time `json:"rango_inicio"`
	RangoFin       time.Time `json:"rango_fin"`
}
