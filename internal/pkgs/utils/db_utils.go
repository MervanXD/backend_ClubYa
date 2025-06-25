package utils

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

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

func ExecSP(ctx context.Context, tx *sql.Tx, name string, args ...interface{}) error {
	q := fmt.Sprintf("CALL ingesoft.%s(%s)", name, placeholders(len(args)))
	if _, err := tx.ExecContext(ctx, q, args...); err != nil {
		return fmt.Errorf("CALL %s: %w", name, err)
	}
	return nil
}

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
