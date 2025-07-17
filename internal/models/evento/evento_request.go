package evento

type EventoRequest struct {
	IdEvento     int     `json:"id_evento"`
	Nombre       string  `json:"nombre"`
	Descripcion  string  `json:"descripcion"`
	Fecha        string  `json:"fecha"`
	Aforo        int     `json:"aforo"`
	Invitados    int     `json:"invitados"`
	Precio       float64 `json:"precio"`
	HoraInicio   string  `json:"hora_inicio"`
	HoraFin      string  `json:"hora_fin"`
	NroInscritos int     `json:"nro_inscritos"`
	IdEspacio    int     `json:"id_espacio"`
	Estado       int     `json:"estado"`
	Imagen       string  `json:"imagen"`
}
