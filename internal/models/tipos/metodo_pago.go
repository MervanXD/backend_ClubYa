package tipos

import (
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type MetodoPago int

const (
	Tarjeta MetodoPago = iota
	Voucher
)

var metodoPagoStr = [...]string{
	"Tarjeta",
	"Voucher",
}

func (d MetodoPago) String() string {
	if int(d) < 0 || int(d) >= len(metodoPagoStr) {
		logs.Logger.Println("Error: MetodoPago fuera de rango en String():", int(d))
		return "Desconocido"
	}
	return metodoPagoStr[d]
}

func (d MetodoPago) MarshalJSON() ([]byte, error) {
	if int(d) < 0 || int(d) >= len(metodoPagoStr) {
		logs.Logger.Println("Error: MetodoPago fuera de rango en MarshalJSON():", int(d))
	}
	return utils.EnumMarshalJSON(int(d), metodoPagoStr[:])
}

func (d *MetodoPago) UnmarshalJSON(data []byte) error {
	idx, err := utils.EnumUnmarshalJSON(data, metodoPagoStr[:])
	*d = MetodoPago(idx)
	if err != nil {
		logs.Logger.Println("Error unmarshaling MetodoPago:", err)
		return err
	}
	return err
}

// Scan implementa la interfaz sql.Scanner para cada enum y asi funcione al momento de recibir de la bd.
func (uc *MetodoPago) Scan(value interface{}) error {
	idx, err := utils.EnumScan(value, metodoPagoStr[:])
	if err != nil {
		logs.Logger.Println("Error scanning MetodoPago:", err)
		return err
	}
	*uc = MetodoPago(idx)
	return err
}
