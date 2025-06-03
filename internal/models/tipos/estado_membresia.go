package tipos

import (
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type EstadoMembresia int

const (
	Vigente EstadoMembresia = iota
	Suspendida
	Cancelada
)

var estadoMembresiaStr = [...]string{
	"Vigente",
	"Suspendida",
	"Cancelada",
}

func (d EstadoMembresia) String() string {
	if int(d) < 0 || int(d) >= len(estadoMembresiaStr) {
		logs.Logger.Println("Error: EstadoMembresia fuera de rango en String():", int(d))
		return "Desconocido"
	}
	return estadoMembresiaStr[d]
}

func (d EstadoMembresia) MarshalJSON() ([]byte, error) {
	if int(d) < 0 || int(d) >= len(estadoMembresiaStr) {
		logs.Logger.Println("Error: EstadoMembresia fuera de rango en MarshalJSON():", int(d))
	}
	return utils.EnumMarshalJSON(int(d), estadoMembresiaStr[:])
}

func (d *EstadoMembresia) UnmarshalJSON(data []byte) error {
	idx, err := utils.EnumUnmarshalJSON(data, estadoMembresiaStr[:])
	*d = EstadoMembresia(idx)
	if err != nil {
		logs.Logger.Println("Error unmarshaling EstadoMembresia:", err)
		return err
	}
	return err
}

func (d *EstadoMembresia) Scan(value interface{}) error {
	idx, err := utils.EnumScan(value, estadoMembresiaStr[:])
	if err != nil {
		logs.Logger.Println("Error scanning EstadoMembresia:", err)
		return err
	}
	*d = EstadoMembresia(idx)
	return err
}
