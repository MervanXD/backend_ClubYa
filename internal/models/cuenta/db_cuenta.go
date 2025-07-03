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

func (r *cuentaRepositoryDB) CrearCuentaAdministrador(cuenta CuentaAdminDTO) error {
	query := "call ingesoft.InsertarPersona(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	_, err := database.DB.Exec(query, cuenta.Persona.Nombre, cuenta.Persona.Apellidos, cuenta.Persona.Sexo.String(),
		cuenta.Persona.TipoDocumento.String(), cuenta.Persona.NroDocumento, cuenta.Persona.FechaNacimiento,
		cuenta.Persona.Telefono, cuenta.Persona.Pais, cuenta.Persona.Provincia, cuenta.Persona.Distrito,
		cuenta.Persona.TipoVia.String(),
		cuenta.Persona.Direccion, cuenta.Persona.Referencia, cuenta.Persona.CodigoPostal)
	if err != nil {
		logs.Logger.Println("Error al insertar Familiar: ", err)
		return err
	}

	var idPersona int64
	err = database.DB.QueryRow("SELECT LAST_INSERT_ID()").Scan(&idPersona)
	if err != nil {
		logs.Logger.Println("Error al obtener el ID de la persona: ", err)
		return err
	}

	usernameEncriptado := security.Hash256(cuenta.Username)
	passwordEncriptado := security.Hash256(cuenta.Contrasena)
	query = "CALL ingesoft.InsertarCuentaAdmin(?, ?, ?, ?, ?)"
	_, err = database.DB.Exec(query, usernameEncriptado, cuenta.Email, passwordEncriptado, cuenta.Rol.String(), idPersona)
	if err != nil {
		logs.Logger.Println("Error al crear cuenta de administrador: ", err)
		return err
	}

	return nil
}
