package tipos

import (
	"fmt"

	"github.com/MervanXD/backend_ClubYa/logs"
)

type Rol int

const (
	Titular Rol = iota
	Conyuge
	Administrador
	AdministradorActividad
	AdminitradorMembresias
	AdministradorCanchas
	AdministradorEventos
)

func (d Rol) String() string {
	return [...]string{"Titular", "Conyuge", "Administrador", "Administrador_Academias", "Administrador_Membresias", "Administrador_Canchas", "Administrador_Eventos"}[d]
}

func (d Rol) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *Rol) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"Titular"`:
		*d = Titular
	case `"Conyuge"`:
		*d = Conyuge
	case `"Administrador"`:
		*d = Administrador
	case `"Administrador_Actividad"`:
		*d = AdministradorActividad
	case `"Adminitrador_Membresias"`:
		*d = AdminitradorMembresias
	case `"Administrador_Canchas"`:
		*d = AdministradorCanchas
	case `"Administrador_Eventos"`:
		*d = AdministradorEventos
	default:
		*d = -1
	}
	return nil
}

// Scan implementa la interfaz sql.Scanner para UbicacionCustom.
func (uc *Rol) Scan(value interface{}) error {
	if value == nil {
		// Manejar el caso de NULL de la base de datos si es necesario
		// *uc = 0 // o algún valor por defecto
		return nil
	}

	var strValue string
	switch v := value.(type) {
	case []byte:
		strValue = string(v)
	case string:
		strValue = v
	default:
		// Manejar el caso de tipo de dato no soportado
		logs.Logger.Fatalf("tipo de dato no soportado para rol: %T", value)
		return fmt.Errorf("tipo de dato no soportado para rol: %T", value)
	}

	// Aquí tu lógica para convertir strValue ("Puerta 1") a un int
	// Aquí tu lógica para convertir strValue ("Puerta 1") a un enum Ubicacion
	switch strValue {
	case "Titular":
		*uc = Titular
	case "conyuge":
		*uc = Conyuge
	case "Administrador":
		*uc = Administrador
	case "Administrador_Actividad":
		*uc = AdministradorActividad
	case "Adminitrador_Membresias":
		*uc = AdminitradorMembresias
	case "Administrador_Canchas":
		*uc = AdministradorCanchas
	case "Administrador_Eventos":
		*uc = AdministradorEventos
	default:
		return fmt.Errorf("tipo de rol desconocido: %s", strValue)
	}
	return nil
}
