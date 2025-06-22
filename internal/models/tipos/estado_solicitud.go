package tipos

import (
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type EstadoSolicitud int

const (
	PendienteSolicitud EstadoSolicitud = iota
	Aceptada
	Rechazada
	Pagada
)

var estadoSolicitudStr = [...]string{
	"Pendiente",
	"Aceptada",
	"Rechazada",
	"Pagada",
}

func (d EstadoSolicitud) String() string {
	if int(d) < 0 || int(d) >= len(estadoSolicitudStr) {
		logs.Logger.Println("Error: EstadoSolicitud fuera de rango en String():", int(d))
		return "Desconocido"
	}
	return estadoSolicitudStr[d]
}

func (d EstadoSolicitud) MarshalJSON() ([]byte, error) {
	if int(d) < 0 || int(d) >= len(estadoSolicitudStr) {
		logs.Logger.Println("Error: EstadoSolicitud fuera de rango en MarshalJSON():", int(d))
	}
	return utils.EnumMarshalJSON(int(d), estadoSolicitudStr[:])
}

func (d *EstadoSolicitud) UnmarshalJSON(data []byte) error {
	idx, err := utils.EnumUnmarshalJSON(data, estadoSolicitudStr[:])
	*d = EstadoSolicitud(idx)
	if err != nil {
		logs.Logger.Println("Error unmarshaling EstadoSolicitud:", err)
		return err
	}
	return err
}

func (d *EstadoSolicitud) Scan(value interface{}) error {
	idx, err := utils.EnumScan(value, estadoSolicitudStr[:])
	if err != nil {
		logs.Logger.Println("Error scanning EstadoSolicitud:", err)
		return err
	}
	*d = EstadoSolicitud(idx)
	return err
}
