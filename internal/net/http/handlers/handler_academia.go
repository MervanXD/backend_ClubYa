package handlers

import (
	"net/http"

	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/academia"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

func InsertarAcademia(c *fiber.Ctx) error {
	var aca academia.Academia
	if err := c.BodyParser(&aca); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(http.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}

	repo := academia.NewAcademiaRepositoryDB()
	err := repo.InsertarAcademia(&aca)
	if err != nil {
		logs.Logger.Println("Error al insertar la academia: ", err)
		return c.Status(http.StatusInternalServerError).JSON(models.Error("Error al insertar la academia", nil))
	}

	return c.Status(http.StatusCreated).JSON(models.Succes("Academia creada con éxito", nil))
}

func ListarAcademiasGenerales(c *fiber.Ctx) error {
	repo := academia.NewAcademiaRepositoryDB()
	academias, err := repo.ListarAcademiasGenerales()
	if err != nil {
		logs.Logger.Println("Error al listar las academias: ", err)
		return c.Status(http.StatusInternalServerError).JSON(models.Error("Error al listar las academias", nil))
	}

	return c.Status(http.StatusOK).JSON(models.Succes("Lista de academias obtenida con éxito", academias))
}
