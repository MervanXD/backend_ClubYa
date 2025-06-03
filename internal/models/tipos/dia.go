package tipos

import (
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
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
	"Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo",
}

func (d Dia) String() string {
	if int(d) < 0 || int(d) >= len(diasStr) {
		return "Desconocido"
	}
	return diasStr[d]
}

func (d Dia) MarshalJSON() ([]byte, error) {
	return utils.EnumMarshalJSON(int(d), diasStr[:])
}

func (d *Dia) UnmarshalJSON(data []byte) error {
	idx, err := utils.EnumUnmarshalJSON(data, diasStr[:])
	*d = Dia(idx)
	if err != nil {
		logs.Logger.Println("Error unmarshaling Dia:", err)
		return err
	}
	return err
}

func (d *Dia) Scan(value interface{}) error {
	idx, err := utils.EnumScan(value, diasStr[:])
	if err != nil {
		logs.Logger.Println("Error scanning Dia:", err)
		return err
	}
	*d = Dia(idx)
	return err
}
