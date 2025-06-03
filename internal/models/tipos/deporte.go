package tipos

import (
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type Deporte int

const (
	Tenis Deporte = iota
	Fútbol
	Vóley
	Natación
	FullBody
	KungFuWushu
	Handball
	Aquaeróbicos
	AltoRendimiento
	Hockey
	Básquet
	FortalecimientoFlexibilidad
	TenisDeMesa
	Steps
	PreparaciónFísica
	Funcional
	Psicomotricidad
	Baile
	BodyBalancePilates
	Pilates
)

var deportesStr = [...]string{
	"Tenis",
	"Fútbol",
	"Vóley",
	"Natación",
	"Full Body",
	"Kung Fu Wushu",
	"Handball",
	"Aquaeróbicos",
	"Alto Rendimiento",
	"Hockey",
	"Básquet",
	"Fortalecimiento y Flexibilidad",
	"Tenis de Mesa",
	"Steps",
	"Preparación Física",
	"Funcional",
	"Psicomotricidad",
	"Baile",
	"Body Balance y Pilates",
	"Pilates",
}

func (d Deporte) String() string {
	if int(d) < 0 || int(d) >= len(deportesStr) {
		logs.Logger.Println("Error: Deporte fuera de rango en String():", int(d))
		return "Desconocido"
	}
	return deportesStr[d]
}

func (d Deporte) MarshalJSON() ([]byte, error) {
	if int(d) < 0 || int(d) >= len(deportesStr) {
		logs.Logger.Println("Error: Deporte fuera de rango en MarshalJSON():", int(d))
	}
	return utils.EnumMarshalJSON(int(d), deportesStr[:])
}

func (d *Deporte) UnmarshalJSON(data []byte) error {
	idx, err := utils.EnumUnmarshalJSON(data, deportesStr[:])
	*d = Deporte(idx)
	if err != nil {
		logs.Logger.Println("Error unmarshaling Deporte:", err)
		return err
	}
	return err
}

func (d *Deporte) Scan(value interface{}) error {
	idx, err := utils.EnumScan(value, deportesStr[:])
	if err != nil {
		logs.Logger.Println("Error scanning Deporte:", err)
		return err
	}
	*d = Deporte(idx)
	return err
}
