package handlers

import (
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
