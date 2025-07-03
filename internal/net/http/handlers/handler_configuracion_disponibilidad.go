package handlers

import (
	"strconv"

	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/configuracion_disponibilidad"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

func ListarConfiguracionDisponibilidad(c *fiber.Ctx) error {
	idEspacio, err := strconv.Atoi(c.Params("id_espacio"))
	if err != nil {
		logs.Logger.Println("Error al convertir el id_espacio: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al convertir el id_espacio", nil))
	}
	repo := configuracion_disponibilidad.NewConfiguracionDisponibilidadRepositoryDB()
	configuraciones, err := repo.ObtenerConfiguracionDisponibilidad(c.Context(), idEspacio)
	if err != nil {
		logs.Logger.Println("Error al obtener las configuraciones de disponibilidad: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener las configuraciones de disponibilidad", nil))
	}
	return c.Status(fiber.StatusOK).JSON(models.Succes("Configuraciones de disponibilidad obtenidas con éxito", configuraciones))
}

func ActualizarConfiguracionDisponibilidad(c *fiber.Ctx) error {
	var configuraciones []configuracion_disponibilidad.ConfiguracionDisponibilidad
	if err := c.BodyParser(&configuraciones); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}
	repo := configuracion_disponibilidad.NewConfiguracionDisponibilidadRepositoryDB()
	err := repo.ActualizarConfiguracionDisponibilidad(c.Context(), configuraciones)
	if err != nil {
		logs.Logger.Println("Error al actualizar la configuracion de disponibilidad: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al actualizar la configuracion de disponibilidad", nil))
	}
	return c.Status(fiber.StatusOK).JSON(models.Succes("Configuraciones de disponibilidad actualizadas con éxito", nil))
}
