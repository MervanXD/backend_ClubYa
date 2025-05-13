package tipos

import (
	"fmt"

	"github.com/MervanXD/backend_ClubYa/logs"
)

type EstadoSolicitud int

const (
	PendienteSolicitud EstadoSolicitud = iota
	Aceptada
	Rechazada
)

func (d EstadoSolicitud) String() string {
	return [...]string{"Pendiente", "Aceptada", "Rechazada"}[d]
}

func (d EstadoSolicitud) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *EstadoSolicitud) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"Pendiente"`:
		*d = PendienteSolicitud
	case `"Aceptada"`:
		*d = Aceptada
	case `"Rechazada"`:
		*d = Rechazada
	default:
		*d = -1
	}
	return nil
}

// Scan implementa la interfaz sql.Scanner para cada enum y asi funcione al momento de recibir de la bd.
func (uc *EstadoSolicitud) Scan(value interface{}) error {
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
		logs.Logger.Fatalf("tipo de dato no soportado para EstadoSolicitud: %T", value)
		return fmt.Errorf("tipo de dato no soportado para EstadoSolicitud: %T", value)
	}

	switch strValue {
	case "Pendiente":
		*uc = PendienteSolicitud
	case "Aceptada":
		*uc = Aceptada
	case "Rechazada":
		*uc = Rechazada
	default:
		return fmt.Errorf("EstadoSoliciutd desconocido: %s", strValue)
	}
	return nil
}
