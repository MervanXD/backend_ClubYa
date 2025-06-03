package tipos

import (
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type EstadoDisponibilidad int

const (
	NoDisponible EstadoDisponibilidad = iota
	Disponible
	Reservado
	Bloqueado
)

var estadoDisponibilidadStr = [...]string{
	"No disponible",
	"Disponible",
	"Reservado",
	"Bloqueado",
}

func (d EstadoDisponibilidad) String() string {
	if int(d) < 0 || int(d) >= len(estadoDisponibilidadStr) {
		logs.Logger.Println("Error: EstadoDisponibilidad fuera de rango en String():", int(d))
		return "Desconocido"
	}
	return estadoDisponibilidadStr[d]
}

func (d EstadoDisponibilidad) MarshalJSON() ([]byte, error) {
	if int(d) < 0 || int(d) >= len(estadoDisponibilidadStr) {
		logs.Logger.Println("Error: EstadoDisponibilidad fuera de rango en MarshalJSON():", int(d))
	}
	return utils.EnumMarshalJSON(int(d), estadoDisponibilidadStr[:])
}

func (d *EstadoDisponibilidad) UnmarshalJSON(data []byte) error {
	idx, err := utils.EnumUnmarshalJSON(data, estadoDisponibilidadStr[:])
	*d = EstadoDisponibilidad(idx)
	if err != nil {
		logs.Logger.Println("Error unmarshaling EstadoDisponibilidad:", err)
		return err
	}
	return err
}

func (d *EstadoDisponibilidad) Scan(value interface{}) error {
	idx, err := utils.EnumScan(value, estadoDisponibilidadStr[:])
	if err != nil {
		logs.Logger.Println("Error scanning EstadoDisponibilidad:", err)
		return err
	}
	*d = EstadoDisponibilidad(idx)
	return err
}
