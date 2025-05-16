package handlers

import (
	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/evento"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

func ListarEventos(c *fiber.Ctx) error {
	eventos, err := evento.ListarEventos()
	if err != nil {
		logs.Logger.Fatal("Error al listar eventos: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al listar eventos", nil))
	}
	return c.Status(fiber.StatusOK).JSON(models.Succes("Eventos obtenidos con éxito", eventos))
}
