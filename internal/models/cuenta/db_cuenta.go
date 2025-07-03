package cuenta

import (
	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/security"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type cuentaRepositoryDB struct{}

func NewCuentaRepositoryDB() CuentaRepository {
	return &cuentaRepositoryDB{}
}

func (r *cuentaRepositoryDB) CrearCuenta(cuenta Cuenta) (int64, error) {
	usernameEncriptado := security.Hash256(cuenta.Username)
	passwordEncriptado := security.Hash256(cuenta.Contrasena)
	query := "CALL ingesoft.InsertarCuenta(?, ?, ?)"
	var idCuenta int64

	err := database.DB.QueryRow(query, usernameEncriptado, cuenta.Email, passwordEncriptado).Scan(&idCuenta)
	if err != nil {
		return 0, err
	}

	return idCuenta, nil
}

func (r *cuentaRepositoryDB) LogIn(cuenta Cuenta) (DTOCuenta, error) {
	usernameEncriptado := security.Hash256(cuenta.Username)
	passwordEncriptado := security.Hash256(cuenta.Contrasena)
	query := "call ingesoft.LogIn(?, ?,@c_fid_persona,@c_rol,@c_esPostulante, @c_estadoSolicitud, @c_id_membresia, @c_id_solicitud)"
	_, err := database.DB.Exec(query, usernameEncriptado, passwordEncriptado)
	var cuentaDTO DTOCuenta
	cuentaDTO.Username = cuenta.Username
	if err != nil {
		logs.Logger.Println("Error al iniciar sesion: ", err)
		return cuentaDTO, err
	}
	err = database.DB.QueryRow("SELECT @c_fid_persona, @c_rol,@c_esPostulante, @c_estadoSolicitud, @c_id_membresia,@c_id_solicitud").Scan(&cuentaDTO.IdPersona, &cuentaDTO.Rol, &cuentaDTO.Postulante, &cuentaDTO.EstadoSolicitud, &cuentaDTO.IdMembresia, &cuentaDTO.IdSolicitud)
	if err != nil {
		logs.Logger.Println("Error al obtener idPersona, rol y postulante: ", err)
		return cuentaDTO, err
	}
	return cuentaDTO, nil
}

func (r *cuentaRepositoryDB) CrearCuentaAdministrador(cuenta Cuenta) error {
	usernameEncriptado := security.Hash256(cuenta.Username)
	passwordEncriptado := security.Hash256(cuenta.Contrasena)
	query := "CALL ingesoft.InsertarCuentaAdministrador(?, ?, ?)"
	_, err := database.DB.Exec(query, usernameEncriptado, cuenta.Email, passwordEncriptado)
	if err != nil {
		logs.Logger.Println("Error al crear cuenta de administrador: ", err)
		return err
	}
	return nil
}
