package handlers

import (
	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/persona"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

func CrearTitular(c *fiber.Ctx) error {
	var personaData persona.Titular

	if err := c.BodyParser(&personaData); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}
	var idTitular int
	idTitular, err := persona.InsertarTitular(personaData)
	if err != nil {
		logs.Logger.Println("Error al insertar la persona: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al insertar la persona", nil))
	}

	return c.Status(fiber.StatusCreated).JSON(models.Succes("Persona creada con éxito", idTitular))
}

func RegistrarPostulante(c *fiber.Ctx) error {
	var personaData persona.Titular
	if err := c.BodyParser(&personaData); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}

	var idTitular int
	idTitular, err := persona.InsertarTitular(personaData)
	if err != nil {
		logs.Logger.Println("Error al insertar la persona: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al insertar la persona", nil))
	}

	err = CrearCuentaDummy(idTitular)
	if err != nil {
		logs.Logger.Println("Error al crear cuenta", err, idTitular)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al registrar postulante", nil))
	}

	return c.Status(fiber.StatusCreated).JSON(models.Succes("Persona creada con éxito", idTitular))
}
