package handlers

import (
	"strconv"

	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/tarifas"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

func ObtenerTarifaMembresia(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		logs.Logger.Println("ID inválido: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID inválido", nil))
	}

	repo := tarifas.NewTarifaMembresiaRespositoryDB()
	tarifaMembresia, err := repo.ObtenerTarifasMembresiaPorId(int64(id))
	if err != nil {
		logs.Logger.Println("Error al obtener la tarifa de membresía: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener la tarifa de membresía", err))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Tarifa de membresía obtenida con éxito", tarifaMembresia))
}
