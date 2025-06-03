package tipos

import (
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type TipoDocumento int

const (
	DNI TipoDocumento = iota
	Pasaporte
	Cedula
)

var tipoDocumentoStr = [...]string{
	"DNI",
	"Pasaporte",
	"Cedula",
}

func (d TipoDocumento) String() string {
	if int(d) < 0 || int(d) >= len(tipoDocumentoStr) {
		logs.Logger.Println("Error: TipoDocumento fuera de rango en String():", int(d))
		return "Desconocido"
	}
	return tipoDocumentoStr[d]
}

func (d TipoDocumento) MarshalJSON() ([]byte, error) {
	if int(d) < 0 || int(d) >= len(tipoDocumentoStr) {
		logs.Logger.Println("Error: TipoDocumento fuera de rango en MarshalJSON():", int(d))
	}
	return utils.EnumMarshalJSON(int(d), tipoDocumentoStr[:])
}

func (d *TipoDocumento) UnmarshalJSON(data []byte) error {
	idx, err := utils.EnumUnmarshalJSON(data, tipoDocumentoStr[:])
	*d = TipoDocumento(idx)
	if err != nil {
		logs.Logger.Println("Error unmarshaling TipoDocumento:", err)
		return err
	}
	return err
}

func (d *TipoDocumento) Scan(value interface{}) error {
	idx, err := utils.EnumScan(value, tipoDocumentoStr[:])
	if err != nil {
		logs.Logger.Println("Error scanning TipoDocumento:", err)
		return err
	}
	*d = TipoDocumento(idx)
	return err
}
