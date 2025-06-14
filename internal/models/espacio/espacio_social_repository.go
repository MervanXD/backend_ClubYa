package espacio

import "context"

type EspacioSocialRepository interface {
	InsertarEspacioSocial(es EspacioSocial) error
	ObtenerEspaciosSociales() ([]EspacioSocial, error)
	ObtenerEspacioSocialPorID(ctx context.Context, id int) (*EspacioSocial, error)
	ObtenerEspaciosSocialesHorarios() ([]EspacioSocialHorarioDTO, error)
	ActualizarParcial(id int, dto EspacioSocialUpdateDTO) error 
}
