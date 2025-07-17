package reserva

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/tipos"
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
)

type ReservaRequest struct {
	NombreSocio      string            `json:"nombre_socio"`
	IdReserva        int               `json:"id_reserva"`
	FechaReserva     utils.NullString  `json:"fecha_reserva"`   /*Fecha en la que se realizo la reserva*/
	Fecha            utils.NullString  `json:"fecha_reservada"` /*Fecha de la reserva*/
	HoraInicio       utils.NullString  `json:"hora_inicio"`
	HoraFin          utils.NullString  `json:"hora_fin"`
	Estado           tipos.Estado      `json:"estado"`
	AnulacionReserva *AnulacionReserva `json:"anulacion_reserva,omitempty"`
}
