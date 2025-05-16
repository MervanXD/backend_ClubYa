package cuota

import(
	"time"

	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
) 

type Cuota struct{
	IdCuota				int 				`json:"idCuota"`
	Periodo				int					`json:"periodo"`
	FechaEmision		time.Time			`json:"fechaEmision"`
	FechaVencimiento	time.Time			`json:"fechaVencimiento"`
	MontoTotal			float64				`json:"montoTotal"`
	Estado				tipos.EstadoCuota	`json:"estado_cuota"`
}