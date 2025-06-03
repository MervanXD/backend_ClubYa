package tipos

import (
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type Estado int

const (
	Confirmada Estado = iota
	Anulada
	Finalizada
	Pendiente
	Expirada
)

var estadoStr = [...]string{
	"Confirmada",
	"Anulada",
	"Finalizada",
	"Pendiente",
	"Expirada",
}

func (d Estado) String() string {
	if int(d) < 0 || int(d) >= len(estadoStr) {
		logs.Logger.Println("Error: Estado fuera de rango en String():", int(d))
		return "Desconocido"
	}
	return estadoStr[d]
}

func (d Estado) MarshalJSON() ([]byte, error) {
	if int(d) < 0 || int(d) >= len(estadoStr) {
		logs.Logger.Println("Error: Estado fuera de rango en MarshalJSON():", int(d))
	}
	return utils.EnumMarshalJSON(int(d), estadoStr[:])
}

func (d *Estado) UnmarshalJSON(data []byte) error {
	idx, err := utils.EnumUnmarshalJSON(data, estadoStr[:])
	*d = Estado(idx)
	if err != nil {
		logs.Logger.Println("Error unmarshaling Estado:", err)
		return err
	}
	return err
}

func (d *Estado) Scan(value interface{}) error {
	idx, err := utils.EnumScan(value, estadoStr[:])
	if err != nil {
		logs.Logger.Println("Error scanning Estado:", err)
		return err
	}
	*d = Estado(idx)
	return err
}
