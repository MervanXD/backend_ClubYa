package handlers

import (
	"strconv"

	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	solicitud "github.com/MervanXD/backend_ClubYa/internal/models/solicitud_membresia"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

func ListarSolicitudMembresia(c *fiber.Ctx) error {
	solicitudes, err := solicitud.ObtenerSolicitudesMembresia()
	if err != nil {
		logs.Logger.Fatal("Error al obtener las solicitudes de membresia: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener las solicitudes de membresia", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Solicitudes de membresia obtenidas con éxito", solicitudes))
}

type actualizarEstadoRequest struct {
	EstadoSolicitud string `json:"estado_solicitud"`
}

func ActualizarEstadoSolicitud(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		logs.Logger.Println("ID inválido: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID inválido", nil))
	}

	var body actualizarEstadoRequest
	if err := c.BodyParser(&body); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}

	err = solicitud.ActualizarEstadoSolicitud(id, body.EstadoSolicitud)
	if err != nil {
		logs.Logger.Println("Error al actualizar estado: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("No se pudo actualizar el estado", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Estado actualizado correctamente", nil))
}

func DatosSolicitudId(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		logs.Logger.Println("ID inválido: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID inválido", nil))
	}

	solicitud, err := solicitud.ObtenerDatosSolicitudPorId(id)
	if err != nil {
		logs.Logger.Fatal("Error al obtener la informacion de la solicitud", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener la informacion de la solicitud", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Informacion obtenida con exito", solicitud))
}

func ListarFamiliaresSolicitudId(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		logs.Logger.Println("ID inválido: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID inválido", nil))
	}

	familiares, err := solicitud.ObtenerFamiliaresPorIdSolicitud(id)
	if err != nil {
		logs.Logger.Fatal("Error al obtener la informacion del familiar", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener la informacion de los familiares ", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Informacion de los familiares obtenida con exito", familiares))
}

func ObtenerPersonaPorSolicitudId(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		logs.Logger.Println("ID inválido: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID inválido", nil))
	}

	persona, err := solicitud.ObtenerDatosPersonaPorIdSolicitud(id)
	if err != nil {
		logs.Logger.Fatal("Error al obtener la informacion de la persona", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener la informacion de la persona", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Informacion obtenida con exito", persona))
}
