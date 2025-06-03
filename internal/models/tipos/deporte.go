package tipos

import (
	"fmt"
	"strings"

	"github.com/MervanXD/backend_ClubYa/logs"
)

type Deporte int

const (
	Tenis Deporte = iota
	Fútbol
	Vóley
	Natación
	FullBody
	KungFuWushu
	Handball
	Aquaeróbicos
	AltoRendimiento
	Hockey
	Básquet
	FortalecimientoFlexibilidad
	TenisDeMesa
	Steps
	PreparaciónFísica
	Funcional
	Psicomotricidad
	Baile
	BodyBalancePilates
	Pilates
)

var deportesStr = [...]string{
	"Tenis",
	"Fútbol",
	"Vóley",
	"Natación",
	"Full Body",
	"Kung Fu Wushu",
	"Handball",
	"Aquaeróbicos",
	"Alto Rendimiento",
	"Hockey",
	"Básquet",
	"Fortalecimiento y Flexibilidad",
	"Tenis de Mesa",
	"Steps",
	"Preparación Física",
	"Funcional",
	"Psicomotricidad",
	"Baile",
	"Body Balance y Pilates",
	"Pilates",
}

func (d Deporte) String() string {
	if int(d) < 0 || int(d) >= len(deportesStr) {
		return "Desconocido"
	}
	return deportesStr[d]
}

func (d Deporte) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *Deporte) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), `"`)
	for i, nombre := range deportesStr {
		if str == nombre {
			*d = Deporte(i)
			return nil
		}
	}
	*d = -1
	return nil
}

// Scan implementa la interfaz sql.Scanner para cada enum y así funcione al recibir de la BD.
func (uc *Deporte) Scan(value interface{}) error {
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
		logs.Logger.Fatalf("tipo de dato no soportado para Deporte: %T", value)
		return fmt.Errorf("tipo de dato no soportado para Deporte: %T", value)
	}

	for i, nombre := range deportesStr {
		if strValue == nombre {
			*uc = Deporte(i)
			return nil
		}
	}
	return fmt.Errorf("Deporte desconocido: %s", strValue)
}
