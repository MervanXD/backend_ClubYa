package utils

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

// placeholders genera una cadena de texto que contiene una cantidad especificada de signos de interrogación ("?"),
// separados por comas y espacios. Es útil para construir consultas SQL con parámetros.
// Por ejemplo, placeholders(3) devuelve "?, ?, ?".
// Si n es menor o igual a 0, retorna una cadena vacía.
//
// Parámetros:
//   - n: número de marcadores de posición a generar.
//
// Retorna:
//   - Una cadena con n signos de interrogación separados por coma y espacio.
func placeholders(n int) string {
	if n <= 0 {
		return ""
	}
	parts := make([]string, n)
	for i := range parts {
		parts[i] = "?"
	}
	// join with comma+space
	return strings.Join(parts, ", ")
}

// ExecSP ejecuta un procedimiento almacenado en la base de datos utilizando una transacción existente.
// Construye la consulta SQL utilizando el nombre del procedimiento y los argumentos proporcionados.
// Si ocurre un error durante la ejecución del procedimiento, lo envuelve y lo retorna; de lo contrario, retorna nil.
//
// Parámetros:
//   - ctx: contexto para controlar la cancelación y los tiempos de espera de la operación.
//   - tx: transacción SQL activa donde se ejecutará el procedimiento almacenado.
//   - name: nombre del procedimiento almacenado a ejecutar.
//   - args: argumentos que se pasarán al procedimiento almacenado.
//
// Retorna:
//   - error: error si la ejecución falla, nil en caso contrario.
func ExecSP(ctx context.Context, tx *sql.Tx, name string, args ...interface{}) error {
	q := fmt.Sprintf("CALL ingesoft.%s(%s)", name, placeholders(len(args)))
	if _, err := tx.ExecContext(ctx, q, args...); err != nil {
		return fmt.Errorf("CALL %s: %w", name, err)
	}
	return nil
}

// ExecSPWithOut ejecuta un procedimiento almacenado que tiene un parámetro de salida (OUT).
// Utiliza una transacción existente para ejecutar el procedimiento y luego recupera el valor del parámetro OUT.
// Construye la consulta SQL para llamar al procedimiento y luego realiza una consulta para obtener el valor del parámetro OUT.
// // Parámetros:
//   - ctx: contexto para controlar la cancelación y los tiempos de espera de la operación.
//   - tx: transacción SQL activa donde se ejecutará el procedimiento almacenado.
//   - name: nombre del procedimiento almacenado a ejecutar.
//   - outVar: nombre de la variable de usuario OUT que se utilizará para almacenar el resultado.
//   - dest: puntero al destino donde se almacenará el valor del parámetro OUT.
//   - args: argumentos que se pasarán al procedimiento almacenado.
//
// // Retorna:
//   - error: error si la ejecución falla o si no se puede recuperar el valor del parámetro OUT, nil en caso contrario.
func ExecSPWithOut(ctx context.Context, tx *sql.Tx, name, outVar string, dest interface{}, args ...interface{}) error {
	// Build the CALL clause: e.g. "CALL ingesoft.MyProc(?, ?, @outVar)"
	inPlaceholders := placeholders(len(args))
	var callArgs string
	if len(args) > 0 {
		callArgs = fmt.Sprintf("%s, @%s", inPlaceholders, outVar)
	} else {
		callArgs = fmt.Sprintf("@%s", outVar)
	}
	callQ := fmt.Sprintf("CALL ingesoft.%s(%s)", name, callArgs)

	if _, err := tx.ExecContext(ctx, callQ, args...); err != nil {
		return fmt.Errorf("CALL %s: %w", name, err)
	}

	// Fetch the OUT user-variable
	selectQ := fmt.Sprintf("SELECT @%s", outVar)
	if err := tx.QueryRowContext(ctx, selectQ).Scan(dest); err != nil {
		return fmt.Errorf("select @%s: %w", outVar, err)
	}

	return nil
}

// QuerySP ejecuta un procedimiento almacenado y retorna el resultado de la consulta.
// Puede ejecutarse tanto en una transacción como directamente en la base de datos.
//
// Parámetros:
//   - ctx: contexto para controlar la cancelación y los tiempos de espera.
//   - dbContext: puede ser *sql.DB (conexión directa) o *sql.Tx (transacción).
//   - name: nombre del procedimiento almacenado.
//   - args: argumentos para el procedimiento.
//
// Retorna:
//   - *sql.Rows: resultado de la consulta.
//   - error: error si la ejecución falla.
func QuerySP(ctx context.Context, dbContext interface{}, name string, args ...interface{}) (*sql.Rows, error) {
	// Construimos la consulta
	q := fmt.Sprintf("CALL ingesoft.%s(%s)", name, placeholders(len(args)))

	// Determinamos si es transacción o conexión directa
	switch v := dbContext.(type) {
	case *sql.DB:
		return v.QueryContext(ctx, q, args...)
	case *sql.Tx:
		return v.QueryContext(ctx, q, args...)
	default:
		return nil, fmt.Errorf("dbContext debe ser *sql.DB o *sql.Tx, recibido: %T", dbContext)
	}
}
