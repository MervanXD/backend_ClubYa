package tipos

import (
	"fmt"

	"github.com/MervanXD/backend_ClubYa/logs"
)

type EstadoCuota int

const (
	PendienteCuota EstadoCuota = iota
	Pagado
	Vencido
)

func (d EstadoCuota) String() string {
	return [...]string{"Pagado", "Pendiente", "Vencido"}[d]
}

func (d EstadoCuota) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *EstadoCuota) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"Pagado"`:
		*d = Pagado
	case `"Pendiente"`:
		*d = PendienteCuota
	case `"Vencido"`:
		*d = Vencido
	default:
		*d = -1
	}
	return nil
}

// Scan implementa la interfaz sql.Scanner para cada enum y asi funcione al momento de recibir de la bd.
func (uc *EstadoCuota) Scan(value interface{}) error {
	if value == nil {
		// Manejar el caso de NULL de la base de datos si es necesario
		// *uc = 0 // o algún valor por defecto
		return nil
	}

	var strValue string
	switch v := value.(type) {
	case []byte:
		strValue = string(v)
	case string:
		strValue = v
	default:
		// Manejar el caso de tipo de dato no soportado
		logs.Logger.Fatalf("tipo de dato no soportado para EstadoCuota: %T", value)
		return fmt.Errorf("tipo de dato no soportado para EstadoCuota: %T", value)
	}

	switch strValue {
	case "Pagado":
		*uc = Pagado
	case "Pendiente":
		*uc = PendienteCuota
	case "Vencido":
		*uc = Vencido
	default:
		return fmt.Errorf("EstadoCuota desconocido: %s", strValue)
	}
	return nil
}