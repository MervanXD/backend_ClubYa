package routes

import (
	"github.com/MervanXD/backend_ClubYa/internal/net/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RouteCuenta(api fiber.Router) {
	api.Post("/cuenta/crear", handlers.CrearCuenta)
	api.Post("/cuenta/login", handlers.LogIn)
	api.Post("/cuenta/crear-administrador", handlers.CrearCuentaAdministrador)
	api.Get("/cuenta/administradores", handlers.ObtenerAdministradores)
	api.Get("/cuenta/perfil/:id", handlers.ObtenerPerfilPorIdCuenta)
	api.Patch("/cuenta/actualizar-perfil/:id", handlers.ActualizarCuenta)
	api.Get("/cuenta/id-por-persona/:idPersona", handlers.ObtenerIdCuentaPorPersona)
	api.Get("/cuenta/usuarios", handlers.ObtenerUsuarios)
	api.Post("/cuenta/registrar-gmail", handlers.RegistrarGmail)
	api.Post("/cuenta/login-gmail", handlers.LoginGmail)
}
