package tipos

import (
	"fmt"

	"github.com/MervanXD/backend_ClubYa/logs"
)

type TipoDocumento int

const (
	Dni TipoDocumento = iota
	ActaMatrimonio
	ActaNacimiento
)

func (d TipoDocumento) String() string {
	return [...]string{"DNI", "Acta_Matrimonio", "Acta_Nacimiento"}[d]
}

func (d TipoDocumento) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *TipoDocumento) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"DNI"`:
		*d = Dni
	case `"Acta_Matrimonio"`:
		*d = ActaMatrimonio
	case `"Acta_Nacimiento"`:
		*d = ActaNacimiento
	default:
		*d = -1
	}
	return nil
}

// Scan implementa la interfaz sql.Scanner para UbicacionCustom.
func (uc *TipoDocumento) Scan(value interface{}) error {
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
		logs.Logger.Fatalf("tipo de dato no soportado para Sexo: %T", value)
		return fmt.Errorf("tipo de dato no soportado para Sexo: %T", value)
	}

	// Aquí tu lógica para convertir strValue ("Puerta 1") a un int
	// Aquí tu lógica para convertir strValue ("Puerta 1") a un enum Ubicacion
	switch strValue {
	case "DNI":
		*uc = Dni
	case "Acta_Matrimonio":
		*uc = ActaMatrimonio
	case "Acta_Nacimiento":
		*uc = ActaNacimiento
	default:
		return fmt.Errorf("ubicación desconocida: %s", strValue)
	}
	return nil
}
