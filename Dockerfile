# Etapa 1: build
FROM golang:1.21-alpine

WORKDIR /app

# Copia los archivos go mod y sum, e instala dependencias
COPY go.mod go.sum ./
RUN go mod download


ENV PATH="/usr/local/go/bin:${PATH}"


# Copia el resto del código fuente
COPY . .

# Compila la aplicación (ajusta main.go si tu archivo principal tiene otro nombre)
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

# Etapa 2: imagen final
FROM alpine:latest

WORKDIR /root/

# Copia el binario compilado desde la etapa anterior
COPY --from=builder /app/app .

# Expone el puerto típico de Fiber (ajusta si usas otro)
EXPOSE 4000

# Ejecuta la app
CMD ["go", "test", "./internal/tests/..."]