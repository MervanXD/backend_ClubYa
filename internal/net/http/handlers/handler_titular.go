package handlers

import (
	"strconv"

	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/persona"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

func CrearTitular(c *fiber.Ctx) error {
	var personaData persona.Titular
	idC := c.Params("idCuenta")

	logs.Logger.Println("ID recibido:", idC)

	if err := c.BodyParser(&personaData); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}

	idCuenta, err := strconv.Atoi(idC)
	if err != nil {
		logs.Logger.Println("ID de membresía inválido:", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID de membresía inválido", nil))
	}

	idTitular, err := persona.InsertarTitular(personaData, idCuenta)
	if err != nil {
		logs.Logger.Println("Error al insertar la persona: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al insertar la persona", nil))
	}

	return c.Status(fiber.StatusCreated).JSON(models.Succes("Persona creada con éxito", idTitular))
}

type TitularRequest struct {
	IdPersona int `json:"id_persona" validate:"required"`
}

func ObtenerTitularPorID(c *fiber.Ctx) error {
	var titularRequest TitularRequest

	if err := c.BodyParser(&titularRequest); err != nil {
		logs.Logger.Println("Error al parsear el body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.BadRequest("Datos inválidos", nil))
	}

	titular, err := persona.ObtenerTitularPorID(c.Context(), titularRequest.IdPersona)
	if err != nil {
		logs.Logger.Println("Error al obtener la información: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener la información", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Información de titular obtenida correctamente", titular))
}
