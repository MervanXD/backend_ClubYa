package handlers

import (
	"context"
	"strconv"
	"time"

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
	repo := persona.NewTitularRepositoryDB()
	idTitular, err := repo.InsertarTitular(personaData, idCuenta)
	if err != nil {
		logs.Logger.Println("Error al insertar la persona: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al insertar la persona", nil))
	}

	return c.Status(fiber.StatusCreated).JSON(models.Succes("Persona creada con éxito", idTitular))
}

func ObtenerTitularPorID(c *fiber.Ctx) error {
	titularIdStr := c.Params("idTitular")
	idTitular, err := strconv.Atoi(titularIdStr)
	if err != nil {
		logs.Logger.Println("ID de titular inválido:", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.BadRequest("ID de titular inválido", nil))
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()
	repo := persona.NewTitularRepositoryDB()
	titular, err := repo.ObtenerTitularPorID(ctx, idTitular)
	if err != nil {
		logs.Logger.Println("Error al obtener la información: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener la información", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Información de titular obtenida correctamente", titular))
}
