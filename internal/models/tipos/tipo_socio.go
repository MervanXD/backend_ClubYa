package tipos

import (
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type TipoSocio int

const (
	Socio TipoSocio = iota
	Externo
)

var tipoSocioStr = [...]string{
	"Socio",
	"Externo",
}

func (t TipoSocio) String() string {
	if int(t) < 0 || int(t) >= len(tipoSocioStr) {
		logs.Logger.Println("Error: TipoSocio fuera de rango en String():", int(t))
		return "Desconocido"
	}
	return tipoSocioStr[t]
}

func (t TipoSocio) MarshalJSON() ([]byte, error) {
	if int(t) < 0 || int(t) >= len(tipoSocioStr) {
		logs.Logger.Println("Error: TipoSocio fuera de rango en MarshalJSON():", int(t))
	}
	return utils.EnumMarshalJSON(int(t), tipoSocioStr[:])
}

func (t *TipoSocio) UnmarshalJSON(data []byte) error {
	idx, err := utils.EnumUnmarshalJSON(data, tipoSocioStr[:])
	*t = TipoSocio(idx)
	if err != nil {
		logs.Logger.Println("Error unmarshaling TipoSocio:", err)
		return err
	}
	return err
}

// Scan implementa la interfaz sql.Scanner para cada enum y así funcione al recibir de la BD.
func (t *TipoSocio) Scan(value interface{}) error {
	idx, err := utils.EnumScan(value, tipoSocioStr[:])
	if err != nil {
		logs.Logger.Println("Error scanning TipoSocio:", err)
		return err
	}
	*t = TipoSocio(idx)
	return err
}
