package tipos

type Estado int

const (
	Confirmada Estado = iota
	Anulada
	Finalizada
	Pendiente
	Expirada
)

func (d Estado) String() string {
	return [...]string{"Confirmada", "Anulada", "Finalizada", "Pendiente", "Expirada"}[d]
}

func (d Estado) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *Estado) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"Confirmada"`:
		*d = Confirmada
	case `"Anulada"`:
		*d = Anulada
	case `"Finalizada"`:
		*d = Finalizada
	case `"Pendiente"`:
		*d = Pendiente
	case `"Expirada"`:
		*d = Expirada
	default:
		*d = -1
	}
	return nil
}
