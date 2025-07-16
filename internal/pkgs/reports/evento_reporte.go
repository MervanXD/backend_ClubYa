package reports

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/evento"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func GeneraReporteEventos(eventos []evento.ReporteEventoDTO) ([]byte, error) {
	m, err := getReporteBase()
	if err != nil {
		return nil, err
	}

	// Título principal del reporte
	m.AddRow(25, getTituloTabla("Reporte de Eventos")...)

	// Encabezados de la tabla
	encabezados := []Encabezado{
		{Nombre: "Nombre del Evento", Size: 3},
		{Nombre: "Fecha del Evento", Size: 2},
		{Nombre: "Duración", Size: 2},
		{Nombre: "Nro. Inscritos", Size: 2},
		{Nombre: "Ingreso Total", Size: 3},
	}
	encabezadosCols := getEncabezadosTabla(encabezados)
	m.AddRow(10, encabezadosCols...)

	// Datos de la tabla de eventos
	eventosRows := getDatosTablaEventos(eventos)
	m.AddRows(eventosRows...)

	// Generar y retornar el PDF
	document, err := m.Generate()
	if err != nil {
		return nil, err
	}

	return document.GetBytes(), nil
}

func getDatosTablaEventos(eventos []evento.ReporteEventoDTO) []core.Row {
	var rows []core.Row
	for i, evento := range eventos {
		ingresoTotal := FormatearMoneda(evento.IngresoTotal)
		nroInscritos := FormatearNumero(evento.NroInscritos)
		duracionHoras := ConvertirHoraAHoras(evento.Duracion) + " hrs"

		r := row.New(10).Add(
			text.NewCol(3, evento.NombreEvento, props.Text{Size: 8, Align: align.Left}),
			text.NewCol(2, evento.FechaEvento, props.Text{Size: 8, Align: align.Center}),
			text.NewCol(2, duracionHoras, props.Text{Size: 8, Align: align.Center}),
			text.NewCol(2, nroInscritos, props.Text{Size: 8, Align: align.Center, Color: getBlueColor()}),
			text.NewCol(3, ingresoTotal, props.Text{Size: 8, Align: align.Right, Color: getBlueColor()}),
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
