package persona

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type Persona struct {
	Id              int           `json:"id"`
	Nombre          string        `json:"nombre"`
	Apellidos       string        `json:"apellidos"`
	Sexo            tipos.Sexo    `json:"sexo"`
	Dni             string        `json:"dni"`
	FechaNacimiento string        `json:"fecha_nacimiento"`
	Telefono        string        `json:"telefono"`
	Pais            string        `json:"pais"`
	Provincia       string        `json:"provincia"`
	Distrito        string        `json:"distrito"`
	Direccion       string        `json:"direccion"`
	TipoVia         tipos.TipoVia `json:"tipo_via"`
	Referencia      string        `json:"referencia"`
	Ciudad          string        `json:"ciudad"`
	CodigoPostal    string        `json:"codigo_postal"`
}
