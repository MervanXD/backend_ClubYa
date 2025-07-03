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
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}
	repo := cuenta.NewCuentaRepositoryDB()
	idCuenta, err := repo.CrearCuenta(cuentaDTO)
	if err != nil {
		logs.Logger.Println("Error al insertar la cuenta: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al insertar la cuenta", nil))
	}

	return c.Status(fiber.StatusCreated).JSON(models.Succes("Cuenta creada con éxito", idCuenta))
}

func LogIn(c *fiber.Ctx) error {
	var cuentaDTO cuenta.Cuenta
	if err := c.BodyParser(&cuentaDTO); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}
	var cuentaMandar cuenta.DTOCuenta
	repo := cuenta.NewCuentaRepositoryDB()
	cuentaMandar, err := repo.LogIn(cuentaDTO)
	if err != nil {
		logs.Logger.Println("Error al iniciar sesión: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al iniciar sesión", nil))
	}
	return c.Status(fiber.StatusOK).JSON(models.Succes("Inicio de sesión exitoso", cuentaMandar))
}

func CrearCuentaAdministrador(c *fiber.Ctx) error {
	var cuentaDTO cuenta.CuentaAdminDTO
	if err := c.BodyParser(&cuentaDTO); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}
	repo := cuenta.NewCuentaRepositoryDB()
	err := repo.CrearCuentaAdministrador(cuentaDTO)
	if err != nil {
		logs.Logger.Println("Error al crear la cuenta de administrador: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al crear la cuenta de administrador", nil))
	}

	return c.Status(fiber.StatusCreated).JSON(models.Succes("Cuenta de administrador creada con éxito", nil))
}

func ObtenerAdministradores(c *fiber.Ctx) error {
	repo := cuenta.NewCuentaRepositoryDB()
	administradores, err := repo.ObtenerAdministradores()
	if err != nil {
		logs.Logger.Println("Error al obtener administradores: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener administradores", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Administradores obtenidos con éxito", administradores))
}
