package tipos

type Ubicacion int

const (
	Principal Ubicacion = iota
	Puerta1
	Puerta2
	Puerta3
)

func (d Ubicacion) String() string {
	return [...]string{"Principal", "Puerta 1", "Puerta 2", "Puerta 3"}[d]
}

func (d Ubicacion) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *Ubicacion) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"Principal"`:
		*d = Principal
	case `"Puerta 1"`:
		*d = Puerta1
	case `"Puerta 2"`:
		*d = Puerta2
	case `"Puerta 3"`:
		*d = Puerta3
	default:
		*d = -1
	}
	return nil
}
