package handlers

import (
	"strings"

	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/pago"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
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

func PagarTodosPendientes(c *fiber.Ctx) error {
	var request pago.PagoMultipleRequest

	// Parsear el body de la request
	if err := c.BodyParser(&request); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud:", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.BadRequest("Error al parsear el cuerpo de la solicitud", nil))
	}

	// Validaciones básicas
	if request.IdTitular <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.BadRequest("ID de titular inválido", nil))
	}

	if request.MetodoPago.String() != "Tarjeta" && request.MetodoPago.String() != "Voucher" {
		return c.Status(fiber.StatusBadRequest).JSON(models.BadRequest("Método de pago inválido. Use 'Tarjeta' o 'Voucher'", nil))
	}

	// Procesar los pagos pendientes
	idsPagos, err := pago.PagarTodosPagosPendientes(request.IdTitular, request.MetodoPago.String())
	if err != nil {
		logs.Logger.Printf("Error al procesar pagos pendientes para titular %d: %v", request.IdTitular, err)

		// Manejo específico de errores
		errMsg := err.Error()
		if strings.Contains(errMsg, "no se encontraron pagos pendientes") {
			return c.Status(fiber.StatusNotFound).JSON(models.Error("No se encontraron pagos pendientes para este titular", nil))
		}

		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al procesar los pagos pendientes", nil))
	}

	// Crear respuesta de éxito
	response := pago.PagoMultipleResponse{
		TotalPagos:     len(idsPagos),
		IdsPagos:       idsPagos,
		PagosAfectados: len(idsPagos),
	}

	logs.Logger.Printf("Pagos procesados exitosamente para titular %d. Total: %d pagos", request.IdTitular, len(idsPagos))

	return c.Status(fiber.StatusOK).JSON(models.Succes("Todos los pagos pendientes han sido procesados exitosamente", response))
}
