package tipos

import (
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
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
		return "Desconocido"
	}
	return deportesStr[d]
}

func (d Deporte) MarshalJSON() ([]byte, error) {
	return utils.EnumMarshalJSON(int(d), deportesStr[:])
}

func (d *Deporte) UnmarshalJSON(data []byte) error {
	idx, err := utils.EnumUnmarshalJSON(data, deportesStr[:])
	*d = Deporte(idx)
	return err
}

func (d *Deporte) Scan(value interface{}) error {
	idx, err := utils.EnumScan(value, deportesStr[:])
	*d = Deporte(idx)
	return err
}
