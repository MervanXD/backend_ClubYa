package handlers

import (
	"database/sql"
	"strconv"

	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/academia"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

func ListarAcademias(c *fiber.Ctx) error {
	academiasDeportivas, err := academia.ObtenerAcademias()
	if err != nil {
		logs.Logger.Println("Error al obtener las academias deportivas: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener las academias", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Academias deportivas obtenidas con exito", academiasDeportivas))
}

func ObtenerAcademiaId(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Logger.Println("ID inválido: ", err)
		return c.Status(fiber.StatusBadRequest).SendString("ID inválido")
	}
	academia, err := academia.ObtenerAcademiaPorId(id)
	if err != nil {
		if err == sql.ErrNoRows {
			logs.Logger.Println("No se encontró ninguna academia con ese ID", err)
			return c.Status(fiber.StatusNotFound).JSON(models.NotFound("No encontrado"))
		}
		logs.Logger.Println("Error al obtener la academia: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener la infromacion de la academia", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Se logró obtener la informacion de la academia", academia))
}
