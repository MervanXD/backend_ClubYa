package tipos

import (
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type TipoVia int

const (
	Jr TipoVia = iota
	Av
	Calle
	Pasaje
	Plaza
	Mza
	Prolongacion
)

var tipoViaStr = [...]string{
	"Jr.",
	"Av.",
	"Calle",
	"Pasaje",
	"Plaza",
	"Mza.",
	"Prolongación",
}

func (d TipoVia) String() string {
	if int(d) < 0 || int(d) >= len(tipoViaStr) {
		logs.Logger.Println("Error: TipoVia fuera de rango en String():", int(d))
		return "Desconocido"
	}
	return tipoViaStr[d]
}

func (d TipoVia) MarshalJSON() ([]byte, error) {
	if int(d) < 0 || int(d) >= len(tipoViaStr) {
		logs.Logger.Println("Error: TipoVia fuera de rango en MarshalJSON():", int(d))
	}
	return utils.EnumMarshalJSON(int(d), tipoViaStr[:])
}

func (d *TipoVia) UnmarshalJSON(data []byte) error {
	idx, err := utils.EnumUnmarshalJSON(data, tipoViaStr[:])
	*d = TipoVia(idx)
	if err != nil {
		logs.Logger.Println("Error unmarshaling TipoVia:", err)
		return err
	}
	return err
}

func (d *TipoVia) Scan(value interface{}) error {
	idx, err := utils.EnumScan(value, tipoViaStr[:])
	if err != nil {
		logs.Logger.Println("Error scanning TipoVia:", err)
		return err
	}
	*d = TipoVia(idx)
	return err
}
