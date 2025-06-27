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
	IdEspacio    int `json:"id_espacio"`
	IdHorarioDia int `json:"id_horario_dia"`
}

type AnulacionReservaRequest struct {
	IdReserva    int    `json:"id_reserva"`
	IdEspacio    int    `json:"id_espacio"`
	IdHorarioDia int    `json:"id_horario_dia"`
	Motivo       string `json:"motivo"`
}

func ReservarEspacio(c *fiber.Ctx) (err error) {
	var body reserva.ReservaEspacio
	if err := c.BodyParser(&body); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}
	ctx := c.Context()
	repo := reserva.NewReservaRepositoryDB()
	err = repo.ReservarEspacio(ctx, body)
	if err != nil {
		logs.Logger.Println("Error al reservar el espacio: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("No se pudo reservar el espacio", nil))
	}
	return c.Status(fiber.StatusOK).JSON(models.Succes("Espacio reservado correctamente", nil))
}

func AnularReservarEspacio(c *fiber.Ctx) error {
	var reservaAnulada AnulacionReservaRequest

	if err := c.BodyParser(&reservaAnulada); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}
	repo := reserva.NewReservaRepositoryDB()
	err := repo.AnulacionReservaEspacio(reservaAnulada.IdReserva, reservaAnulada.IdEspacio, reservaAnulada.IdHorarioDia, reservaAnulada.Motivo)
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
	repo := reserva.NewReservaRepositoryDB()
	espaciosReserva, err := repo.ObtenerReservasEspaciosSocialesSocio(id)
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

func ListarCanchasSocio(c *fiber.Ctx) error {
	idStr := c.Params("idSocio")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Logger.Println("ID inválido: ", err)
		return c.Status(fiber.StatusBadRequest).SendString("ID inválido")
	}
	repo := reserva.NewReservaRepositoryDB()
	canchaReserva, err := repo.ObtenerReservasCanchasSocio(id)
	if err != nil {
		if err == sql.ErrNoRows {
			logs.Logger.Println("No se encontraron reservas con ese ID", err)
			return c.Status(fiber.StatusNotFound).JSON(models.NotFound("No encontrado"))
		}
		logs.Logger.Println("Error al obtener los espacios sociales del socio: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener las reservas del socio", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Se logró obtener la loza deportiva", canchaReserva))
}

func ListarReservasPorEspacio(c *fiber.Ctx) error {
	idStr := c.Params("idEspacio")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Logger.Println("ID inválido: ", err)
		return c.Status(fiber.StatusBadRequest).SendString("ID inválido")
	}
	repo := reserva.NewReservaRepositoryDB()
	reservas, err := repo.ObtenerReservasPorEspacio(id)
	if err != nil {
		logs.Logger.Println("Error al obtener las reservas del espacio: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener las reservas del espacio", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Se logró obtener las reservas del espacio", reservas))
}

func AceptarDevolucionAnulacionReserva(c *fiber.Ctx) error {
	var anulacionReserva reserva.AnulacionReservaRequest

	if err := c.BodyParser(&anulacionReserva); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}
	repo := reserva.NewReservaRepositoryDB()
	if err := repo.AceptarDevolucionAnulacionReserva(anulacionReserva); err != nil {
		logs.Logger.Println("Error al aceptar la devolución de la anulación de reserva: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("No se pudo aceptar la devolución de la anulación de reserva", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Devolución de anulación aceptada correctamente", nil))
}
