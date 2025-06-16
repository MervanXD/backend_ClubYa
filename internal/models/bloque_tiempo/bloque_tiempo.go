package bloque_tiempo

type BloqueTiempo struct {
	IdBloqueTiempo int    `json:"id_bloque_tiempo"`
	RangoInicio    string `json:"rango_inicio"`
	RangoFin       string `json:"rango_fin"`
}
