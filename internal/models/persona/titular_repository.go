package persona

import "context"

type TitularRepository interface {
	InsertarTitular(p Titular, idCuenta int) (int, error) 
	ObtenerTitularPorID(ctx context.Context, idPersona int) (*Titular, error)
}
