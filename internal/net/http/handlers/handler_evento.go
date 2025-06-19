package handlers

import (
	"strconv"

	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/evento"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

func ListarEventos(c *fiber.Ctx) error {
	repo := evento.NewEventoRepositoryDB()
	eventos, err := repo.ListarEventos()
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
	repo := evento.NewEventoRepositoryDB()
	eventoInfo, err := repo.BuscarEventoPorID(id)
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

func CrearEvento(c *fiber.Ctx) error {
	var eventoData evento.EventoRequest

	if err := c.BodyParser(&eventoData); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear los datos del evento", nil))
	}
	repo := evento.NewEventoRepositoryDB()
	idEvento, err := repo.InsertarEvento(eventoData)
	if err != nil {
		logs.Logger.Println("Error al insertar el evento: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al crear el evento", nil))
	}

	return c.Status(fiber.StatusCreated).JSON(models.Succes("Evento creado con éxito", idEvento))
}

func ModificarEvento(c *fiber.Ctx) error {
    var eventoData evento.EventoRequest

    idStr := c.Params("id")
    id, err := strconv.Atoi(idStr)
    if err != nil || id <= 0 {
        logs.Logger.Println("ID de evento inválido:", idStr)
        return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID de evento inválido", nil))
    }

    if err := c.BodyParser(&eventoData); err != nil {
        logs.Logger.Println("Error al parsear el cuerpo de la solicitud:", err)
        return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear los datos del evento", nil))
    }

    eventoData.IdEvento = id

	repo := evento.NewEventoRepositoryDB()
    if err := repo.ModificarEvento(eventoData); err != nil {
        logs.Logger.Println("Error al modificar el evento:", err)
        return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al modificar el evento", nil))
    }

    return c.Status(fiber.StatusOK).JSON(models.Succes("Evento modificado con éxito", nil))
}

func CancelarEvento(c *fiber.Ctx) error {
    idStr := c.Params("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        logs.Logger.Println("Error al convertir id a entero:", err)
        return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID inválido", nil))
    }

	repo := evento.NewEventoRepositoryDB()
    if err := repo.CancelarEvento(id); err != nil {
        logs.Logger.Println("Error al cancelar el evento:", err)
        return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al cancelar el evento", nil))
    }

    return c.Status(fiber.StatusOK).JSON(models.Succes("Evento cancelado con éxito", nil))
}


func EliminarEvento(c *fiber.Ctx) error {
    idStr := c.Params("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        logs.Logger.Println("Error al convertir id a entero:", err)
        return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID inválido", nil))
    }

	repo := evento.NewEventoRepositoryDB()
    if err := repo.EliminarEvento(id); err != nil {
        logs.Logger.Println("Error al eliminar el evento:", err)
        return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al cancelar el evento", nil))
    }

    return c.Status(fiber.StatusOK).JSON(models.Succes("Evento cancelado con éxito", nil))
}
