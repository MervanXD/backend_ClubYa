package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Horario struct {
	Id         int    `json:"id"`
	HoraInicio string `json:"horaInicio"`
	HoraFin    string `json:"horaFin"`
	Disponible bool   `json:"disponible"`
	Fecha      string `json:"fecha"`
}

func main() {
	// esto es para crear una app y tener handlers y eso
	app := fiber.New()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Log file
	logFile, err := os.OpenFile("logs/app.log", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("No se pudo abrir el archivo de log: ", err)
	}
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	log.SetFlags(log.LstdFlags | log.Lshortfile) // Fecha y Hora y linea de origen

	PORT := os.Getenv("DB_PORT")
	PASS := os.Getenv("DB_PASS")
	HOST := os.Getenv("DB_HOST")
	USER := os.Getenv("DB_USER")
	NAME := os.Getenv("DB_NAME")

	llave := sha256.Sum256([]byte("[Code Here]"))
	llaveFin := llave[:]

	portDecrypt, err := decryptAES(PORT, string(llaveFin))
	passDecrypt, err := decryptAES(PASS, string(llaveFin))
	hostDecrypt, err := decryptAES(HOST, string(llaveFin))
	userDecrypt, err := decryptAES(USER, string(llaveFin))
	nameDecrypt, err := decryptAES(NAME, string(llaveFin))

	// fmt.Println("Desencriptado HOST:", hostDecrypt)
	// fmt.Println("Desencriptado PORT:", portDecrypt)
	// fmt.Println("Desencriptado USER:", userDecrypt)
	// fmt.Println("Desencriptado PASS:", passDecrypt)
	// fmt.Println("Desencriptado NAME:", nameDecrypt)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		userDecrypt, passDecrypt, hostDecrypt, portDecrypt, nameDecrypt)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error abriendo la conexion: %v", err)
	}
	defer db.Close()
	log.Println("🟢  Conexion completa....")

	if err := db.Ping(); err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	log.Println("🟢  Conectado a la base de datos")
	fmt.Println("Conectado a la base de datos :D")

	// hacemos el handler para el Horario
	app.Post("/horario", func(c *fiber.Ctx) error {
		var h Horario
		if err := c.BodyParser(&h); err != nil {
			log.Println("Error al hacer el parse JSON:", err)
			return c.Status(400).SendString("JSON inválido")
		}
		query := `INSERT INTO horario (horaInicio,horaFin,disponible,fecha) VALUES (?,?,?,?)`
		_, err := db.Exec(query, h.HoraInicio, h.HoraFin, h.Disponible, h.Fecha)
		if err != nil {
			log.Println("Error al insertar en la base de datos: ", err)
			return c.Status(500).SendString("Error al insertar horario")
		}

		log.Println("✅ Horario insertado con éxito")
		return c.SendString("Horario insertado")
	})

	app.Get("/horario", func(c *fiber.Ctx) error {
		rows, err := db.Query("SELECT id,horaInicio,horaFin,disponible, fecha FROM horario")
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

func decryptAES(encryptBase64, key string) (string, error) {
	ciphertext, _ := base64.StdEncoding.DecodeString(encryptBase64)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func encryptAES(plaintext, key string) string {
	block, _ := aes.NewCipher([]byte(key))
	aesGCM, _ := cipher.NewGCM(block)
	nonce := make([]byte, aesGCM.NonceSize())
	_, _ = rand.Read(nonce)
	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext)
}
