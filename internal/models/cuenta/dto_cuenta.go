package cuenta

import "github.com/MervanXD/backend_ClubYa/internal/models/tipos"

type DTOCuenta struct {
	Username        string                `json:"username" `
	Contrasena      string                `json:"contrasena" `
	IdPersona       int                   `json:"id_persona" `
	Postulante      bool                  `json:"postulante" `
	Rol             tipos.Rol             `json:"rol" `
	EstadoSolicitud tipos.EstadoSolicitud `json:"estado_solicitud" `
	IdMembresia     int                   `json:"id_membresia,omitempty"` // Solo si es titular
	IdSolicitud     int                   `json:"id_solicitud,omitempty"` // Solo si es titular
}
