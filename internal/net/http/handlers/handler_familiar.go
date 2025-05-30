package handlers

import (
	"strconv"

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

	idSolicitud, err := persona.RegistrarFamiliares(req)
	if err != nil {
		logs.Logger.Println("Error al registrar familiares: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al registrar familiares", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Familiares registrados correctamente", idSolicitud))
}

func ObtenerFamiliaresPorTitular(c *fiber.Ctx) error {
	idTitularStr := c.Params("id")
	idTitular, err := strconv.Atoi(idTitularStr)
	if err != nil {
		logs.Logger.Println("ID inválido:", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID inválido", nil))
	}

	familiares, err := persona.ObtenerIdsFamiliaresPorTitular(idTitular)
	if err != nil {
		logs.Logger.Println("Error al obtener familiares:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("No se pudo obtener los familiares", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Familiares obtenidos correctamente", familiares))
}

func ObtenerFamiliarPorIdPersona(c *fiber.Ctx) error {
	idPersonStr := c.Params("id")
	idPersona, err := strconv.Atoi(idPersonStr)
	if err != nil {
		logs.Logger.Println("ID de persona inválido:", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID invalido", nil))
	}
	familiar, err := persona.ObtenerFamiliarPorIDPersona(idPersona)
	if err != nil {
		logs.Logger.Println("Error al obtener los datos del familiar:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("No se pudo obtener los datos del familiar", nil))
	}
	return c.Status(fiber.StatusOK).JSON(models.Succes("Familiar obtenido correctamente", familiar))
}
