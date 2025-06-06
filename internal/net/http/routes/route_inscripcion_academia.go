package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RutasInscripcionAcademia(app *fiber.App) {
	app.Get("/academias-deportivas", handlers.ListarAcademias)
	app.Get("/academia-deportiva-informacion/:id", handlers.ObtenerAcademiaId)
}
