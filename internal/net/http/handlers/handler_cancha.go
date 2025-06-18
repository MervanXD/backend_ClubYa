package handlers

import (
	"strconv"

	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/espacio"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

func ListarCanchasHorarios(c *fiber.Ctx) error {
	repo := espacio.NewCanchaRepositoryDB()
	horariosCanchas, err := repo.ObtenerCanchasHorarios()
	if err != nil {
		logs.Logger.Println("Error al obtener las lozas deportivas: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener las lozas deportivas", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Horarios de lozas deportivas obtenidos con éxito", horariosCanchas))
}

func InsertarCancha(c *fiber.Ctx) error {
	var cancha espacio.Cancha
	if err := c.BodyParser(&cancha); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}

	repo := espacio.NewCanchaRepositoryDB()
	if err := repo.InsertarCancha(cancha); err != nil {
		logs.Logger.Println("Error al insertar la cancha: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al insertar la cancha", nil))
	}

	return c.Status(fiber.StatusCreated).JSON(models.Succes("Cancha insertada con éxito", nil))
}

func ActualizarCancha(c *fiber.Ctx) error {
	var dto espacio.CanchaUpdateDTO
	if err := c.BodyParser(&dto); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}

	idStr := c.Params("id")
	if idStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID de la cancha es requerido", nil))
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Logger.Println("Error al convertir el ID de la cancha: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID de la cancha inválido", nil))
	}
	repo := espacio.NewCanchaRepositoryDB()
	if err := repo.ActualizarParcial(id, dto); err != nil {
		logs.Logger.Println("Error al actualizar la cancha: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al actualizar la cancha", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Cancha actualizada con éxito", nil))
}


func ListarCanchasConfiguracion(c *fiber.Ctx) error {
	repo := espacio.NewCanchaRepositoryDB()
	canchas, err := repo.ObtenerCanchasConfiguracion()
	if err != nil {
		logs.Logger.Println("Error al obtener las canchas: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener las canchas", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Canchas obtenidos con éxito", canchas))
}
