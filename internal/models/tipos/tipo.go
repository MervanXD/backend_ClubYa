package tipos

type Tipo int

const (
	Manual Tipo = iota
	Automatico
)

func (d Tipo) String() string {
	return [...]string{"Manual", "Automatico"}[d]
}

func (d Tipo) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *Tipo) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"Manual"`:
		*d = Manual
	case `"Automatico"`:
		*d = Automatico
	default:
		*d = -1
	}
	return nil
}
