package reports

import "github.com/johnfercher/maroto/v2/pkg/props"

// Estructura para métricas generales reutilizable
type MetricaItem struct {
	Label string
	Valor string
	Color *props.Color
}
