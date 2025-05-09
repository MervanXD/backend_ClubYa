package tipos

type Deporte int

const (
	Tennis Deporte = iota
	Futbol
	Volley
	Basket
	Waterpolo
)

func (d Deporte) String() string {
	return [...]string{"Tennis", "Fútbol", "Volley", "Basket", "Waterpolo"}[d]
}

func (d Deporte) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *Deporte) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"Tennis"`:
		*d = Tennis
	case `"Fútbol"`:
		*d = Futbol
	case `"Volley"`:
		*d = Volley
	case `"Basket"`:
		*d = Basket
	case `"Waterpolo"`:
		*d = Waterpolo
	default:
		*d = -1
	}
	return nil
}
