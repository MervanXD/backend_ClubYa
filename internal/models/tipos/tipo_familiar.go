package tipos

import (
	"fmt"

	"github.com/MervanXD/backend_ClubYa/logs"
)

type TipoFamiliar int

const (
	Esposa TipoFamiliar = iota
	Esposo
	Papa
	Mama
	Abuela
	Abuelo
	Hijo
	Hija
)

func (d TipoFamiliar) String() string {
	return [...]string{"Esposa", "Esposo", "Papá", "Mamá", "Abuela", "Abuelo", "Hijo", "Hija"}[d]
}

func (d TipoFamiliar) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *TipoFamiliar) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"Esposa"`:
		*d = Esposa
	case `"Esposo"`:
		*d = Esposo
	case `"Papá"`:
		*d = Papa
	case `"Mamá"`:
		*d = Mama
	case `"Abuela"`:
		*d = Abuela
	case `"Abuelo"`:
		*d = Abuelo
	case `"Hijo"`:
		*d = Hijo
	case `"Hija"`:
		*d = Hija
	default:
		*d = -1
	}
	return nil
}

// Scan implementa la interfaz sql.Scanner para cada enum y asi funcione al momento de recibir de la bd.
func (uc *TipoFamiliar) Scan(value interface{}) error {
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
		logs.Logger.Fatalf("tipo de dato no soportado para TipoFamiliar: %T", value)
		return fmt.Errorf("tipo de dato no soportado para TipoFamiliar: %T", value)
	}

	switch strValue {
	case "Esposa":
		*uc = Esposa
	case "Esposo":
		*uc = Esposo
	case "Papá":
		*uc = Papa
	case "Mamá":
		*uc = Mama
	case "Abuela":
		*uc = Abuela
	case "Abuelo":
		*uc = Abuelo
	case "Hijo":
		*uc = Hijo
	case "Hija":
		*uc = Hija
	default:
		return fmt.Errorf("TipoFamiliar desconocido: %s", strValue)
	}
	return nil
}
