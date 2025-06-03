package tipos

import (
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type TipoMembresia int

const (
	Regular TipoMembresia = iota
	Vitalicia
)

var tipoMembresiaStr = [...]string{
	"Regular",
	"Vitalicia",
}

func (d TipoMembresia) String() string {
	if int(d) < 0 || int(d) >= len(tipoMembresiaStr) {
		logs.Logger.Println("Error: TipoMembresia fuera de rango en String():", int(d))
		return "Desconocido"
	}
	return tipoMembresiaStr[d]
}

func (d TipoMembresia) MarshalJSON() ([]byte, error) {
	if int(d) < 0 || int(d) >= len(tipoMembresiaStr) {
		logs.Logger.Println("Error: TipoMembresia fuera de rango en MarshalJSON():", int(d))
	}
	return utils.EnumMarshalJSON(int(d), tipoMembresiaStr[:])
}

func (d *TipoMembresia) UnmarshalJSON(data []byte) error {
	idx, err := utils.EnumUnmarshalJSON(data, tipoMembresiaStr[:])
	*d = TipoMembresia(idx)
	if err != nil {
		logs.Logger.Println("Error unmarshaling TipoMembresia:", err)
		return err
	}
	return err
}

func (d *TipoMembresia) Scan(value interface{}) error {
	idx, err := utils.EnumScan(value, tipoMembresiaStr[:])
	if err != nil {
		logs.Logger.Println("Error scanning TipoMembresia:", err)
		return err
	}
	*d = TipoMembresia(idx)
	return err
}
