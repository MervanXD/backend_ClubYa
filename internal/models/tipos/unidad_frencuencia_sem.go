package tipos

import (
    "fmt"
    "strings"

    "github.com/MervanXD/backend_ClubYa/logs"
)

type UnidadFrecuenciaSem int

const (
    Sesiones UnidadFrecuenciaSem = iota
    Hora
)

var unidadFrecuenciaSemStr = [...]string{
    "Sesiones",
    "Hora",
}

func (u UnidadFrecuenciaSem) String() string {
    if int(u) < 0 || int(u) >= len(unidadFrecuenciaSemStr) {
        return "Desconocido"
    }
    return unidadFrecuenciaSemStr[u]
}

func (u UnidadFrecuenciaSem) MarshalJSON() ([]byte, error) {
    return []byte(`"` + u.String() + `"`), nil
}

func (u *UnidadFrecuenciaSem) UnmarshalJSON(data []byte) error {
    str := strings.Trim(string(data), `"`)
    for i, nombre := range unidadFrecuenciaSemStr {
        if str == nombre {
            *u = UnidadFrecuenciaSem(i)
            return nil
        }
    }
    *u = -1
    return nil
}

// Scan implementa la interfaz sql.Scanner para cada enum y así funcione al recibir de la BD.
func (u *UnidadFrecuenciaSem) Scan(value interface{}) error {
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
        logs.Logger.Fatalf("tipo de dato no soportado para UnidadFrecuenciaSem: %T", value)
        return fmt.Errorf("tipo de dato no soportado para UnidadFrecuenciaSem: %T", value)
    }

    for i, nombre := range unidadFrecuenciaSemStr {
        if strValue == nombre {
            *u = UnidadFrecuenciaSem(i)
            return nil
        }
    }
    return fmt.Errorf("UnidadFrecuenciaSem desconocido: %s", strValue)
}