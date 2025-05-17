package handlers

import (
	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/membresia"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func ObtenerMembresiaPorSolicitud(c *fiber.Ctx) error {
	idSolicitudStr := c.Params("id")
	idSolicitud, err := strconv.Atoi(idSolicitudStr)
	if err != nil {
		logs.Logger.Println("ID de solicitud inválido: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID de solicitud inválido", nil))
	}

	idMembresia, err := membresia.BuscarMembresiaPorSolicitud(idSolicitud)
	if err != nil {
		logs.Logger.Println("Error al buscar membresía: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al buscar membresía", nil))
	}

	if idMembresia == 0 {
		return c.Status(fiber.StatusNotFound).JSON(models.Error("No se encontró una membresía para esta solicitud", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Membresía encontrada", idMembresia))
}
