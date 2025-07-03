package tipos

import (
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type TipoDocumentoIdentidad int

const (
	DocumentoNacionalIdentidad TipoDocumentoIdentidad = iota
	RUC
	CarnetExtranjeria
	Pasaporte
)

var tiposDocIdentidadStr = [...]string{
	"DNI", "RUC", "Carnet de Extranjeria", "Pasaporte",
}

func (d TipoDocumentoIdentidad) String() string {
	if int(d) < 0 || int(d) >= len(tiposDocIdentidadStr) {
		return "Desconocido"
	}
	return diasStr[d]
}

func (d TipoDocumentoIdentidad) MarshalJSON() ([]byte, error) {
	return utils.EnumMarshalJSON(int(d), tiposDocIdentidadStr[:])
}

func (d *TipoDocumentoIdentidad) UnmarshalJSON(data []byte) error {
	idx, err := utils.EnumUnmarshalJSON(data, tiposDocIdentidadStr[:])
	*d = TipoDocumentoIdentidad(idx)
	if err != nil {
		logs.Logger.Println("Error unmarshaling Dia:", err)
		return err
	}
	return err
}

func (d *TipoDocumentoIdentidad) Scan(value interface{}) error {
	idx, err := utils.EnumScan(value, tiposDocIdentidadStr[:])
	if err != nil {
		logs.Logger.Println("Error scanning Dia:", err)
		return err
	}
	*d = TipoDocumentoIdentidad(idx)
	return err
}
