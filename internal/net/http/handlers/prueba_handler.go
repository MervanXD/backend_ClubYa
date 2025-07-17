package handlers

import (
	"strconv"

	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/prueba"
	"github.com/gofiber/fiber/v2"
)

var repo = prueba.NewPruebaRepositoryDB()

func ListarPruebas(c *fiber.Ctx) error {
	lista, err := repo.ListarPruebas()
	if err != nil {
		return c.Status(500).JSON(models.Error("Error al listar pruebas", nil))
	}
	return c.JSON(models.Succes("Lista de pruebas obtenida", lista))
}

func CrearPrueba(c *fiber.Ctx) error {
	var req prueba.Prueba
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(models.Error("Datos inválidos", nil))
	}
	id, err := repo.InsertarPrueba(req)
	if err != nil {
		return c.Status(500).JSON(models.Error("Error al crear prueba", nil))
	}
	return c.Status(201).JSON(models.Succes("Prueba creada", id))
}

func ModificarPrueba(c *fiber.Ctx) error {
	var req prueba.Prueba
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(models.Error("Datos inválidos", nil))
	}
	if err := repo.ModificarPrueba(req); err != nil {
		return c.Status(500).JSON(models.Error("Error al modificar prueba", nil))
	}
	return c.JSON(models.Succes("Prueba modificada", nil))
}

func EliminarPrueba(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(400).JSON(models.Error("ID inválido", nil))
	}
	if err := repo.EliminarPrueba(id); err != nil {
		return c.Status(500).JSON(models.Error("Error al eliminar prueba", nil))
	}
	return c.JSON(models.Succes("Prueba eliminada", nil))
}
