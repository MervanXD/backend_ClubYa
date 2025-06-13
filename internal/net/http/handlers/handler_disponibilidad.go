package handlers

import (
	"context"
	"database/sql"
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

	espacio, err := detalledisponibilidad.ObtenerDisponibilidadEspacioSocialPorId(ctx, idEspacio, idHorarioDia, idBloqueTiempo)
	if err != nil {
		logs.Logger.Println("Error al obtener el espacio social: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener el espacio social", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Se logró obtener el espacio social", espacio))
}
func ListarDisponibilidadEspacio(c *fiber.Ctx) error {
	idStr := c.Params("idEspacio")
	fecha := c.Params("fecha")

	idEspacio, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Logger.Println("ID de espacio inválido: ", err)
		return c.Status(fiber.StatusBadRequest).SendString("ID de espacio inválido")
	}

	rangos, err := detalledisponibilidad.ObtenerRangosInicioDisponibles(idEspacio, fecha)
	if err != nil {
		if err == sql.ErrNoRows {
			logs.Logger.Println("No se encontraron rangos")
			return c.Status(fiber.StatusNotFound).JSON(models.NotFound("No hay horarios disponibles"))
		}
		logs.Logger.Println("Error al obtener rangos de inicio: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener horarios", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Horarios disponibles encontrados", rangos))
}
