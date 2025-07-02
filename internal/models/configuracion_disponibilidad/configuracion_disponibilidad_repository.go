package configuracion_disponibilidad

import (
	"context"
	"database/sql"
)

type ConfiguracionDisponibilidadRepository interface {
	InsertarConfiguracionDisponibilidad(ctx context.Context, tx *sql.Tx, configuracion ConfiguracionDisponibilidad) error
	ObtenerConfiguracionDisponibilidad(ctx context.Context, tx *sql.Tx, idEspacio int) ([]ConfiguracionDisponibilidad, error)
	EliminarConfiguracionDisponibilidad(ctx context.Context, tx *sql.Tx, configuracion ConfiguracionDisponibilidad) error
}
