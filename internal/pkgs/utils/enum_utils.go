package utils

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/MervanXD/backend_ClubYa/logs"
)

// Busca el índice de un string en un slice
func IndexOf(str string, arr []string) int {
	for i, v := range arr {
		if v == str {
			return i
		}
	}
	return -1
}

// Función genérica para MarshalJSON
func EnumMarshalJSON(e int, arr []string) ([]byte, error) {
	if e < 0 || e >= len(arr) {
		return json.Marshal("Desconocido")
	}
	return json.Marshal(arr[e])
}

// Función genérica para UnmarshalJSON
func EnumUnmarshalJSON(data []byte, arr []string) (int, error) {
	str := strings.Trim(string(data), `"`)
	idx := IndexOf(str, arr)
	if idx == -1 {
		logs.Logger.Printf("valor desconocido: %s", str)
		return -1, fmt.Errorf("valor desconocido: %s", str)
	}
	return idx, nil
}

// Función genérica para Scan
func EnumScan(value interface{}, arr []string) (int, error) {
	if value == nil {
		return -1, nil
	}
	var strValue string
	switch v := value.(type) {
	case []byte:
		strValue = string(v)
	case string:
		strValue = v
	default:
		logs.Logger.Printf("tipo de dato no soportado: %T", value)
		return -1, fmt.Errorf("tipo de dato no soportado: %T", value)
	}
	idx := IndexOf(strValue, arr)
	if idx == -1 {
		logs.Logger.Printf("valor desconocido: %s", strValue)
		return -1, fmt.Errorf("valor desconocido: %s", strValue)
	}
	return idx, nil
}
