package handlers

import (
	"fmt"

	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/cuenta"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

func CrearCuenta(c *fiber.Ctx) error {
	var cuentaDTO cuenta.Cuenta
	if err := c.BodyParser(&cuentaDTO); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}
	if err := cuenta.CrearCuenta(cuentaDTO); err != nil {
		logs.Logger.Println("Error al insertar la cuenta: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al insertar la cuenta", nil))
	}
	return c.Status(fiber.StatusCreated).JSON(models.Succes("Cuenta creada con éxito", nil))
}

func LogIn(c *fiber.Ctx) error {
	var cuentaDTO cuenta.Cuenta
	if err := c.BodyParser(&cuentaDTO); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}
	var cuentaMandar cuenta.DTOCuenta
	cuentaMandar, err := cuenta.LogIn(cuentaDTO)
	if err != nil {
		logs.Logger.Println("Error al iniciar sesión: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al iniciar sesión", nil))
	}
	return c.Status(fiber.StatusOK).JSON(models.Succes("Inicio de sesión exitoso", cuentaMandar))
}

func CrearCuentaDummy(idTitular int) error {
	var cuentaDTO cuenta.Cuenta
	cuentaDTO.IdPersona = idTitular
	cuentaDTO.Username = "dummy"
	cuentaDTO.Contrasena = "dummy"
	cuentaDTO.Email = "dummy@gmail.com"
	err := cuenta.CrearCuenta(cuentaDTO)
	if err != nil {
		logs.Logger.Println("Error al crear cuenta dummy: ", err)
		return fmt.Errorf("error al crear cuenta dummy")
	}
	return nil
}
