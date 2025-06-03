package tipos

import (
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type Sexo int

const (
	Femenino Sexo = iota
	Masculino
	PrefieroNoEspecificar
)

var sexoStr = [...]string{
	"Femenino",
	"Masculino",
	"Prefiero no especificar",
}

func (d Sexo) String() string {
	if int(d) < 0 || int(d) >= len(sexoStr) {
		logs.Logger.Println("Error: Sexo fuera de rango en String():", int(d))
		return "Desconocido"
	}
	return sexoStr[d]
}

func (d Sexo) MarshalJSON() ([]byte, error) {
	if int(d) < 0 || int(d) >= len(sexoStr) {
		logs.Logger.Println("Error: Sexo fuera de rango en MarshalJSON():", int(d))
	}
	return utils.EnumMarshalJSON(int(d), sexoStr[:])
}

func (d *Sexo) UnmarshalJSON(data []byte) error {
	idx, err := utils.EnumUnmarshalJSON(data, sexoStr[:])
	*d = Sexo(idx)
	if err != nil {
		logs.Logger.Println("Error unmarshaling Sexo:", err)
		return err
	}
	return err
}

func (d *Sexo) Scan(value interface{}) error {
	idx, err := utils.EnumScan(value, sexoStr[:])
	if err != nil {
		logs.Logger.Println("Error scanning Sexo:", err)
		return err
	}
	*d = Sexo(idx)
	return err
}
