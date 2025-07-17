# Etapa 1: build
FROM golang:1.23.10-alpine AS builder

# Instala dependencias del sistema
RUN apk add --no-cache git

# Crea el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia go.mod y go.sum y descarga dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el resto del código
COPY . .

# Compila el binario (sin dependencias de C y con flags para tamaño)
RUN go build -ldflags="-s -w" -o app ./cmd/server

# Etapa 2: contenedor mínimo para correr
FROM alpine:latest

# Crea una carpeta para el binario
WORKDIR /root/

# Copia solo el binario desde la etapa anterior
COPY --from=builder /app/app .
COPY --from=builder /app/.env .

# Verificar que .env se copió (DEBUG)
RUN ls -la | grep .env || echo "ERROR: .env no se copió"


# Opcional: indicar que es un contenedor
ENV DOCKER_CONTAINER=true

# Expone el puerto que usas (4000 por ejemplo)
EXPOSE 4000

# Comando por defecto
CMD ["./app"]
