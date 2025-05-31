# backend_ClubYa

backend_ClubYa es el backend desarrollado en Go para la aplicación ClubYa, un sistema de gestión de clubes sociales y deportivos. Este backend expone una API RESTful para la administración de socios, familiares, reservas, eventos, membresías y más.

## Tabla de Contenidos

- [Características](#características)
- [Requisitos](#requisitos)
- [Instalación](#instalación)
- [Configuración](#configuración)
- [Ejecución](#ejecución)
- [Estructura del Proyecto](#estructura-del-proyecto)
- [Endpoints Principales](#endpoints-principales)
- [Logs](#logs)
- [Contribuciones](#contribuciones)
- [Licencia](#licencia)

---

## Características

- Gestión de socios y familiares
- Administración de reservas de espacios y eventos
- Control de membresías y solicitudes
- Manejo de reportes y pagos
- API RESTful desarrollada con Go y Fiber
- Registro de logs de aplicación

---

## Requisitos

- Go 1.20 o superior
- MySql (u otro motor de base de datos compatible)
- Air (opcional, para hot reload en desarrollo)

---

## Instalación

```bash
git clone https://github.com/tu-usuario/backend_ClubYa.git
cd backend_ClubYa
go mod download
```
---

## Configuración

- En desarrollo
```bash
cd cmd/server
go run main.go
```
- O usando Air para que se actualice automaticamente (opcional)
```bash
air
```
---
## Estructura del proyecto

```plaintext
backend_ClubYa/
│
├── cmd/                # Entrada principal del servidor
│   └── server/
│       └── main.go
├── config/             # Configuración de la aplicación
├── database/           # Conexión y lógica de base de datos
├── internal/           # Lógica de negocio y modelos
│   ├── api/
│   ├── models/
│   ├── net/
│   └── pkgs/
├── logs/               # Archivos de logs de la aplicación
├── go.mod
├── go.sum
└── [README.md](http://_vscodecontentref_/2)
```
---

## Endpoints Principales

> **Nota:** La documentación completa de la API está en desarrollo.  
> Algunos endpoints principales son:

- `GET    /api/socios` — Listar socios
- `GET    /api/familiares` — Listar familiares
- `GET    /api/reservas` — Listar reservas
- `GET    /api/eventos` — Listar eventos
- `GET    /api/membresias` — Listar membresías
- `GET    /api/solicitudes` — Listar solicitudes
- `POST   /api/socios` — Crear socio
- `POST   /api/familiares` — Crear familiar
- `POST   /api/reservas` — Crear reserva
- `POST   /api/eventos` — Crear evento
- `POST   /api/membresias` — Crear membresía
- `POST   /api/solicitudes` — Crear solicitud

Consulta el código en `internal/net/http/handlers/` para ver todos los endpoints disponibles.

---

## Logs

Los logs de la aplicación se almacenan en la carpeta `logs/`.  
Revisa estos archivos para depuración y monitoreo de errores.  
Cada archivo de log tiene el formato:  
```
logs/app-YYYY-MM-DD.log
```
Donde `YYYY-MM-DD` corresponde a la fecha del log.

---

