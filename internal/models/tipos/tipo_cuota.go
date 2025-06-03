package tipos

import (
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type EstadoCuota int

const (
	PendienteCuota EstadoCuota = iota
	Pagado
	Vencido
)

var estadoCuotaStr = [...]string{
	"Pendiente",
	"Pagado",
	"Vencido",
}

func (d EstadoCuota) String() string {
	if int(d) < 0 || int(d) >= len(estadoCuotaStr) {
		logs.Logger.Println("Error: EstadoCuota fuera de rango en String():", int(d))
		return "Desconocido"
	}
	return estadoCuotaStr[d]
}

func (d EstadoCuota) MarshalJSON() ([]byte, error) {
	if int(d) < 0 || int(d) >= len(estadoCuotaStr) {
		logs.Logger.Println("Error: EstadoCuota fuera de rango en MarshalJSON():", int(d))
	}
	return utils.EnumMarshalJSON(int(d), estadoCuotaStr[:])
}

func (d *EstadoCuota) UnmarshalJSON(data []byte) error {
	idx, err := utils.EnumUnmarshalJSON(data, estadoCuotaStr[:])
	*d = EstadoCuota(idx)
	if err != nil {
		logs.Logger.Println("Error unmarshaling EstadoCuota:", err)
		return err
	}
	return err
}

// Scan implementa la interfaz sql.Scanner para cada enum y asi funcione al momento de recibir de la bd.
func (d *EstadoCuota) Scan(value interface{}) error {
	idx, err := utils.EnumScan(value, estadoCuotaStr[:])
	if err != nil {
		logs.Logger.Println("Error scanning EstadoCuota:", err)
		return err
	}
	*d = EstadoCuota(idx)
	return err
}
