package tipos

import (
	"fmt"

	"github.com/MervanXD/backend_ClubYa/logs"
)

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

// Scan implementa la interfaz sql.Scanner para cada enum y asi funcione al momento de recibir de la bd.
func (uc *Estado) Scan(value interface{}) error {
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
	case "Confirmada":
		*uc = Confirmada
	case "Anulada":
		*uc = Anulada
	case "Finalizada":
		*uc = Finalizada
	case "Pendiente":
		*uc = Pendiente
	case "Expirada":
		*uc = Expirada
	default:
		return fmt.Errorf("Actividad desconocida: %s", strValue)
	}
	return nil
}
