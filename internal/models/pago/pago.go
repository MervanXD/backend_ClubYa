package pago

import(
	"time"

	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
)

type Pago struct{
	IdPago		int 				`json:"id_pago"`
	fecha 		time.Time 			`json:"fecha"`
	concepto 	string 				`json:"concepto"`
	metodo		tipos.MetodoPago	`json:"metodo"`
	monto 		float64 			`json:"monto"`
}