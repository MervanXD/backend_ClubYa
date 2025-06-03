package tipos

import (
	"fmt"
	"strings"

	"github.com/MervanXD/backend_ClubYa/logs"
)

type TipoSocio int

const (
	Socio TipoSocio = iota
	Externo
)

var tipoSocioStr = [...]string{
	"Socio",
	"Externo",
}

func (t TipoSocio) String() string {
	if int(t) < 0 || int(t) >= len(tipoSocioStr) {
		return "Desconocido"
	}
	return tipoSocioStr[t]
}

func (t TipoSocio) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}

func (t *TipoSocio) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), `"`)
	for i, nombre := range tipoSocioStr {
		if str == nombre {
			*t = TipoSocio(i)
			return nil
		}
	}
	*t = -1
	return nil
}

// Scan implementa la interfaz sql.Scanner para cada enum y así funcione al recibir de la BD.
func (t *TipoSocio) Scan(value interface{}) error {
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
		logs.Logger.Fatalf("tipo de dato no soportado para TipoSocio: %T", value)
		return fmt.Errorf("tipo de dato no soportado para TipoSocio: %T", value)
	}

	for i, nombre := range tipoSocioStr {
		if strValue == nombre {
			*t = TipoSocio(i)
			return nil
		}
	}
	return fmt.Errorf("TipoSocio desconocido: %s", strValue)
}
