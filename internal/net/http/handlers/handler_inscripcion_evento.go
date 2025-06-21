package handlers

import (
	"strconv"

	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/inscripcion_evento"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

type InscripcionRequest struct {
	FidPersona        int `json:"id_persona" validate:"required"`
	IdEvento          int `json:"id_evento" validate:"required"`
	CantidadInvitados int `json:"cantidad_invitados"`
}

type AnulacionRequest struct {
	IdInscripcionEvento int    `json:"id_inscripcion_evento" validate:"required"`
	Motivo              string `json:"motivo"`
}

type PagoInscripcionRequest struct {
	IdTitular   int     `json:"id_titular" validate:"required"`
	IdInscritos []int   `json:"id_inscritos" validate:"required"`
	IdEvento    int     `json:"id_evento" validate:"required"`
	Monto       float64 `json:"monto" validate:"required"`
	Concepto    string  `json:"concepto" validate:"required"`
	MetodoPago  string  `json:"metodo" validate:"required"`
}

func RegistrarInscripcionEvento(c *fiber.Ctx) error {
	var requests []InscripcionRequest

	if err := c.BodyParser(&requests); err != nil {
		logs.Logger.Println("Error al parsear el body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.BadRequest("Datos inválidos", nil))
	}

	if len(requests) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.BadRequest("La lista de inscripciones no puede estar vacía", nil))
	}
	repo := inscripcion_evento.NewInscripcionEventoRepositoryDB()
	for _, req := range requests {
		if err := repo.RegistrarInscripcion(req.FidPersona, req.IdEvento, req.CantidadInvitados); err != nil {
			logs.Logger.Printf("Error al registrar inscripción para persona %d: %v", req.FidPersona, err)
			return c.Status(fiber.StatusInternalServerError).JSON(models.Error("No se pudo registrar la inscripción para una o más personas", nil))
		}
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Todas las inscripciones fueron registradas exitosamente", nil))
}

func ListarEventosSocioId(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Logger.Println("ID inválido: ", err)
		return c.Status(fiber.StatusBadRequest).SendString("ID inválido")
	}
	repo := inscripcion_evento.NewInscripcionEventoRepositoryDB()
	eventos, err := repo.ObtenerEventosSocio(id)
	if err != nil {
		logs.Logger.Println("Error al obtener los eventos del socio: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener los eventos del socio", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Eventos del socio obtenidas con exito", eventos))
}

func AnularInscripcionEvento(c *fiber.Ctx) error {
	var request AnulacionRequest

	if err := c.BodyParser(&request); err != nil {
		logs.Logger.Println("Error al parsear el body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.BadRequest("Datos inválidos", nil))
	}

	if request.IdInscripcionEvento == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.BadRequest("ID de inscripcion no puede ser cero", nil))
	}

	repo := inscripcion_evento.NewInscripcionEventoRepositoryDB()
	if err := repo.AnularInscripcion(request.IdInscripcionEvento, request.Motivo); err != nil {
		logs.Logger.Printf("Error al anular inscripción %d: %v", request.IdInscripcionEvento, err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("No se pudo anular la inscripción", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Inscripción anulada exitosamente", nil))
}

func PagarInscripcionEvento(c *fiber.Ctx) error {
	var request PagoInscripcionRequest

	if err := c.BodyParser(&request); err != nil {
		logs.Logger.Println("Error al parsear el body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.BadRequest("Datos inválidos", nil))
	}

	if request.IdTitular == 0 || request.IdEvento == 0 || request.MetodoPago == "" || request.Monto <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.BadRequest("Datos de pago inválidos", nil))
	}

	repo := inscripcion_evento.NewInscripcionEventoRepositoryDB()
	//registra las inscripciones de los invitados con pago pendiente
	for _, idInscrito := range request.IdInscritos {
		if err := repo.RegistrarInscripcion(idInscrito, request.IdEvento, 0); err != nil {
			logs.Logger.Printf("Error al registrar inscripción para invitado %d: %v", idInscrito, err)
			return c.Status(fiber.StatusInternalServerError).JSON(models.Error("No se pudo registrar la inscripción para uno o más invitados", nil))
		}
	}
	//procesa el pago de la inscripción (acá se actualiza el estado de la inscripción a confirmada)
	idPago, err := repo.PagarInscripcion(request.IdTitular, request.IdEvento, request.Concepto, request.MetodoPago, request.Monto)
	if err != nil {
		logs.Logger.Printf("Error al procesar el pago de inscripción: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("No se pudo procesar el pago de inscripción", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Pago de inscripción procesado exitosamente", map[string]int{"id_pago": idPago}))
}
