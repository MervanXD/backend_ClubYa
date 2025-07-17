package tipos

import (
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type Ubicacion int

const (
	Principal Ubicacion = iota
	Puerta1
	Puerta2
	Puerta3
)

var ubicacionStr = [...]string{
	"Principal",
	"Puerta 1",
	"Puerta 2",
	"Puerta 3",
}

func (d Ubicacion) String() string {
	if int(d) < 0 || int(d) >= len(ubicacionStr) {
		logs.Logger.Println("Error: Ubicacion fuera de rango en String():", int(d))
		return "Desconocido"
	}
	return ubicacionStr[d]
}

func (d Ubicacion) MarshalJSON() ([]byte, error) {
	if int(d) < 0 || int(d) >= len(ubicacionStr) {
		logs.Logger.Println("Error: Ubicacion fuera de rango en MarshalJSON():", int(d))
	}
	return utils.EnumMarshalJSON(int(d), ubicacionStr[:])
}

func (d *Ubicacion) UnmarshalJSON(data []byte) error {
	idx, err := utils.EnumUnmarshalJSON(data, ubicacionStr[:])
	*d = Ubicacion(idx)
	if err != nil {
		logs.Logger.Println("Error unmarshaling Ubicacion:", err)
		return err
	}
	return err
}

func (d *Ubicacion) Scan(value interface{}) error {
	idx, err := utils.EnumScan(value, ubicacionStr[:])
	if err != nil {
		logs.Logger.Println("Error scanning Ubicacion:", err)
		return err
	}
	*d = Ubicacion(idx)
	return err
}
