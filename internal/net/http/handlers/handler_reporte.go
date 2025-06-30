package handlers

import (
	"github.com/MervanXD/backend_ClubYa/internal/models/membresia"
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/reports"
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
