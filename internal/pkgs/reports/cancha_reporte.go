package reports

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/espacio"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func GeneraReporteCanchas(reporte espacio.ReporteCanchaDTO) ([]byte, error) {
	m, err := getReporteBase()
	if err != nil {
		return nil, err
	}

	// Título principal del reporte
	m.AddRow(20, getTituloTabla("Reporte de Canchas Deportivas")...)

	// Sección de métricas generales
	metricasGenerales := []MetricaItem{
		{
			Label: "Ingresos Totales",
			Valor: FormatearMoneda(reporte.Metricas.IngresosTotales),
			Color: getBlueColor(),
		},
		{
			Label: "Horas Reservadas",
			Valor: FormatearNumero(reporte.Metricas.HorasReservadas) + " hrs",
			Color: getBlueColor(),
		},
		{
			Label: "Reservas Confirmadas",
			Valor: FormatearNumero(reporte.Metricas.ReservasConfirmadas),
			Color: getBlueColor(),
		},
		{
			Label: "Reservas Anuladas",
			Valor: FormatearNumero(reporte.Metricas.ReservasAnuladas),
			Color: getRedColor(),
		},
	}

	metricasRows := CrearSeccionMetricas("MÉTRICAS GENERALES", metricasGenerales)
	m.AddRows(metricasRows...)

	// Título de la tabla de socios
	m.AddRow(15, getTituloTabla("DETALLE POR SOCIO")...)

	// Encabezados de la tabla de socios
	encabezados := []Encabezado{
		{Nombre: "Nombre del Socio", Size: 3},
		{Nombre: "Reservas Confirmadas", Size: 2},
		{Nombre: "Reservas Anuladas", Size: 2},
		{Nombre: "Horas Reservadas", Size: 2},
		{Nombre: "Ingresos por Socio", Size: 3},
	}
	encabezadosCols := getEncabezadosTabla(encabezados)
	m.AddRow(10, encabezadosCols...)

	// Datos de la tabla de socios
	sociosRows := getDatosTablaSocios(reporte.DatosSocios)
	m.AddRows(sociosRows...)

	// Generar y retornar el PDF
	document, err := m.Generate()
	if err != nil {
		return nil, err
	}

	return document.GetBytes(), nil
}

func getDatosTablaSocios(socios []espacio.ReporteCanchaSocioDTO) []core.Row {
	var rows []core.Row
	for i, socio := range socios {
		ingresosPorSocio := FormatearMoneda(socio.IngresosPorSocio)
		horasReservadas := FormatearNumero(socio.HorasReservadas) + " hrs"

		r := row.New(10).Add(
			text.NewCol(3, socio.NombreSocio, props.Text{Size: 8, Align: align.Left}),
			text.NewCol(2, FormatearNumero(socio.ReservasConfirmadas), props.Text{Size: 8, Align: align.Center, Color: getBlueColor()}),
			text.NewCol(2, FormatearNumero(socio.ReservasAnuladas), props.Text{Size: 8, Align: align.Center, Color: getRedColor()}),
			text.NewCol(2, horasReservadas, props.Text{Size: 8, Align: align.Center}),
			text.NewCol(3, ingresosPorSocio, props.Text{Size: 8, Align: align.Right, Color: getBlueColor()}),
		)

		// Alternar color de fondo de las filas
		if i%2 == 0 {
			r.WithStyle(&props.Cell{
				BackgroundColor: &props.Color{Red: 240, Green: 240, Blue: 240},
			})
		}

		rows = append(rows, r)
	}
	return rows
}
