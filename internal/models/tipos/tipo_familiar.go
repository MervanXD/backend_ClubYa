package tipos

import (
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type TipoFamiliar int

const (
	Esposa TipoFamiliar = iota
	Esposo
	Papá
	Mamá
	Abuela
	Abuelo
	Hijo
	Hija
)

var tipoFamiliarStr = [...]string{
	"Esposa",
	"Esposo",
	"Papá",
	"Mamá",
	"Abuela",
	"Abuelo",
	"Hijo",
	"Hija",
}

func (d TipoFamiliar) String() string {
	if int(d) < 0 || int(d) >= len(tipoFamiliarStr) {
		logs.Logger.Println("Error: TipoFamiliar fuera de rango en String():", int(d))
		return "Desconocido"
	}
	return tipoFamiliarStr[d]
}

func (d TipoFamiliar) MarshalJSON() ([]byte, error) {
	if int(d) < 0 || int(d) >= len(tipoFamiliarStr) {
		logs.Logger.Println("Error: TipoFamiliar fuera de rango en MarshalJSON():", int(d))
	}
	return utils.EnumMarshalJSON(int(d), tipoFamiliarStr[:])
}

func (d *TipoFamiliar) UnmarshalJSON(data []byte) error {
	idx, err := utils.EnumUnmarshalJSON(data, tipoFamiliarStr[:])
	*d = TipoFamiliar(idx)
	if err != nil {
		logs.Logger.Println("Error unmarshaling TipoFamiliar:", err)
		return err
	}
	return err
}

func (d *TipoFamiliar) Scan(value interface{}) error {
	idx, err := utils.EnumScan(value, tipoFamiliarStr[:])
	if err != nil {
		logs.Logger.Println("Error scanning TipoFamiliar:", err)
		return err
	}
	*d = TipoFamiliar(idx)
	return err
}
