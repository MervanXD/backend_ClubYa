package handlers

import (
	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/cuenta"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

func CrearCuenta(c *fiber.Ctx) error {
	var cuentaDTO cuenta.Cuenta
	if err := c.BodyParser(&cuentaDTO); err != nil {
		logs.Logger.Fatal("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}
	if err := cuenta.CrearCuenta(cuentaDTO); err != nil {
		logs.Logger.Fatal("Error al insertar la cuenta: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al insertar la cuenta", nil))
	}
	return c.Status(fiber.StatusCreated).JSON(models.Succes("Cuenta creada con éxito", nil))
}

func LogIn(c *fiber.Ctx) error {
	var cuentaDTO cuenta.Cuenta
	if err := c.BodyParser(&cuentaDTO); err != nil {
		logs.Logger.Fatal("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}
	var cuentaMandar cuenta.DTOCuenta
	cuentaMandar, err := cuenta.LogIn(cuentaDTO)
	if err != nil {
		logs.Logger.Fatal("Error al iniciar sesión: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al iniciar sesión", nil))
	}
	return c.Status(fiber.StatusOK).JSON(models.Succes("Inicio de sesión exitoso", cuentaMandar))
}
