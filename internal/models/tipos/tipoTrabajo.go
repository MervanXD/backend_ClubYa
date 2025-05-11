package tipos

import (
	"fmt"

	"github.com/MervanXD/backend_ClubYa/logs"
)

type TipoTrabajo int

const (
	Independiente TipoTrabajo = iota
	Empleado
)

func (d TipoTrabajo) String() string {
	return [...]string{"Independiente", "Empleado"}[d]
}

func (d TipoTrabajo) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *TipoTrabajo) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"Independiente"`:
		*d = Independiente
	case `"Empleado"`:
		*d = Empleado
	default:
		*d = -1
	}
	return nil
}

// Scan implementa la interfaz sql.Scanner para UbicacionCustom.
func (uc *TipoTrabajo) Scan(value interface{}) error {
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
	case "Independiente":
		*uc = Independiente
	case "Empleado":
		*uc = Empleado
	default:
		return fmt.Errorf("ubicación desconocida: %s", strValue)
	}
	return nil
}
