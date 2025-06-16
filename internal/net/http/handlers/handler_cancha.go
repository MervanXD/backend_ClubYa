package handlers

import (
	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/espacio"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

func ListarCanchasHorarios(c *fiber.Ctx) error {
	repo := espacio.NewCanchaRepositoryDB()
	horariosCanchas, err := repo.ObtenerCanchasHorarios()
	if err != nil {
		logs.Logger.Println("Error al obtener las lozas deportivas: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener las lozas deportivas", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Horarios de lozas deportivas obtenidos con éxito", horariosCanchas))
}

func ListarCanchasConfiguracion(c *fiber.Ctx) error {
	repo := espacio.NewCanchaRepositoryDB()
	canchas, err := repo.ObtenerCanchasConfiguracion()
	if err != nil {
		logs.Logger.Println("Error al obtener las canchas: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener las canchas", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Canchas obtenidos con éxito", canchas))
}
