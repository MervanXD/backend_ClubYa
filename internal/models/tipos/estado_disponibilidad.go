package tipos

import (
	"fmt"

	"github.com/MervanXD/backend_ClubYa/logs"
)

type EstadoDisponibilidad int

const (
	NoDisponible EstadoDisponibilidad = iota
	Disponible
	Reservado
	Bloqueado
)

func (d EstadoDisponibilidad) String() string {
	return [...]string{"No disponible", "Disponible", "Reservado", "Bloqueado"}[d]
}

func (d EstadoDisponibilidad) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *EstadoDisponibilidad) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"No disponible"`:
		*d = NoDisponible
	case `"Disponible"`:
		*d = Disponible
	case `"Reservado"`:
		*d = Reservado
	case `"Bloqueado"`:
		*d = Bloqueado
	default:
		*d = -1
	}
	return nil
}

// Scan implementa la interfaz sql.Scanner para cada enum y asi funcione al momento de recibir de la bd.
func (uc *EstadoDisponibilidad) Scan(value interface{}) error {
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
		logs.Logger.Fatalf("tipo de dato no soportado para Estado Disponiblidad: %T", value)
		return fmt.Errorf("tipo de dato no soportado para Estado Disponiblidad: %T", value)
	}

	switch strValue {
	case "No disponible":
		*uc = NoDisponible
	case "Disponible":
		*uc = Disponible
	case "Reservado":
		*uc = Reservado
	case "Bloqueado":
		*uc = Bloqueado
	default:
		return fmt.Errorf("Estado Disponiblidad desconocido: %s", strValue)
	}
	return nil
}
