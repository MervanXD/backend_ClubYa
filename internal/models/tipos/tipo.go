package tipos

import (
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type Tipo int

const (
	Manual Tipo = iota
	Automatico
)

var tipoStr = [...]string{
	"Manual",
	"Automatico",
}

func (d Tipo) String() string {
	if int(d) < 0 || int(d) >= len(tipoStr) {
		logs.Logger.Println("Error: Tipo fuera de rango en String():", int(d))
		return "Desconocido"
	}
	return tipoStr[d]
}

func (d Tipo) MarshalJSON() ([]byte, error) {
	if int(d) < 0 || int(d) >= len(tipoStr) {
		logs.Logger.Println("Error: Tipo fuera de rango en MarshalJSON():", int(d))
	}
	return utils.EnumMarshalJSON(int(d), tipoStr[:])
}

func (d *Tipo) UnmarshalJSON(data []byte) error {
	idx, err := utils.EnumUnmarshalJSON(data, tipoStr[:])
	*d = Tipo(idx)
	if err != nil {
		logs.Logger.Println("Error unmarshaling Tipo:", err)
		return err
	}
	return err
}

func (d *Tipo) Scan(value interface{}) error {
	idx, err := utils.EnumScan(value, tipoStr[:])
	if err != nil {
		logs.Logger.Println("Error scanning Tipo:", err)
		return err
	}
	*d = Tipo(idx)
	return err
}
