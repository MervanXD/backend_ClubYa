package handlers

import (
	"strconv"
	"strings"

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
		errMsg := err.Error()
		if strings.Contains(errMsg, "ya está registrado") {
			return c.Status(fiber.StatusConflict).JSON(models.Error("La cuenta Gmail ya está registrada", nil))
		}
		if strings.Contains(errMsg, "no ha sido verificado") {
			return c.Status(fiber.StatusForbidden).JSON(models.Error("La cuenta Gmail no ha sido verificada", nil))
		}
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
		errMsg := err.Error()
		if strings.Contains(errMsg, "ya está registrado") {
			return c.Status(fiber.StatusConflict).JSON(models.Error("La cuenta Gmail ya está registrada", nil))
		}
		if strings.Contains(errMsg, "no ha sido verificado") {
			return c.Status(fiber.StatusForbidden).JSON(models.Error("La cuenta Gmail no ha sido verificada", nil))
		}
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

func ObtenerPerfilPorIdCuenta(c *fiber.Ctx) error {
	idCuentaStr := c.Params("id")
	if idCuentaStr == "" {
		logs.Logger.Println("ID de cuenta no proporcionado")
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID de cuenta no proporcionado", nil))
	}
	idCuenta, err := strconv.Atoi(idCuentaStr)
	if err != nil {
		logs.Logger.Println("Error al convertir ID de cuenta a entero: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID de cuenta inválido", nil))
	}
	idCuentaNew := int64(idCuenta)
	repo := cuenta.NewCuentaRepositoryDB()
	perfil, err := repo.ObtenerPerfilPorIdCuenta(idCuentaNew)
	if err != nil {
		logs.Logger.Println("Error al obtener perfil por ID de cuenta: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener perfil por ID de cuenta", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Perfil obtenido con éxito", perfil))
}

func ActualizarCuenta(c *fiber.Ctx) error {
	idCuentaStr := c.Params("id")
	if idCuentaStr == "" {
		logs.Logger.Println("ID de cuenta no proporcionado")
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID de cuenta no proporcionado", nil))
	}
	idCuenta, err := strconv.Atoi(idCuentaStr)
	if err != nil {
		logs.Logger.Println("Error al convertir ID de cuenta a entero: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID de cuenta inválido", nil))
	}
	idCuentaNew := int64(idCuenta)

	var cuentaDTO cuenta.CuentaAdminUpdateDTO
	if err := c.BodyParser(&cuentaDTO); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}
	repo := cuenta.NewCuentaRepositoryDB()
	err = repo.ActualizarCuentaAParcial(idCuentaNew, cuentaDTO)
	if err != nil {
		logs.Logger.Println("Error al actualizar la cuenta: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al actualizar la cuenta", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Cuenta actualizada con éxito", nil))
}

func ObtenerIdCuentaPorPersona(c *fiber.Ctx) error {
	idPersonaStr := c.Params("idPersona")
	if idPersonaStr == "" {
		logs.Logger.Println("ID de persona no proporcionado")
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID de persona no proporcionado", nil))
	}

	idPersona, err := strconv.Atoi(idPersonaStr)
	if err != nil {
		logs.Logger.Println("Error al convertir ID de persona a entero: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID de persona inválido", nil))
	}

	repo := cuenta.NewCuentaRepositoryDB()
	idCuenta, err := repo.ObtenerIdCuentaPorPersona(int64(idPersona))
	if err != nil {
		logs.Logger.Println("Error al obtener ID de cuenta por persona: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener ID de cuenta por persona", nil))
	}

	if idCuenta == 0 {
		logs.Logger.Println("No se encontró cuenta para la persona con ID:", idPersona)
		return c.Status(fiber.StatusNotFound).JSON(models.Error("No se encontró cuenta para la persona especificada", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("ID de cuenta obtenido con éxito", idCuenta))
}

func ObtenerUsuarios(c *fiber.Ctx) error {
	repo := cuenta.NewCuentaRepositoryDB()
	usuarios, err := repo.ListarUsuarios()
	if err != nil {
		logs.Logger.Println("Error al obtener usuarios: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener usuarios", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Usuarios listados con éxito", usuarios))
}

func RegistrarGmail(c *fiber.Ctx) error {
	var cuentaDTO cuenta.CuentaGmailDTO
	if err := c.BodyParser(&cuentaDTO); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}

	repo := cuenta.NewCuentaRepositoryDB()
	idCuenta, err := repo.RegistrarGmail(cuentaDTO)
	if err != nil {
		logs.Logger.Println("Error al registrar cuenta Gmail: ", err)
		errMsg := err.Error()
		if strings.Contains(errMsg, "ya está registrado") {
			return c.Status(fiber.StatusConflict).JSON(models.Error("La cuenta Gmail ya está registrada", nil))
		}
		if strings.Contains(errMsg, "no ha sido verificado") {
			return c.Status(fiber.StatusForbidden).JSON(models.Error("La cuenta Gmail no ha sido verificada", nil))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al registrar cuenta Gmail", nil))
	}

	return c.Status(fiber.StatusCreated).JSON(models.Succes("Cuenta Gmail registrada con éxito", idCuenta))
}

func LoginGmail(c *fiber.Ctx) error {
	var cuentaDTO cuenta.CuentaGmailDTO
	if err := c.BodyParser(&cuentaDTO); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}

	repo := cuenta.NewCuentaRepositoryDB()
	cuentaMandar, err := repo.LoginGmail(cuentaDTO)
	if err != nil {
		logs.Logger.Println("Error al iniciar sesión con Gmail: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al iniciar sesión con Gmail", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Inicio de sesión con Gmail exitoso", cuentaMandar))
}
