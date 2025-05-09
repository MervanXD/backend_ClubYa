package tipos

import (
	"fmt"

	"github.com/MervanXD/backend_ClubYa/logs"
)

type Actividad int

const (
	SalaDeReuniones Actividad = iota
	Parrilla
	SalaDeFiestas
)

func (d Actividad) String() string {
	return [...]string{"Sala de reuniones", "Parrilla", "Sala de fiestas"}[d]
}

func (d Actividad) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *Actividad) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"Sala de reuniones"`:
		*d = SalaDeReuniones
	case `"Parrilla"`:
		*d = Parrilla
	case `"Sala de fiestas"`:
		*d = SalaDeFiestas
	default:
		*d = -1
	}
	return nil
}

func (d Actividad) FromString(str string) (Actividad, error) {
	switch str {
	case "Sala de reuniones":
		return SalaDeReuniones, nil
	case "Parrilla":
		return Parrilla, nil
	case "Sala de fiestas":
		return SalaDeFiestas, nil
	default:
		return -1, fmt.Errorf("actividad desconocida: %s", str)
	}
}

// Scan implementa la interfaz sql.Scanner para UbicacionCustom.
func (uc *Actividad) Scan(value interface{}) error {
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
	case "Sala de reuniones":
		*uc = SalaDeReuniones
	case "Parrilla":
		*uc = Parrilla
	case "Sala de fiestas":
		*uc = SalaDeFiestas
	default:
		return fmt.Errorf("Actividad desconocida: %s", strValue)
	}
	return nil
}
