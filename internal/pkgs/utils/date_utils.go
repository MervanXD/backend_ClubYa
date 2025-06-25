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
