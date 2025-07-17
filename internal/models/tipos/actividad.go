package tipos

import (
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type Actividad int

const (
	SalaDeReuniones Actividad = iota
	Parrilla
	SalaDeFiestas
)

var actividadStr = [...]string{
	"Sala de reuniones",
	"Parrilla",
	"Sala de fiestas",
}

func (a Actividad) String() string {
	if int(a) < 0 || int(a) >= len(actividadStr) {
		logs.Logger.Println("Error: Actividad fuera de rango en String():", int(a))
		return "Desconocido"
	}
	return actividadStr[a]
}

func (a Actividad) MarshalJSON() ([]byte, error) {
	if int(a) < 0 || int(a) >= len(actividadStr) {
		logs.Logger.Println("Error: Actividad fuera de rango en MarshalJSON():", int(a))
	}
	return utils.EnumMarshalJSON(int(a), actividadStr[:])
}

func (a *Actividad) UnmarshalJSON(data []byte) error {
	idx, err := utils.EnumUnmarshalJSON(data, actividadStr[:])
	*a = Actividad(idx)
	if err != nil {
		logs.Logger.Println("Error unmarshaling Actividad:", err)
		return err
	}
	return err
}

func (a *Actividad) Scan(value interface{}) error {
	idx, err := utils.EnumScan(value, actividadStr[:])
	if err != nil {
		logs.Logger.Println("Error scanning Actividad:", err)
		return err
	}
	*a = Actividad(idx)
	return err
}
