package handlers

import (
	"strconv"

	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/solicitud_membresia"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

type actualizarEstadoRequest struct {
	EstadoSolicitud string `json:"estado_solicitud"`
}

func ActualizarEstadoSolicitud(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		logs.Logger.Println("ID inválido: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID inválido", nil))
	}

	var body actualizarEstadoRequest
	if err := c.BodyParser(&body); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}

	err = solicitud.ActualizarEstadoSolicitud(id, body.EstadoSolicitud)
	if err != nil {
		logs.Logger.Println("Error al actualizar estado: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("No se pudo actualizar el estado", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Estado actualizado correctamente", nil))
}
