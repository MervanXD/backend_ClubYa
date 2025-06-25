package utils

import (
	"fmt"
	"sort"
	"time"

	"github.com/MervanXD/backend_ClubYa/internal/models/bloque_tiempo"
)

// Interval representa un intervalo de tiempo con inicio y fin
type Interval struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// ConvertBloquesToIntervals convierte una lista de bloques de tiempo a una lista de intervalos
// de tiempo, donde cada bloque se convierte en un intervalo con hora de inicio y fin.
// Los bloques de tiempo deben tener los campos RangoInicio y RangoFin en formato "HH:mm".
// Si los bloques de tiempo no están en el formato correcto, la función retornará
// un error al intentar parsear las horas.
// Parámetros:
//   - bloques: lista de bloques de tiempo a convertir.
//
// Retorna:
//   - Una lista de intervalos de tiempo, donde cada intervalo tiene un campo Start y End en formato "HH:mm".
//   - Un error si ocurre algún problema al convertir los bloques de tiempo.
//   - Si la lista de bloques está vacía, retorna un error indicando que la lista de intervalos está vacía.
//
// Ejemplo de uso:
//   - bloques := []bloque_tiempo.BloqueTiempo{
//     {RangoInicio: "08:00", RangoFin: "10:00"},
//     {RangoInicio: "10:00", RangoFin: "12:00"},
//     }
//   - intervals := ConvertBloquesToIntervals(bloques)
//   - fmt.Println(intervals) // Output: [{08:00 10:00} {10:00 12:00}]
//
// // Ejemplo de error:
//   - bloques := []bloque_tiempo.BloqueTiempo{
//     {RangoInicio: "08:00", RangoFin: "10:00"},
//     {RangoInicio: "10:00", RangoFin: "09:00"}, // Hora de fin antes de hora de inicio
//     }
//   - intervals, err := ConvertBloquesToIntervals(bloques)
//   - if err != nil {
//     fmt.Println("Error:", err)
//     } else {
//     fmt.Println(intervals)
//     }
func ConvertBloquesToIntervals(bloques []bloque_tiempo.BloqueTiempo) []Interval {
	intervals := make([]Interval, len(bloques))
	for i, bloque := range bloques {
		intervals[i] = Interval{
			Start: bloque.RangoInicio,
			End:   bloque.RangoFin,
		}
	}
	return intervals
}

// ValidateAndSortIntervals valida y ordena los intervalos de tiempo
func ValidateAndSortIntervals(intervals []Interval) (string, string, error) {
	if len(intervals) == 0 {
		return "", "", fmt.Errorf("la lista de intervalos está vacía")
	}

	// Validar y parsear todos los intervalos primero
	parsed := make([]struct {
		original Interval
		start    time.Time
		end      time.Time
	}, len(intervals))

	for i, interval := range intervals {
		start, err := time.Parse("15:04", interval.Start)
		if err != nil {
			return "", "", fmt.Errorf("formato inválido en hora de inicio: %s", interval.Start)
		}

		end, err := time.Parse("15:04", interval.End)
		if err != nil {
			return "", "", fmt.Errorf("formato inválido en hora de fin: %s", interval.End)
		}

		if end.Before(start) {
			return "", "", fmt.Errorf("hora de fin es anterior a hora de inicio en intervalo: %s-%s", interval.Start, interval.End)
		}

		parsed[i] = struct {
			original Interval
			start    time.Time
			end      time.Time
		}{
			original: interval,
			start:    start,
			end:      end,
		}
	}

	// Ordenar intervalos por hora de inicio
	sort.Slice(parsed, func(i, j int) bool {
		return parsed[i].start.Before(parsed[j].start)
	})

	// Reconstruir lista ordenada y verificar continuidad
	sorted := make([]Interval, len(parsed))
	var minStart, maxEnd time.Time

	for i, p := range parsed {
		sorted[i] = p.original

		if i == 0 {
			minStart = p.start
			maxEnd = p.end
		} else {
			// Verificar si el intervalo actual comienza donde termina el anterior
			if !p.start.Equal(parsed[i-1].end) {
				return "", "", fmt.Errorf("intervalos no son consecutivos: %s != %s",
					parsed[i-1].original.End, p.original.Start)
			}

			// Actualizar hora de fin máxima
			if p.end.After(maxEnd) {
				maxEnd = p.end
			}
		}
	}

	// Convertir a strings "HH:mm"
	startStr := minStart.Format("15:04")
	endStr := maxEnd.Format("15:04")

	return startStr, endStr, nil
}
