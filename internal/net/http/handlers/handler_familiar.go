package handlers

import (
	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/persona"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

func InsertarFamiliares(c *fiber.Ctx) error {
	var req persona.FamiliarResquest
	if err := c.BodyParser(&req); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}

	if err := persona.RegistrarFamiliares(req); err != nil {
		logs.Logger.Println("Error al registrar familiares: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al registrar familiares", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Familiares registrados correctamente", nil))

}
