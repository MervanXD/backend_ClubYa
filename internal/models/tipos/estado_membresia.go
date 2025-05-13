package tipos

import (
	"fmt"

	"github.com/MervanXD/backend_ClubYa/logs"
)

type EstadoMembresia int

const (
	Vigente EstadoMembresia = iota
	Suspendida
	Cancelada
)

func (d EstadoMembresia) String() string {
	return [...]string{"Vigente", "Suspendida", "Cancelada"}[d]
}

func (d EstadoMembresia) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *EstadoMembresia) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"Vigente"`:
		*d = Vigente
	case `"Suspendida"`:
		*d = Suspendida
	case `"Cancelada"`:
		*d = Cancelada
	default:
		*d = -1
	}
	return nil
}

// Scan implementa la interfaz sql.Scanner para cada enum y asi funcione al momento de recibir de la bd.
func (uc *EstadoMembresia) Scan(value interface{}) error {
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
		logs.Logger.Fatalf("tipo de dato no soportado para EstadoMembresia: %T", value)
		return fmt.Errorf("tipo de dato no soportado para EstadoMembresia: %T", value)
	}

	switch strValue {
	case "Vigente":
		*uc = Vigente
	case "Suspendida":
		*uc = Suspendida
	case "Cancelada":
		*uc = Cancelada
	default:
		return fmt.Errorf("EstadoMembresia desconocido: %s", strValue)
	}
	return nil
}
