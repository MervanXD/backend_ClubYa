package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasInscripcionAcademia(api fiber.Router) {
	api.Get("/academias-deportivas", handlers.ListarAcademias)
	api.Get("/academia-deportiva-informacion/:id", handlers.ObtenerAcademiaId)
	api.Get("/academia-deportiva-informacion-familiares/:id", handlers.ListarFamiliaresSocioInscritos)
	api.Post("/academia/inscripcion", handlers.RegistrarInscripcionAcademia)
	api.Put("/anulacion-inscripcion-academia", handlers.AnularInscripcionAcademia)
	api.Put("/anulacion-inscripcion-academia/aceptar/:id", handlers.AceptarDevolucionAnulacionAcademia)
}
