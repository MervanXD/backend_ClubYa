package tipos

import (
	"fmt"

	"github.com/MervanXD/backend_ClubYa/logs"
)

type Ubicacion int

const (
	Principal Ubicacion = iota
	Puerta1
	Puerta2
	Puerta3
)

func (d Ubicacion) String() string {
	return [...]string{"Principal", "Puerta 1", "Puerta 2", "Puerta 3"}[d]
}

func (d Ubicacion) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *Ubicacion) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"Principal"`:
		*d = Principal
	case `"Puerta 1"`:
		*d = Puerta1
	case `"Puerta 2"`:
		*d = Puerta2
	case `"Puerta 3"`:
		*d = Puerta3
	default:
		*d = -1
	}
	return nil
}

// Scan implementa la interfaz sql.Scanner para UbicacionCustom.
func (uc *Ubicacion) Scan(value interface{}) error {
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
		logs.Logger.Fatalf("tipo de dato no soportado para UbicacionCustom: %T", value)
		return fmt.Errorf("tipo de dato no soportado para UbicacionCustom: %T", value)
	}

	// Aquí tu lógica para convertir strValue ("Puerta 1") a un int
	// Aquí tu lógica para convertir strValue ("Puerta 1") a un enum Ubicacion
	switch strValue {
	case "Puerta 1":
		*uc = Puerta1
	case "Puerta 2":
		*uc = Puerta2
	case "Puerta 3":
		*uc = Puerta3
	case "Principal":
		*uc = Principal
	default:
		return fmt.Errorf("ubicación desconocida: %s", strValue)
	}
	return nil
}
