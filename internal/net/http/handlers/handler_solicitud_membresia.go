package handlers

import (
	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	solicitud "github.com/MervanXD/backend_ClubYa/internal/models/solicitud_membresia"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

func ListarSolicitudMembresia(c *fiber.Ctx) error {
	solicitudes, err := solicitud.ObtenerSolicitudesMembresia()
	if err != nil {
		logs.Logger.Fatal("Error al obtener las solicitudes de membresia: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener las solicitudes de membresia", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Solicitudes de membresia obtenidas con éxito", solicitudes))
}
