package reports

import (
	"strconv"

	"github.com/MervanXD/backend_ClubYa/internal/models/membresia"
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func GeneraReporteMembresias(membresias []membresia.MembresiaDTO) ([]byte, error) {
	m, err := getReporteBase()
	if err != nil {
		return nil, err
	}

	// Agregar titulo de la tabla
	m.AddRow(20, getTituloTabla("Reporte de Membresias")...)

	// Agregar encabezados de la tabla
	encabezados := []Encabezado{
		{Nombre: "Id", Size: 1},
		{Nombre: "Nombre Titular", Size: 3},
		{Nombre: "Cuota", Size: 2},
		{Nombre: "Fecha Inicio", Size: 2},
		{Nombre: "Fecha Fin", Size: 2},
	}
	encabezadosCols := getEncabezadosTabla(encabezados)
	m.AddRow(10, encabezadosCols...)

	// Agregar datos de las membresias
	membresiasRows := getDatosTabla(membresias)
	m.AddRows(membresiasRows...)

	// Genera y retorna el PDF
	document, err := m.Generate()
	if err != nil {
		return nil, err
	}

	return document.GetBytes(), nil
}

func getDatosTabla(membresias []membresia.MembresiaDTO) []core.Row {
	var rows []core.Row
	for i, entry := range membresias {
		fechaInicio := utils.FormatDate(entry.Membresia.FechaInicio)
		fechaFin := utils.FormatDate(entry.Membresia.FechaFin.String)
		r := row.New(10).Add(
			text.NewCol(1, strconv.Itoa(entry.Membresia.Id), props.Text{Size: 9, Align: align.Center}),
			text.NewCol(3, entry.NombreTitular, props.Text{Size: 9, Align: align.Center}),
			text.NewCol(2, strconv.FormatFloat(entry.CuotaBase, 'f', 2, 64), props.Text{Size: 9, Align: align.Center}),
			text.NewCol(2, fechaInicio, props.Text{Size: 9, Align: align.Center}),
			text.NewCol(2, fechaFin, props.Text{Size: 9, Align: align.Center}),
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
