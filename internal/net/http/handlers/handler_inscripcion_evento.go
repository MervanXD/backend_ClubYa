package handlers

import (
    "github.com/MervanXD/backend_ClubYa/internal/api/models"
    inscripcion "github.com/MervanXD/backend_ClubYa/internal/models/inscripcion_evento"
    "github.com/MervanXD/backend_ClubYa/logs"
    "github.com/gofiber/fiber/v2"
)

type InscripcionRequest struct {
    FidPersona         int `json:"fidPersona"`
    IdEvento           int `json:"idEvento"`
    CantidadInvitados  int `json:"cantidadInvitados"`
}

func RegistrarInscripcionEvento(c *fiber.Ctx) error {
    var req InscripcionRequest

    if err := c.BodyParser(&req); err != nil {
        logs.Logger.Println("Error al parsear el body:", err)
        return c.Status(fiber.StatusBadRequest).JSON(models.Error("Datos inválidos", nil))
    }

    if err := inscripcion.RegistrarInscripcion(req.FidPersona, req.IdEvento, req.CantidadInvitados); err != nil {
        logs.Logger.Println("Error al registrar inscripción:", err)
        return c.Status(fiber.StatusInternalServerError).JSON(models.Error("No se pudo registrar la inscripción", nil))
    }

    return c.Status(fiber.StatusOK).JSON(models.Succes("Inscripción registrada exitosamente", nil))
}
