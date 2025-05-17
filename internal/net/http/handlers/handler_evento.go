package handlers

import (
	"strconv"

	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/evento"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

func ListarEventos(c *fiber.Ctx) error {
	eventos, err := evento.ListarEventos()
	if err != nil {
		logs.Logger.Println("Error al listar eventos: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al listar eventos", nil))
	}
	return c.Status(fiber.StatusOK).JSON(models.Succes("Eventos obtenidos con éxito", eventos))
}

func ObtenerEventoPorId(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID inválido",
			"status":  400,
			"data":    nil,
		})
	}

	eventoInfo, err := evento.BuscarEventoPorID(id)
	if err != nil {
		logs.Logger.Println("Error al obtener evento por ID: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error al obtener el evento",
			"status":  500,
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Evento encontrado",
		"status":  200,
		"data":    eventoInfo,
	})
}