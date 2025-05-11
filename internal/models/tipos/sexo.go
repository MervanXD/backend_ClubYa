package tipos

import (
	"fmt"

	"github.com/MervanXD/backend_ClubYa/logs"
)

type Sexo int

const (
	Femenino Sexo = iota
	Masculino
	PrefieroNoEspecificar
)

func (d Sexo) String() string {
	return [...]string{"Femenino", "Masculino", "Prefiero no especificar"}[d]
}

func (d Sexo) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *Sexo) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"Femenino"`:
		*d = Femenino
	case `"Masculino"`:
		*d = Masculino
	case `"Prefiero no especificar"`:
		*d = PrefieroNoEspecificar
	default:
		*d = -1
	}
	return nil
}

// Scan implementa la interfaz sql.Scanner para UbicacionCustom.
func (uc *Sexo) Scan(value interface{}) error {
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
	case "Femenino":
		*uc = Femenino
	case "Masculino":
		*uc = Masculino
	case "Prefiero no especificar":
		*uc = PrefieroNoEspecificar
	default:
		return fmt.Errorf("ubicación desconocida: %s", strValue)
	}
	return nil
}
