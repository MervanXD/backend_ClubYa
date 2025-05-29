package handlers

import (
	"database/sql"
	"strconv"

	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/reserva"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

type EstadoReservaEspacioRequest struct {
	IdEspacio      int `json:"id_espacio"`
	IdHorarioDia   int `json:"id_horario_dia"`
	IdBloqueTiempo int `json:"id_bloque_tiempo"`
}

type AnulacionReservaRequest struct {
	IdReserva      int    `json:"id_reserva"`
	IdEspacio      int    `json:"id_espacio"`
	IdHorarioDia   int    `json:"id_horario_dia"`
	IdBloqueTiempo int    `json:"id_bloque_tiempo"`
	Motivo         string `json:"motivo"`
}

func ReservarEspacio(c *fiber.Ctx) (err error) {
	var body EstadoReservaEspacioRequest
	if err := c.BodyParser(&body); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}

	err = reserva.ReservarEspacio(body.IdEspacio, body.IdHorarioDia, body.IdBloqueTiempo)
	if err != nil {
		logs.Logger.Println("Error al actualizar estado: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("No se pudo actualizar el estado", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Estado actualizado correctamente", nil))
}

func ReservarEspacioSocial(c *fiber.Ctx) error {
	var reservaEspacioSocial reserva.ReservaEspacio

	if err := c.BodyParser(&reservaEspacioSocial); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}
	if err := reserva.ReservarEspacioSocial(reservaEspacioSocial); err != nil {
		logs.Logger.Println("Error al reservar el espacio social: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al reservar el espacio social", nil))
	}

	return c.Status(fiber.StatusCreated).JSON(models.Succes("Espacio social reservado con exito", nil))
}

func AnularReservarEspacioSocial(c *fiber.Ctx) error {
	var reservaAnulada AnulacionReservaRequest

	if err := c.BodyParser(&reservaAnulada); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}

	err := reserva.AnulacionReservaEspacioSocial(reservaAnulada.IdReserva, reservaAnulada.IdEspacio, reservaAnulada.IdHorarioDia,
		reservaAnulada.IdBloqueTiempo, reservaAnulada.Motivo)
	if err != nil {
		logs.Logger.Println("Error al anular la reserva: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("No se pudo anular la reserva", nil))
	}
	return c.Status(fiber.StatusCreated).JSON(models.Succes("Espacio social reservado anulado con exito", nil))
}

func ListarEspaciosSocialesSocio(c *fiber.Ctx) error {
	idStr := c.Params("idSocio")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Logger.Println("ID inválido: ", err)
		return c.Status(fiber.StatusBadRequest).SendString("ID inválido")
	}
	espaciosReserva, err := reserva.ObtenerReservasEspaciosSocialesSocio(id)
	if err != nil {
		if err == sql.ErrNoRows {
			logs.Logger.Println("No se encontraron reservas con ese ID", err)
			return c.Status(fiber.StatusNotFound).JSON(models.NotFound("No encontrado"))
		}
		logs.Logger.Println("Error al obtener los espacios sociales del socio: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener las reservas del socio", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Se logró obtener el espacio social", espaciosReserva))
}
