package tipos

import (
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type MiembroFamiliar int

const (
	Cónyuge MiembroFamiliar = iota
	Padres
	Abuelos
	Hijos
	Hermanos
	Pareja
)

var miembroFamiliarStr = [...]string{
	"Cónyuge",
	"Padres",
	"Abuelos",
	"Hijos",
	"Hermanos",
	"Pareja",
}

func (d MiembroFamiliar) String() string {
	if int(d) < 0 || int(d) >= len(miembroFamiliarStr) {
		logs.Logger.Println("Error: MiembroFamiliar fuera de rango en String():", int(d))
		return "Desconocido"
	}
	return miembroFamiliarStr[d]
}

func (d MiembroFamiliar) MarshalJSON() ([]byte, error) {
	if int(d) < 0 || int(d) >= len(miembroFamiliarStr) {
		logs.Logger.Println("Error: MiembroFamiliar fuera de rango en MarshalJSON():", int(d))
	}
	return utils.EnumMarshalJSON(int(d), miembroFamiliarStr[:])
}

func (d *MiembroFamiliar) UnmarshalJSON(data []byte) error {
	idx, err := utils.EnumUnmarshalJSON(data, miembroFamiliarStr[:])
	*d = MiembroFamiliar(idx)
	if err != nil {
		logs.Logger.Println("Error unmarshaling MiembroFamiliar:", err)
		return err
	}
	return err
}

func (d *MiembroFamiliar) Scan(value interface{}) error {
	idx, err := utils.EnumScan(value, miembroFamiliarStr[:])
	if err != nil {
		logs.Logger.Println("Error scanning MiembroFamiliar:", err)
		return err
	}
	*d = MiembroFamiliar(idx)
	return err
}
