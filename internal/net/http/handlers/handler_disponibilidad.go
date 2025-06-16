package handlers

import (
	"context"
	"strconv"
	"time"

	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	detalledisponibilidad "github.com/MervanXD/backend_ClubYa/internal/models/detalle_disponibilidad"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

func ObtenerDisponibilidadEspacioSocialPorId(c *fiber.Ctx) error {
	s_idEspacio := c.Params("id_espacio")
	idEspacio, err := strconv.Atoi(s_idEspacio)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Id espacio inválido")
	}

	s_idHorarioDia := c.Params("id_horario_dia")
	idHorarioDia, err := strconv.Atoi(s_idHorarioDia)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Id horario dia inválido")
	}

	s_idBloqueTiempo := c.Params("id_bloque_tiempo")
	idBloqueTiempo, err := strconv.Atoi(s_idBloqueTiempo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Id bloque tiempo inválido")
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()
	rapo := detalledisponibilidad.NewDetalleDisponibilidadRepositoryDB()
	espacio, err := rapo.ObtenerDisponibilidadEspacioSocialPorId(ctx, idEspacio, idHorarioDia, idBloqueTiempo)
	if err != nil {
		logs.Logger.Println("Error al obtener el espacio social: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener el espacio social", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Se logró obtener el espacio social", espacio))
}

func ActualizarDetalleDisponibilidad(c *fiber.Ctx) error {
	var detalleDisponibilidad detalledisponibilidad.DetalleDisponibilidad
	if err := c.BodyParser(&detalleDisponibilidad); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}

	repo := detalledisponibilidad.NewDetalleDisponibilidadRepositoryDB()
	err := repo.ActualizarEstadoDetalleDisponibilidad(detalleDisponibilidad.IdHorarioDia, detalleDisponibilidad.IdBloqueTiempo, detalleDisponibilidad.EstadoDisponibilidad.String())
	if err != nil {
		logs.Logger.Println("Error al actualizar el detalle de disponibilidad: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al actualizar el detalle de disponibilidad", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Detalle de disponibilidad actualizado correctamente", nil))
}
