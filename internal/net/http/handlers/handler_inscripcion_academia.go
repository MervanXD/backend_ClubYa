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
	MetodoPago string              `json:"metodoPago"`
	MontoTotal float64             `json:"monto_total"`
	IdTitular  int                 `json:"id_titular"`
	Inscritos  []InscritosAcademia `json:"inscritos"`
}

func RegistrarInscripcionAcademia(c *fiber.Ctx) error {
	var requests InscripcionAcademiaRequest

	if err := c.BodyParser(&requests); err != nil {
		logs.Logger.Println("Error al parsear el body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.BadRequest("Datos inválidos", nil))
	}
	repo := inscripcionacademia.NewInscripcionAcademiaRepositoryDB()
	for _, req := range requests.Inscritos {
		if err := repo.RegistrarInscripcionAcademia(req.IDPersona, req.IDGrupo, req.IDTarifa, req.Uniforme, requests.MontoTotal, requests.IdTitular, requests.MetodoPago); err != nil {
			logs.Logger.Printf("Error al registrar inscripción para persona %d: %v", req.IDPersona, err)
			return c.Status(fiber.StatusInternalServerError).JSON(models.Error("No se pudo registrar la inscripción para una o más personas en academias", nil))
		}
	}
	return c.Status(fiber.StatusOK).JSON(models.Succes("Todas las inscripciones en academias fueron registradas exitosamente", nil))
}

func ListarAcademias(c *fiber.Ctx) error {
	repo := academia.NewAcademiaRepositoryDB()
	academiasDeportivas, err := repo.ObtenerAcademias()
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
	repo := academia.NewAcademiaRepositoryDB()
	academia, err := repo.ObtenerAcademiaPorId(id)
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
	repo := inscripcionacademia.NewInscripcionAcademiaRepositoryDB()
	academia, err := repo.ObtenerFamiliaresInscritosAcademia(id)
	if err != nil {
		logs.Logger.Println("Error al obtener los familiares inscritos a las academias deportivas: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener los familiares inscritos", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Academias deportivas del socio obtenidas con exito", academia))
}

func AnularInscripcionAcademia(c *fiber.Ctx) error {
	var anulacionAcademia AnulacionAcademiaRequest

	if err := c.BodyParser(&anulacionAcademia); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}
	repo := inscripcionacademia.NewInscripcionAcademiaRepositoryDB()
	if err := repo.AnularInscripcionAcademia(anulacionAcademia.IdInscripcion, anulacionAcademia.IdPersona, anulacionAcademia.Motivo); err != nil {
		logs.Logger.Println("Error al anular la inscripción a la academia: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al anular la inscripción a la academia", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Inscripción a la academia anulada con éxito", nil))
}

type AnulacionAcademiaRequest struct {
	IdInscripcion int    `json:"id_inscripcion"`
	IdPersona     int    `json:"id_persona"`
	Motivo        string `json:"motivo"`
}
