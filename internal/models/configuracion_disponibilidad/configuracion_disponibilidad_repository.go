package configuracion_disponibilidad

import (
	"context"
	"database/sql"
)

type ConfiguracionDisponibilidadRepository interface {
	InsertarConfiguracionDisponibilidad(ctx context.Context, tx *sql.Tx, configuracion ConfiguracionDisponibilidad) error
	ObtenerConfiguracionDisponibilidad(ctx context.Context, idEspacio int, tx ...*sql.Tx) ([]ConfiguracionDisponibilidad, error)
	EliminarConfiguracionDisponibilidad(ctx context.Context, tx *sql.Tx, configuracion ConfiguracionDisponibilidad) error
	ActualizarConfiguracionDisponibilidad(ctx context.Context, configuraciones []ConfiguracionDisponibilidadDTO) error
	ObtenerCrucesEspacio(ctx context.Context, idEspacio int) ([]NroCruces, error)
}
