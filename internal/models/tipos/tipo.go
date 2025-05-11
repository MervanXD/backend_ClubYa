package tipos

import (
	"fmt"

	"github.com/MervanXD/backend_ClubYa/logs"
)

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

// Scan implementa la interfaz sql.Scanner para cada enum y asi funcione al momento de recibir de la bd.
func (uc *Tipo) Scan(value interface{}) error {
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
		logs.Logger.Fatalf("tipo de dato no soportado para Actividad: %T", value)
		return fmt.Errorf("tipo de dato no soportado para Actividad: %T", value)
	}

	// Aquí tu lógica para convertir strValue ("Puerta 1") a un int
	// Aquí tu lógica para convertir strValue ("Puerta 1") a un enum Ubicacion
	switch strValue {
	case "Manual":
		*uc = Manual
	case "Automatico":
		*uc = Automatico
	default:
		return fmt.Errorf("Actividad desconocida: %s", strValue)
	}
	return nil
}
