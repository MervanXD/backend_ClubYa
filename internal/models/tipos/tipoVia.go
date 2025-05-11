package tipos

import (
	"fmt"

	"github.com/MervanXD/backend_ClubYa/logs"
)

type TipoVia int

const (
	Jr TipoVia = iota
	Av
	Calle
	Pasaje
	Plaza
	Mza
	Prolongacion
)

func (d TipoVia) String() string {
	return [...]string{"Jr.", "Av.", "Calle", "Pasaje", "Plaza", "Mza.", "Prolongacion"}[d]
}

func (d TipoVia) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *TipoVia) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"Jr."`:
		*d = Jr
	case `"Av."`:
		*d = Av
	case `"Calle"`:
		*d = Calle
	case `"Pasaje"`:
		*d = Pasaje
	case `"Plaza"`:
		*d = Plaza
	case `"Mza."`:
		*d = Mza
	case `"Prolongacion"`:
		*d = Prolongacion
	default:
		*d = -1
	}
	return nil
}

// Scan implementa la interfaz sql.Scanner para UbicacionCustom.
func (uc *TipoVia) Scan(value interface{}) error {
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
		logs.Logger.Fatalf("tipo de dato no soportado para tipo de via: %T", value)
		return fmt.Errorf("tipo de dato no soportado para tipo de via: %T", value)
	}

	// Aquí tu lógica para convertir strValue ("Puerta 1") a un int
	// Aquí tu lógica para convertir strValue ("Puerta 1") a un enum Ubicacion
	switch strValue {
	case "Jr.":
		*uc = Jr
	case "Av.":
		*uc = Av
	case "Calle":
		*uc = Calle
	case "Pasaje":
		*uc = Pasaje
	case "Plaza":
		*uc = Plaza
	case "Mza.":
		*uc = Mza
	case "Prolongacion":
		*uc = Prolongacion
	default:
		return fmt.Errorf("tipo de via desconocida: %s", strValue)
	}
	return nil
}
