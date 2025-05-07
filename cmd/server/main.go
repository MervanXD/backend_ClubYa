package main

import (
	"log"

	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/gofiber/fiber/v2"
)

type Horario struct {
	Id         int    `json:"id"`
	HoraInicio string `json:"horaInicio"`
	HoraFin    string `json:"horaFin"`
	Disponible bool   `json:"disponible"`
	Fecha      string `json:"fecha"`
}

func main() {
	logs.InitLogger()
	defer logs.CloseLogger()

	// inicializamos la base de datos
	database.InitDB()
	defer database.CloseDB()

	// esto es para crear una app y tener handlers y eso
	app := fiber.New()

	// hacemos el handler para el Horario

	app.Post("/horario", func(c *fiber.Ctx) error {
		var h Horario
		if err := c.BodyParser(&h); err != nil {
			log.Println("Error al hacer el parse JSON:", err)
			return c.Status(400).SendString("JSON inválido")
		}
		query := `INSERT INTO horario (horaInicio,horaFin,disponible,fecha) VALUES (?,?,?,?)`
		_, err := database.DB.Exec(query, h.HoraInicio, h.HoraFin, h.Disponible, h.Fecha)
		if err != nil {
			log.Println("Error al insertar en la base de datos: ", err)
			return c.Status(500).SendString("Error al insertar horario")
		}

		log.Println("✅ Horario insertado con éxito")
		return c.SendString("Horario insertado")
	})

	app.Get("/horario", func(c *fiber.Ctx) error {
		rows, err := database.DB.Query("SELECT id,horaInicio,horaFin,disponible, fecha FROM horario")
		if err != nil {
			log.Println("Error al consultar horarios: ", err)
			return c.Status(500).SendString("Error al consultar horarios")
		}
		defer rows.Close()

		var horarios []Horario
		for rows.Next() {
			var h Horario
			if err := rows.Scan(&h.Id, &h.HoraInicio, &h.HoraFin, &h.Disponible, &h.Fecha); err != nil {
				log.Println("Error al leer horario: ", err)
				continue
			}
			horarios = append(horarios, h)
		}
		return c.JSON(horarios)
	})

	log.Fatal(app.Listen(":4000"))

}
