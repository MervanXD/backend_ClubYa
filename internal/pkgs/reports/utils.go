package reports

import (
	"strconv"

	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func getEncabezadosTabla(encabezados []Encabezado) []core.Col {
	var cols []core.Col
	for _, encabezado := range encabezados {
		col := text.NewCol(encabezado.Size, encabezado.Nombre, props.Text{
			Top:   1.5,
			Size:  9,
			Style: fontstyle.Bold,
			Align: align.Center,
		})
		cols = append(cols, col)
	}
	return cols
}

func getTituloTabla(title string) []core.Col {
	return []core.Col{
		text.NewCol(10, title, getTituloStyle(align.Center)),
	}
}

func getPiePagina() core.Row {
	return row.New(20).Add(
		col.New(12).Add(
			text.New("ClubYA!", props.Text{
				Top:   13,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Left,
				Color: getBlueColor(),
			}),
			text.New("Lima, Perú - Club social", props.Text{
				Top:   16,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Left,
				Color: getBlueColor(),
			}),
		),
	)
}

func getReporteBase() (core.Maroto, error) {
	cfg := config.NewBuilder().
		WithPageNumber().
		Build()
	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	err := m.RegisterFooter(getPiePagina())
	if err != nil {
		logs.Logger.Println(err.Error())
		return nil, err
	}

	return m, nil
}

// ------------------------------ Tipos ------------------------------
type Encabezado struct {
	Nombre string
	Size   int
}

// ------------------------------ Estilos Texto ------------------------------
func getTituloStyle(align align.Type) props.Text {
	return props.Text{
		Top:   3,
		Size:  20,
		Style: fontstyle.Bold,
		Align: align,
	}
}

// ------------------------------ Colores ------------------------------

func getGrayColor() *props.Color {
	return &props.Color{
		Red:   200,
		Green: 200,
		Blue:  200,
	}
}

func getBlueColor() *props.Color {
	return &props.Color{
		Red:   10,
		Green: 10,
		Blue:  150,
	}
}

func getRedColor() *props.Color {
	return &props.Color{
		Red:   150,
		Green: 10,
		Blue:  10,
	}
}

// Función reutilizable para crear sección de métricas
func CrearSeccionMetricas(titulo string, metricas []MetricaItem) []core.Row {
	var rows []core.Row

	// Título de la sección
	tituloRow := row.New(15).Add(
		text.NewCol(12, titulo, props.Text{
			Size:  12,
			Style: fontstyle.Bold,
			Align: align.Center,
			Color: getBlueColor(),
		}),
	)
	rows = append(rows, tituloRow)

	// Crear filas de métricas (2 métricas por fila)
	for i := 0; i < len(metricas); i += 2 {
		metricaRow := row.New(12)

		// Primera métrica
		metricaRow.Add(
			text.NewCol(3, metricas[i].Label+":", props.Text{
				Size:  10,
				Style: fontstyle.Bold,
				Align: align.Right,
			}),
			text.NewCol(3, metricas[i].Valor, props.Text{
				Size:  10,
				Align: align.Left,
				Color: metricas[i].Color,
			}),
		)

		// Segunda métrica (si existe)
		if i+1 < len(metricas) {
			metricaRow.Add(
				text.NewCol(3, metricas[i+1].Label+":", props.Text{
					Size:  10,
					Style: fontstyle.Bold,
					Align: align.Right,
				}),
				text.NewCol(3, metricas[i+1].Valor, props.Text{
					Size:  10,
					Align: align.Left,
					Color: metricas[i+1].Color,
				}),
			)
		} else {
			// Llenar espacio vacío
			metricaRow.Add(
				text.NewCol(3, "", props.Text{}),
				text.NewCol(3, "", props.Text{}),
			)
		}

		rows = append(rows, metricaRow)
	}

	// Línea separadora
	separadorRow := row.New(8).Add(
		text.NewCol(12, "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━", props.Text{
			Size:  8,
			Align: align.Center,
			Color: getGrayColor(),
		}),
	)
	rows = append(rows, separadorRow)

	return rows
}

// Función para formatear moneda
func FormatearMoneda(valor float64) string {
	return "S/ " + strconv.FormatFloat(valor, 'f', 2, 64)
}

// Función para formatear número con separadores
func FormatearNumero(valor int) string {
	return strconv.Itoa(valor)
}
