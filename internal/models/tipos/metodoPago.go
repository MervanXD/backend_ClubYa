package tipos

import (
	"fmt"

	"github.com/MervanXD/backend_ClubYa/logs"
)

type MetodoPago int

const (
	Tarjeta MetodoPago = iota
	Voucher
)

func (d MetodoPago) String() string {
	return [...]string{"Tarjeta", "Voucher"}[d]
}

func (d MetodoPago) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *MetodoPago) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"Tarjeta"`:
		*d = Tarjeta
	case `"Voucher"`:
		*d = Voucher
	default:
		*d = -1
	}
	return nil
}

// Scan implementa la interfaz sql.Scanner para cada enum y asi funcione al momento de recibir de la bd.
func (uc *MetodoPago) Scan(value interface{}) error {
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
		logs.Logger.Fatalf("tipo de dato no soportado para metodoPago: %T", value)
		return fmt.Errorf("tipo de dato no soportado para metodoPago: %T", value)
	}

	// Aquí tu lógica para convertir strValue ("Puerta 1") a un int
	// Aquí tu lógica para convertir strValue ("Puerta 1") a un enum Ubicacion
	switch strValue {
	case `"Tarjeta"`:
		*uc = Tarjeta
	case `"Voucher"`:
		*uc = Voucher
	default:
		return fmt.Errorf("metodo de pago desconocido: %s", strValue)
	}
	return nil
}
