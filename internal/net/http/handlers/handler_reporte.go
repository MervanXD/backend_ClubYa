package handlers

import (
	"github.com/MervanXD/backend_ClubYa/internal/api/models"
	"github.com/MervanXD/backend_ClubYa/internal/models/academia"
	"github.com/MervanXD/backend_ClubYa/internal/models/membresia"
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/reports"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

func GeneraReporteMembresias(c *fiber.Ctx) error {
	membresias, err := membresia.ListarMembresias()
	if err != nil {
		return c.Status(500).SendString("Error al obtener las membresias: " + err.Error())
	}

	pdfBytes, err := reports.GeneraReporteMembresias(membresias)
	if err != nil {
		return c.Status(500).SendString("Error generating report: " + err.Error())
	}

	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", "inline; filename=reporte_membresias.pdf")
	return c.Send(pdfBytes)
}

func GeneraReporteAcademias(c *fiber.Ctx) error {
	var filtros academia.ReporteAcademiaRequest

	// Parsear el body con los filtros
	if err := c.BodyParser(&filtros); err != nil {
		logs.Logger.Println("Error al parsear filtros del reporte:", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.BadRequest("Filtros inválidos", nil))
	}

	// Validaciones básicas
	if filtros.Anio <= 0 || filtros.MesInicio < 1 || filtros.MesInicio > 12 ||
		filtros.MesFin < 1 || filtros.MesFin > 12 || filtros.MesInicio > filtros.MesFin {
		return c.Status(fiber.StatusBadRequest).JSON(models.BadRequest("Filtros de fecha inválidos", nil))
	}

	// Obtener datos del reporte
	repo := academia.NewAcademiaRepositoryDB()
	academias, err := repo.GenerarReporteAcademias(filtros)
	if err != nil {
		logs.Logger.Println("Error al obtener datos del reporte:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al generar el reporte", nil))
	}

	// Generar PDF
	pdfBytes, err := reports.GeneraReporteAcademias(academias)
	if err != nil {
		logs.Logger.Println("Error al generar PDF del reporte:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error("Error al generar el PDF", nil))
	}

	// Configurar headers de respuesta
	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", "inline; filename=reporte_academias.pdf")

	return c.Send(pdfBytes)
}
