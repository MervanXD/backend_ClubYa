package handlers

import (
	"database/sql"
	"strconv"

	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/academia"
	inscripcionacademia "github.com/MervanXD/backend_ClubYa/internal/models/inscripcion_academia"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

type InscritosAcademia struct {
	IDPersona int    `json:"id_persona"`
	IDGrupo   int    `json:"id_grupo"`
	IDTarifa  int    `json:"id_tarifa"`
	Uniforme  int    `json:"uniforme"`
	TipoSocio string `json:"tipoSocio"`
}

type InscripcionAcademiaRequest struct {
	MontoTotal float64             `json:"monto_total"`
	Inscritos  []InscritosAcademia `json:"inscritos"`
}

func RegistrarInscripcionAcademia(c *fiber.Ctx) error {
	var requests InscripcionAcademiaRequest

	if err := c.BodyParser(&requests); err != nil {
		logs.Logger.Println("Error al parsear el body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.BadRequest("Datos inválidos", nil))
	}

	for _, req := range requests.Inscritos {
		if err := inscripcionacademia.RegistrarInscripcionAcademia(req.IDPersona, req.IDGrupo, req.IDTarifa, req.Uniforme); err != nil {
			logs.Logger.Printf("Error al registrar inscripción para persona %d: %v", req.IDPersona, err)
			return c.Status(fiber.StatusInternalServerError).JSON(models.Error("No se pudo registrar la inscripción para una o más personas en academias", nil))
		}
	}
	return c.Status(fiber.StatusOK).JSON(models.Succes("Todas las inscripciones en academias fueron registradas exitosamente", nil))
}

func ListarAcademias(c *fiber.Ctx) error {
	academiasDeportivas, err := academia.ObtenerAcademias()
	if err != nil {
		logs.Logger.Println("Error al obtener las academias deportivas: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener las academias", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Academias deportivas obtenidas con exito", academiasDeportivas))
}

func ObtenerAcademiaId(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Logger.Println("ID inválido: ", err)
		return c.Status(fiber.StatusBadRequest).SendString("ID inválido")
	}
	academia, err := academia.ObtenerAcademiaPorId(id)
	if err != nil {
		if err == sql.ErrNoRows {
			logs.Logger.Println("No se encontró ninguna academia con ese ID", err)
			return c.Status(fiber.StatusNotFound).JSON(models.NotFound("No encontrado"))
		}
		logs.Logger.Println("Error al obtener la academia: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener la infromacion de la academia", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Se logró obtener la informacion de la academia", academia))
}

func ListarFamiliaresSocioInscritos(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Logger.Println("ID inválido: ", err)
		return c.Status(fiber.StatusBadRequest).SendString("ID inválido")
	}
	academia, err := inscripcionacademia.ObtenerFamiliaresInscritosAcademia(id)
	if err != nil {
		logs.Logger.Println("Error al obtener los familiares inscritos a las academias deportivas: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener los familiares inscritos", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Academias deportivas del socio obtenidas con exito", academia))
}
