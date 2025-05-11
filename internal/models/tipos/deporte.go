package tipos

import (
	"fmt"

	"github.com/MervanXD/backend_ClubYa/logs"
)

type Deporte int

const (
	Tennis Deporte = iota
	Futbol
	Volley
	Basket
	Waterpolo
)

func (d Deporte) String() string {
	return [...]string{"Tennis", "Fútbol", "Volley", "Basket", "Waterpolo"}[d]
}

func (d Deporte) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *Deporte) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"Tennis"`:
		*d = Tennis
	case `"Fútbol"`:
		*d = Futbol
	case `"Volley"`:
		*d = Volley
	case `"Basket"`:
		*d = Basket
	case `"Waterpolo"`:
		*d = Waterpolo
	default:
		*d = -1
	}
	return nil
}

// Scan implementa la interfaz sql.Scanner para cada enum y asi funcione al momento de recibir de la bd.
func (uc *Deporte) Scan(value interface{}) error {
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
	case "Tennis":
		*uc = Tennis
	case "Futbol":
		*uc = Futbol
	case "Volley":
		*uc = Volley
	case "Basket":
		*uc = Basket
	case "Waterpolo":
		*uc = Waterpolo
	default:
		return fmt.Errorf("Actividad desconocida: %s", strValue)
	}
	return nil
}
