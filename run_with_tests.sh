#!/bin/bash

echo "Ejecutando tests..."
go test ./...

if [ $? -ne 0 ]; then
  echo "❌ Al menos un test falló. No se ejecutará el programa."
  exit 1
fi

echo "✅ Todos los tests pasaron. Ejecutando el programa..."
go run ./cmd/server/main.go