package handlers

import (
	"strconv"

	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/tarifas"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

func ListarTarifasMoras(c *fiber.Ctx) error {
	repo := tarifas.NewTarifaMoraRespositoryDB()
	moras, err := repo.ListarTarifaMora()
	if err != nil {
		logs.Logger.Println("Error al listar las moras: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al listar las moras", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Moras obtenidas con éxito", moras))
}

func ObtenerTarifaMoraPorId(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		logs.Logger.Println("ID inválido: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID inválido", nil))
	}

	repo := tarifas.NewTarifaMoraRespositoryDB()
	mora, err := repo.ObtenerTarifaMoraPorId(int64(id))
	if err != nil {
		logs.Logger.Println("Error al obtener la mora: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener la mora", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Mora obtenida con éxito", mora))
}

func ModificarMora(c *fiber.Ctx) error {
	var mora tarifas.TarifaMora
	if err := c.BodyParser(&mora); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}

	repo := tarifas.NewTarifaMoraRespositoryDB()

	if err := repo.ModificarTarifaMora(mora); err != nil {
		logs.Logger.Printf("Error al actualizar la mora (ID: %d): %v\n", mora.IdTarifaMora, err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al actualizar la mora", nil))
	}
	return c.Status(fiber.StatusOK).JSON(models.Succes("Mora actualizada con éxito", nil))
}

func InsertarMora(c *fiber.Ctx) error {
	var mora tarifas.TarifaMora
	if err := c.BodyParser(&mora); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}

	repo := tarifas.NewTarifaMoraRespositoryDB()
	if err := repo.InsertarTarifaMora(mora); err != nil {
		logs.Logger.Println("Error al insertar la mora: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al insertar la mora", nil))
	}

	return c.Status(fiber.StatusCreated).JSON(models.Succes("Mora insertada con éxito", nil))
}
