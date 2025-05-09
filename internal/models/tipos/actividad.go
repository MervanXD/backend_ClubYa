package tipos

type Actividad int

const (
	SalaDeReuniones Actividad = iota
	Parrilla
	SalaDeFiestas
)

func (d Actividad) String() string {
	return [...]string{"Sala De Reuniones", "Parrilla", "Sala De Fiestas"}[d]
}

func (d Actividad) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *Actividad) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"Sala De Reuniones"`:
		*d = SalaDeReuniones
	case `"Parrilla"`:
		*d = Parrilla
	case `"Sala De Fiestas"`:
		*d = SalaDeFiestas
	default:
		*d = -1
	}
	return nil
}
