package tipos

import (
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type TipoTrabajo int

const (
	Independiente TipoTrabajo = iota
	Empleado
)

var tipoTrabajoStr = [...]string{
	"Independiente",
	"Empleado",
}

func (d TipoTrabajo) String() string {
	if int(d) < 0 || int(d) >= len(tipoTrabajoStr) {
		logs.Logger.Println("Error: TipoTrabajo fuera de rango en String():", int(d))
		return "Desconocido"
	}
	return tipoTrabajoStr[d]
}

func (d TipoTrabajo) MarshalJSON() ([]byte, error) {
	if int(d) < 0 || int(d) >= len(tipoTrabajoStr) {
		logs.Logger.Println("Error: TipoTrabajo fuera de rango en MarshalJSON():", int(d))
	}
	return utils.EnumMarshalJSON(int(d), tipoTrabajoStr[:])
}

func (d *TipoTrabajo) UnmarshalJSON(data []byte) error {
	idx, err := utils.EnumUnmarshalJSON(data, tipoTrabajoStr[:])
	*d = TipoTrabajo(idx)
	if err != nil {
		logs.Logger.Println("Error unmarshaling TipoTrabajo:", err)
		return err
	}
	return err
}

func (d *TipoTrabajo) Scan(value interface{}) error {
	idx, err := utils.EnumScan(value, tipoTrabajoStr[:])
	if err != nil {
		logs.Logger.Println("Error scanning TipoTrabajo:", err)
		return err
	}
	*d = TipoTrabajo(idx)
	return err
}
