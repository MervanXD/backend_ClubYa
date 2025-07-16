package reports

import (
    "strconv"

    "github.com/MervanXD/backend_ClubYa/internal/models/academia"
    "github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
    "github.com/johnfercher/maroto/v2/pkg/components/row"
    "github.com/johnfercher/maroto/v2/pkg/components/text"
    "github.com/johnfercher/maroto/v2/pkg/consts/align"
    "github.com/johnfercher/maroto/v2/pkg/core"
    "github.com/johnfercher/maroto/v2/pkg/props"
)

func GeneraReporteAcademias(academias []academia.ReporteAcademiaDTO) ([]byte, error) {
    m, err := getReporteBase()
    if err != nil {
        return nil, err
    }

    // Agregar título de la tabla
    m.AddRow(20, getTituloTabla("Reporte de Academias Deportivas")...)

    // Agregar encabezados de la tabla
    encabezados := []Encabezado{
        {Nombre: "Nombre Academia", Size: 2},
        {Nombre: "Deporte", Size: 2},
        {Nombre: "Entrenador", Size: 2},
        {Nombre: "Inscritos", Size: 1},
        {Nombre: "Fecha Inicio", Size: 1},
        {Nombre: "Fecha Fin", Size: 1},
        {Nombre: "Ingreso Total", Size: 3},
    }
    encabezadosCols := getEncabezadosTabla(encabezados)
    m.AddRow(10, encabezadosCols...)

    // Agregar datos de las academias
    academiasRows := getDatosTablaAcademias(academias)
    m.AddRows(academiasRows...)

    // Genera y retorna el PDF
    document, err := m.Generate()
    if err != nil {
        return nil, err
    }

    return document.GetBytes(), nil
}

func getDatosTablaAcademias(academias []academia.ReporteAcademiaDTO) []core.Row {
    var rows []core.Row
    for i, entry := range academias {
        fechaInicio := utils.FormatDate(entry.FechaInicio)
        fechaFin := utils.FormatDate(entry.FechaFin)
        ingresoTotal := "S/ " + strconv.FormatFloat(entry.IngresoTotal, 'f', 2, 64)
        
        r := row.New(10).Add(
            text.NewCol(2, entry.NombreAcademia, props.Text{Size: 8, Align: align.Left}),
            text.NewCol(2, entry.Deporte, props.Text{Size: 8, Align: align.Center}),
            text.NewCol(2, entry.Entrenador, props.Text{Size: 8, Align: align.Center}),
            text.NewCol(1, strconv.Itoa(entry.Inscritos), props.Text{Size: 8, Align: align.Center}),
            text.NewCol(1, fechaInicio, props.Text{Size: 8, Align: align.Center}),
            text.NewCol(1, fechaFin, props.Text{Size: 8, Align: align.Center}),
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