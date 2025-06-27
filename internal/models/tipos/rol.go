package tipos

import (
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type Rol int

const (
	Titular Rol = iota
	Conyuge
	Administrador
	AdministradorAcademias
	AdminitradorMembresias
	AdministradorCanchas
	AdministradorEventos
)

var rolStr = [...]string{
	"Titular",
	"Conyuge",
	"Administrador",
	"Administrador_Academias",
	"Administrador_Membresias",
	"Administrador_Canchas",
	"Administrador_Eventos",
}

func (d Rol) String() string {
	if int(d) < 0 || int(d) >= len(rolStr) {
		logs.Logger.Println("Error: Rol fuera de rango en String():", int(d))
		return "Desconocido"
	}
	return rolStr[d]
}

func (d Rol) MarshalJSON() ([]byte, error) {
	if int(d) < 0 || int(d) >= len(rolStr) {
		logs.Logger.Println("Error: Rol fuera de rango en MarshalJSON():", int(d))
	}
	return utils.EnumMarshalJSON(int(d), rolStr[:])
}

func (d *Rol) UnmarshalJSON(data []byte) error {
	idx, err := utils.EnumUnmarshalJSON(data, rolStr[:])
	*d = Rol(idx)
	if err != nil {
		logs.Logger.Println("Error unmarshaling Rol:", err)
		return err
	}
	return err
}

func (d *Rol) Scan(value interface{}) error {
	idx, err := utils.EnumScan(value, rolStr[:])
	if err != nil {
		logs.Logger.Println("Error scanning Rol:", err)
		return err
	}
	*d = Rol(idx)
	return err
}
