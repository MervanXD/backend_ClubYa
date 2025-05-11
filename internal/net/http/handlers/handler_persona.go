package handlers

import (
	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/persona"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

func CrearPersona(c *fiber.Ctx) error {
	var personaData persona.Persona

	if err := c.BodyParser(&personaData); err != nil {
		logs.Logger.Fatal("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}

	if err := persona.InsertarPersona(personaData); err != nil {
		logs.Logger.Fatal("Error al insertar la persona: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al insertar la persona", nil))
	}

	return c.Status(fiber.StatusCreated).JSON(models.Succes("Persona creada con éxito", nil))
}