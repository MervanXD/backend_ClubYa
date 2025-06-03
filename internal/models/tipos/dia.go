package tipos

import (
	"fmt"
	"strings"

	"github.com/MervanXD/backend_ClubYa/logs"
)

type Dia int

const (
	Lunes Dia = iota
	Martes
	Miércoles
	Jueves
	Viernes
	Sábado
	Domingo
)

var diasStr = [...]string{
	"Lunes",
	"Martes",
	"Miercoles",
	"Jueves",
	"Viernes",
	"Sábado",
	"Domingo",
}

func (d Dia) String() string {
	if int(d) < 0 || int(d) >= len(diasStr) {
		return "Desconocido"
	}
	return diasStr[d]
}

func (d Dia) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *Dia) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), `"`)
	for i, nombre := range diasStr {
		if str == nombre {
			*d = Dia(i)
			return nil
		}
	}
	*d = -1
	return nil
}

// Scan implementa la interfaz sql.Scanner para cada enum y así funcione al recibir de la BD.
func (d *Dia) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var strValue string
	switch v := value.(type) {
	case []byte:
		strValue = string(v)
	case string:
		strValue = v
	default:
		logs.Logger.Fatalf("tipo de dato no soportado para Dia: %T", value)
		return fmt.Errorf("tipo de dato no soportado para Dia: %T", value)
	}

	for i, nombre := range diasStr {
		if strValue == nombre {
			*d = Dia(i)
			return nil
		}
	}
	return fmt.Errorf("Dia desconocido: %s", strValue)
}
