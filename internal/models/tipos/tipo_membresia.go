package tipos

import (
	"fmt"

	"github.com/MervanXD/backend_ClubYa/logs"
)

type TipoMembresia int

const (
	Regular TipoMembresia = iota
	Vitalicia
)

func (d TipoMembresia) String() string {
	return [...]string{"Regular", "Vitalicia"}[d]
}

func (d TipoMembresia) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *TipoMembresia) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"Regular"`:
		*d = Regular
	case `"Vitalicia"`:
		*d = Vitalicia
	default:
		*d = -1
	}
	return nil
}

// Scan implementa la interfaz sql.Scanner para cada enum y asi funcione al momento de recibir de la bd.
func (uc *TipoMembresia) Scan(value interface{}) error {
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
		logs.Logger.Fatalf("tipo de dato no soportado para TipoMembresia: %T", value)
		return fmt.Errorf("tipo de dato no soportado para TipoMembresia: %T", value)
	}

	// Aquí tu lógica para convertir strValue ("Puerta 1") a un int
	// Aquí tu lógica para convertir strValue ("Puerta 1") a un enum Ubicacion
	switch strValue {
	case "Regular":
		*uc = Regular
	case "Vitalicia":
		*uc = Vitalicia
	default:
		return fmt.Errorf("TipoMembresia desconocido: %s", strValue)
	}
	return nil
}
