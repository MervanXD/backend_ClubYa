package handlers

import (
	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	inscripcion "github.com/MervanXD/backend_ClubYa/internal/models/inscripcion_evento"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

type InscripcionRequest struct {
	FidPersona        int `json:"id_persona" validate:"required"`
	IdEvento          int `json:"id_evento" validate:"required"`
	CantidadInvitados int `json:"cantidad_invitados"`
}

func RegistrarInscripcionEvento(c *fiber.Ctx) error {
	var requests []InscripcionRequest

	if err := c.BodyParser(&requests); err != nil {
		logs.Logger.Println("Error al parsear el body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.BadRequest("Datos inválidos", nil))
	}

	if len(requests) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.BadRequest("La lista de inscripciones no puede estar vacía", nil))
	}

	for _, req := range requests {
		if err := inscripcion.RegistrarInscripcion(req.FidPersona, req.IdEvento, req.CantidadInvitados); err != nil {
			logs.Logger.Printf("Error al registrar inscripción para persona %d: %v", req.FidPersona, err)
			return c.Status(fiber.StatusInternalServerError).JSON(models.Error("No se pudo registrar la inscripción para una o más personas", nil))
		}
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Todas las inscripciones fueron registradas exitosamente", nil))
}
