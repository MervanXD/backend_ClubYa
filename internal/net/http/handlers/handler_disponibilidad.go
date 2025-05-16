package handlers

import (
	"database/sql"
	"strconv"

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
	espacio, err := detalledisponibilidad.ObtenerDisponibilidadEspacioSocialPorId(c.Context(), idEspacio, idHorarioDia, idBloqueTiempo)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(models.NotFound("No encontrado"))
		}
		logs.Logger.Fatal("Error al obtener el espacio social: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener el espacio social", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Se logró obtener el espacio social", espacio))
}
