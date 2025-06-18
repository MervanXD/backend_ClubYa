package handlers

import (
	"database/sql"
	"strconv"

	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	bloqueTiempo "github.com/MervanXD/backend_ClubYa/internal/models/bloque_tiempo"
	detalledisponibilidad "github.com/MervanXD/backend_ClubYa/internal/models/detalle_disponibilidad"
	"github.com/MervanXD/backend_ClubYa/internal/models/espacio"
	"github.com/MervanXD/backend_ClubYa/internal/models/horario"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

func CrearEspacioSocial(c *fiber.Ctx) error {
	var espacioSocial espacio.EspacioSocial

	if err := c.BodyParser(&espacioSocial); err != nil {
		logs.Logger.Println("Error al parsear el cuerpo de la solicitud: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear el cuerpo de la solicitud", nil))
	}
	repo := espacio.NewEspacioSocialRepositoryDB()
	if err := repo.InsertarEspacioSocial(espacioSocial); err != nil {
		logs.Logger.Println("Error al insertar el espacio social: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al insertar el espacio social", nil))
	}

	return c.Status(fiber.StatusCreated).JSON(models.Succes("Espacio social creado con éxito", nil))
}

func ListarEspaciosSociales(c *fiber.Ctx) error {
	repo := espacio.NewEspacioSocialRepositoryDB()
	espacios, err := repo.ObtenerEspaciosSociales()
	if err != nil {
		logs.Logger.Println("Error al obtener los espacios sociales: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener los espacios sociales", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Espacios sociales obtenidos con éxito", espacios))
}

func ObtenerEspacioSocialPorId(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Logger.Println("ID inválido: ", err)
		return c.Status(fiber.StatusBadRequest).SendString("ID inválido")
	}
	repo := espacio.NewEspacioSocialRepositoryDB()
	espacio, err := repo.ObtenerEspacioSocialPorID(c.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			logs.Logger.Println("No se encontró ningun espacio social con ese ID", err)
			return c.Status(fiber.StatusNotFound).JSON(models.NotFound("No encontrado"))
		}
		logs.Logger.Println("Error al obtener el espacio social: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener el espacio social", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Se logró obtener el espacio social", espacio))
}

func ListarEspaciosSocialesHorarios(c *fiber.Ctx) error {
	repo := espacio.NewEspacioSocialRepositoryDB()
	horariosEspacioSocial, err := repo.ObtenerEspaciosSocialesHorarios()
	if err != nil {
		logs.Logger.Println("Error al obtener los horarios espacios sociales: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener los horarios espacios sociales", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Horarios de espacios sociales obtenidos con éxito", horariosEspacioSocial))
}

func ActualizarEspacioSocial(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		logs.Logger.Println("ID inválido: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("ID inválido", nil))
	}

	var input espacio.EspacioSocialUpdateDTO
	if err := c.BodyParser(&input); err != nil {
		logs.Logger.Println("Error al parsear JSON: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error("Error al parsear JSON", nil))
	}
	repo := espacio.NewEspacioSocialRepositoryDB()
	err = repo.ActualizarParcial(id, input)
	if err != nil {
		logs.Logger.Println("Error al actualizar el espacio social: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al actualizar", nil))
	}

	return c.JSON(models.Succes("Actualizado correctamente", nil))
}

func ListarEspaciosSocialesConfiguracion(c *fiber.Ctx) error {
	repo := espacio.NewEspacioSocialRepositoryDB()
	espacios, err := repo.ObtenerEspaciosSocialesConfiguracion()
	if err != nil {
		logs.Logger.Println("Error al obtener los espacios sociales: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener los espacios sociales", nil))
	}

	return c.Status(fiber.StatusOK).JSON(models.Succes("Espacios sociales obtenidos con éxito", espacios))
}

func ListarDisponibilidadEspacio(c *fiber.Ctx) error {
	var requests DisponibilidadEspacioRequest

	if err := c.BodyParser(&requests); err != nil {
		logs.Logger.Println("Error al parsear el body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.BadRequest("Datos inválidos", nil))
	}
	repo := horario.NewHorarioDiaRepositoryDB()
	horarios, err := repo.ObtenerInscritosEspacioFecha(requests.IdEspacio, requests.FechaActual)
	if err != nil {
		logs.Logger.Println("Error al obtener la disponibilidad del espacio social de la fecha: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener la disponibilidad del espacio social de la fecha:", nil))
	}

	repo2 := bloqueTiempo.NewBloqueTiempoRepositoryDB()
	bloques, err := repo2.ObtenerBloquesTiempoEspacio(requests.IdEspacio)
	if err != nil {
		logs.Logger.Println("Error al obtener los bloques de tiempo ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener los bloques de tiempo", nil))
	}

	repo3 := detalledisponibilidad.NewDetalleDisponibilidadRepositoryDB()
	detalles, err := repo3.ObtenerDetalleDisponibilidadEspacioFechaId(requests.IdEspacio, requests.FechaActual)
	if err != nil {
		logs.Logger.Println("Error al obtener los detalles del espacio ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al obtener los detalles del espacio", nil))
	}
	var disponibilidades DisponibilidadEspacio
	disponibilidades.HorariosDias = horarios

	disponibilidades.Bloques = bloques
	disponibilidades.DetallesDisponibilidad = detalles

	return c.Status(fiber.StatusOK).JSON(models.Succes("Todas la información de disponibilidad de el espacio fue obtenido con éxito", disponibilidades))
}

type DisponibilidadEspacioRequest struct {
	IdEspacio   int    `json:"id_espacio"`
	FechaActual string `json:"fecha_actual"`
}

type DisponibilidadEspacio struct {
	HorariosDias           []horario.HorarioDiaDTO                          `json:"horariosDia"`
	Bloques                []bloqueTiempo.BloqueTiempo                      `json:"bloquesTiempo"`
	DetallesDisponibilidad []detalledisponibilidad.DetalleDisponibilidadDto `json:"detalleDisponibilidad"`
}
