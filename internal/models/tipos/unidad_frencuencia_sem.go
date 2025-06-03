package tipos

import (
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
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
		logs.Logger.Println("Error: UnidadFrecuenciaSem fuera de rango en String():", int(u))
		return "Desconocido"
	}
	return unidadFrecuenciaSemStr[u]
}

func (u UnidadFrecuenciaSem) MarshalJSON() ([]byte, error) {
	if int(u) < 0 || int(u) >= len(unidadFrecuenciaSemStr) {
		logs.Logger.Println("Error: UnidadFrecuenciaSem fuera de rango en MarshalJSON():", int(u))
	}
	return utils.EnumMarshalJSON(int(u), unidadFrecuenciaSemStr[:])
}

func (u *UnidadFrecuenciaSem) UnmarshalJSON(data []byte) error {
	idx, err := utils.EnumUnmarshalJSON(data, unidadFrecuenciaSemStr[:])
	*u = UnidadFrecuenciaSem(idx)
	if err != nil {
		logs.Logger.Println("Error unmarshaling UnidadFrecuenciaSem:", err)
		return err
	}
	return err
}

// Scan implementa la interfaz sql.Scanner para cada enum y así funcione al recibir de la BD.
func (u *UnidadFrecuenciaSem) Scan(value interface{}) error {
	idx, err := utils.EnumScan(value, unidadFrecuenciaSemStr[:])
	if err != nil {
		logs.Logger.Println("Error scanning UnidadFrecuenciaSem:", err)
		return err
	}
	*u = UnidadFrecuenciaSem(idx)
	return err
}
