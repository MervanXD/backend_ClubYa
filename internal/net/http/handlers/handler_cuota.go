package handlers

import (
	"strconv"

	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/cuota"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

func ObtenerCuotasPorMembresia(c *fiber.Ctx) error {
	idStr := c.Params("id")
	limiteStr := c.Params("limite")

	logs.Logger.Println("ID recibido:", idStr)
	logs.Logger.Println("Limite recibido:", limiteStr)


	idMembresia, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Logger.Println("ID de membresía inválido:", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID de membresía inválido", nil))
	}

	numLimite, err := strconv.Atoi(limiteStr)
	if err != nil || numLimite <= 0 {
		logs.Logger.Println("Límite inválido:", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Límite inválido", nil))
	}

	cuotas, err := cuota.ObtenerCuotasPorMembresiaLimitado(idMembresia, numLimite)
	if err != nil {
		logs.Logger.Println("Error al obtener cuotas:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("No se pudieron obtener las cuotas", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Cuotas obtenidas correctamente", cuotas))
}
