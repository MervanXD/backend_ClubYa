package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/pago"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type PagoRequest struct {
	IDMembresia int    `json:"idMembresia"`
	Concepto    string `json:"concepto"`
	Metodo      string `json:"metodo"` // "Tarjeta" o "Voucher"
}

func RegistrarPago(c *fiber.Ctx) error {
	var data PagoRequest

	if err := c.BodyParser(&data); err != nil {
		logs.Logger.Println("Error al parsear datos: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Datos inválidos", nil))
	}

	idPago, err := pago.InsertarPagoPorMembresia(data.IDMembresia, data.Concepto, data.Metodo)
	if err != nil {
		logs.Logger.Println("Error al insertar pago: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("No se pudo insertar el pago", nil))
	}

	return c.Status(fiber.StatusCreated).JSON(models.Succes("Pago creado exitosamente", idPago))
}
