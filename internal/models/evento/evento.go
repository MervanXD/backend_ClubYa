package evento

import "time"

type Evento struct {
	IdEvento    int       `json:"id_evento"`
	Nombre      string    `json:"nombre"`
	Descripcion string    `json:"descripcion"`
	Fecha       time.Time `json:"fecha"`
	HoraInicio  time.Time `json:"hora_inicio"`
	HoraFin     time.Time `json:"hora_fin"`
	Aforo       int       `json:"aforo"`
	Invitados   int       `json:"invitados"`
	Precio      float64   `json:"precio"`
	Imagen      []byte    `json:"imagen"`
	Reglamento  []byte    `json:"reglamento"`
	NroInscritos int	  `json:"nro_inscritos"`
	Estado 		int		  `json:"estado"`
}
