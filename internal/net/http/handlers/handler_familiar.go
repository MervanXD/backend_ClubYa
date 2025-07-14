package handlers

import (
	"strconv"

	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/persona"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

func InsertarFamiliares(c *fiber.Ctx) error {
	var req persona.FamiliarResquest
	if err := c.BodyParser(&req); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}
	repo := persona.NewFamiliarRepositoryDB()
	idSolicitud, err := repo.RegistrarFamiliares(req)
	if err != nil {
		logs.Logger.Println("Error al registrar familiares: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al registrar familiares", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Familiares registrados correctamente", idSolicitud))
}

func ObtenerFamiliaresPorTitular(c *fiber.Ctx) error {
	idTitularStr := c.Params("id")
	idTitular, err := strconv.Atoi(idTitularStr)
	if err != nil {
		logs.Logger.Println("ID inválido:", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID inválido", nil))
	}
	repo := persona.NewFamiliarRepositoryDB()
	familiares, err := repo.ObtenerIdsFamiliaresPorTitular(idTitular)
	if err != nil {
		logs.Logger.Println("Error al obtener familiares:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("No se pudo obtener los familiares", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Familiares obtenidos correctamente", familiares))
}

func ObtenerFamiliarPorIdPersona(c *fiber.Ctx) error {
	idPersonStr := c.Params("id")
	idPersona, err := strconv.Atoi(idPersonStr)
	if err != nil {
		logs.Logger.Println("ID de persona inválido:", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID invalido", nil))
	}
	repo := persona.NewFamiliarRepositoryDB()
	familiar, err := repo.ObtenerFamiliarPorIDPersona(idPersona)
	if err != nil {
		logs.Logger.Println("Error al obtener los datos del familiar:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("No se pudo obtener los datos del familiar", nil))
	}
	return c.Status(fiber.StatusOK).JSON(models.Succes("Familiar obtenido correctamente", familiar))
}

func RegistrarFamiliarConSolicitud(c *fiber.Ctx) error {
	var familiar persona.Familiar

	if err := c.BodyParser(&familiar); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}

	idTitular, err := strconv.Atoi(c.Params("id"))
	if err != nil || idTitular == 0 {
		logs.Logger.Println("idTitular inválido o no proporcionado en la URL")
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("idTitular inválido o no proporcionado en la URL", nil))
	}

	repo := persona.NewFamiliarRepositoryDB()
	err = repo.RegistrarFamiliarConSolicitud(familiar, idTitular)
	if err != nil {
		logs.Logger.Println("Error al registrar familiar con solicitud: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al registrar familiar con solicitud", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Familiar y solicitud registrados correctamente", nil))
}

func CrearSolicitudRetiro(c *fiber.Ctx) error {
	idFamiliar, err := strconv.Atoi(c.Params("id"))
	if err != nil || idFamiliar <= 0 {
		logs.Logger.Println("idFamiliar inválido o no proporcionado en la URL")
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("idFamiliar inválido o no proporcionado en la URL", nil))
	}

	var req struct {
		Motivo string `json:"motivo"`
	}
	
	if err := c.BodyParser(&req); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}

	repo := persona.NewFamiliarRepositoryDB()
	err = repo.CrearSolicitudRetiro(idFamiliar, req.Motivo)
	if err != nil {
		logs.Logger.Println("Error al crear solicitud de retiro: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al crear solicitud de retiro", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Solicitud de retiro creada correctamente", nil))
}
