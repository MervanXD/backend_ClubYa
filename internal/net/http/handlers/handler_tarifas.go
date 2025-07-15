package handlers

import (
	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/tarifas"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

func ListarTarifasMiembrosFamiliar(c *fiber.Ctx) error {
	repo := tarifas.NewTarifaMiembroAdicionalRespositoryDB()
	tarifas, err := repo.ListarTarifaMiembroAdicional()
	if err != nil {
		logs.Logger.Println("Error al listar las tarifas por familiar: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al listar las tarifas por miembro de familia", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Tarifas obtenidas con éxito", tarifas))
}

func ModificarTarifaMiembroAdicional(c *fiber.Ctx) error {
	var tarifa TarifaMembresiaRequest
	if err := c.BodyParser(&tarifa); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}

	repo := tarifas.NewTarifaMiembroAdicionalRespositoryDB()

	if err := repo.ModificarTarifaMiembroAdicional(tarifa.IdTarifaMiembroAdicional, tarifa.Monto); err != nil {
		logs.Logger.Printf("Error al actualizar la tarifa (ID: %d, Monto: %.2f): %v\n", tarifa.IdTarifaMiembroAdicional, tarifa.Monto, err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al actualizar la tarifa", nil))
	}
	return c.Status(fiber.StatusOK).JSON(models.Succes("Tarifa actualizada con éxito", nil))
}

func InsertarTarifaMiembroAdicional(c *fiber.Ctx) error {
	var tarifa tarifas.TarifaMiembroAdicional
	if err := c.BodyParser(&tarifa); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}

	repo := tarifas.NewTarifaMiembroAdicionalRespositoryDB()
	if err := repo.InsertarTarifaMiembroAdicional(tarifa); err != nil {
		logs.Logger.Println("Error al insertar la tarifa: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al insertar la tarifa", nil))
	}

	return c.Status(fiber.StatusCreated).JSON(models.Succes("Tarifa insertada con éxito", nil))
}

type TarifaMembresiaRequest struct {
	IdTarifaMiembroAdicional int     `json:"idTarifaMiembroAdicional"`
	Monto                    float64 `json:"monto"`
}
