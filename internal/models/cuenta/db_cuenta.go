package cuenta

import (
	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/security"
	"github.com/MervanXD/backend_ClubYa/logs"
)

func CrearCuenta(cuenta Cuenta) error {
	usernameEncriptado := security.Hash256(cuenta.Username)
	passwordEncriptado := security.Hash256(cuenta.Contrasena)
	emailEncriptado := security.Hash256(cuenta.Email)
	query := "call ingesoft.InsertarCuenta(?, ?, ?, ?)"
	_, err := database.DB.Exec(query, usernameEncriptado, emailEncriptado, passwordEncriptado, cuenta.IdPersona)
	if err != nil {
		return err
	}
	return nil
}

func LogIn(cuenta Cuenta) (DTOCuenta, error) {
	usernameEncriptado := security.Hash256(cuenta.Username)
	passwordEncriptado := security.Hash256(cuenta.Contrasena)
	query := "call ingesoft.LogIn(?, ?,@c_fid_persona,@c_rol,@c_esPostulante)"
	_, err := database.DB.Exec(query, usernameEncriptado, passwordEncriptado)
	var cuentaDTO DTOCuenta
	cuentaDTO.Username = cuenta.Username
	cuentaDTO.Contrasena = cuenta.Contrasena
	if err != nil {
		logs.Logger.Println("Error al iniciar sesion: ", err)
		return cuentaDTO, err
	}
	err = database.DB.QueryRow("SELECT @c_fid_persona, @c_rol,@c_esPostulante").Scan(&cuentaDTO.IdPersona, &cuentaDTO.Rol, &cuentaDTO.Postulante)
	if err != nil {
		logs.Logger.Println("Error al obtener idPersona, rol y postulante: ", err)
		return cuentaDTO, err
	}
	return cuentaDTO, nil
}
