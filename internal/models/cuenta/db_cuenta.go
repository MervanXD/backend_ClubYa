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

	passwordEncriptado := security.Hash256(cuenta.Contrasena)
	query := "CALL ingesoft.InsertarCuenta(?, ?, ?)"
	var idCuenta int64

	err := database.DB.QueryRow(query, cuenta.Username, cuenta.Email, passwordEncriptado).Scan(&idCuenta)
	if err != nil {
		return 0, err
	}

	return idCuenta, nil
}

func (r *cuentaRepositoryDB) LogIn(cuenta Cuenta) (DTOCuenta, error) {
	passwordEncriptado := security.Hash256(cuenta.Contrasena)
	query := "call ingesoft.LogIn(?, ?,@c_fid_persona,@c_rol,@c_esPostulante, @c_estadoSolicitud, @c_id_membresia, @c_id_solicitud)"
	_, err := database.DB.Exec(query, cuenta.Username, passwordEncriptado)
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
	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	// 1. Insertar persona
	query := "call ingesoft.InsertarPersona(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	_, err = tx.Exec(query,
		cuenta.Persona.Nombre,
		cuenta.Persona.Apellidos,
		cuenta.Persona.Sexo.String(),
		cuenta.Persona.TipoDocumento.String(),
		cuenta.Persona.NroDocumento,
		cuenta.Persona.FechaNacimiento,
		cuenta.Persona.Telefono,
		cuenta.Persona.Pais,
		cuenta.Persona.Provincia,
		cuenta.Persona.Distrito,
		cuenta.Persona.TipoVia.String(),
		cuenta.Persona.Direccion,
		cuenta.Persona.Referencia,
		cuenta.Persona.Ciudad,
		cuenta.Persona.CodigoPostal)
	if err != nil {
		logs.Logger.Println("Error al insertar datos Personales: ", err)
		tx.Rollback()
		return err
	}

	var idPersona int64
	err = tx.QueryRow("SELECT LAST_INSERT_ID()").Scan(&idPersona)
	if err != nil {
		logs.Logger.Println("Error al obtener el ID de la persona: ", err)
		tx.Rollback()
		return err
	}

	// 2. Insertar cuenta admin
	passwordEncriptado := security.Hash256(cuenta.Contrasena)
	query = "CALL ingesoft.InsertarCuentaAdmin(?, ?, ?, ?, ?)"
	_, err = tx.Exec(query, cuenta.Username, cuenta.Email, passwordEncriptado, cuenta.Rol.String(), idPersona)
	if err != nil {
		logs.Logger.Println("Error al crear cuenta de administrador: ", err)
		tx.Rollback()
		return err
	}

	// 3. Commit si todo salió bien
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *cuentaRepositoryDB) ObtenerAdministradores() ([]CuentaAdminRequest, error) {
	query := "CALL ingesoft.ObtenerAdministradores()"
	rows, err := database.DB.Query(query)
	if err != nil {
		logs.Logger.Println("Error al obtener administradores: ", err)
		return nil, err
	}
	defer rows.Close()
	var administradores []CuentaAdminRequest
	for rows.Next() {
		var cuenta CuentaAdminRequest
		err := rows.Scan(
			&cuenta.IdCuenta,
			&cuenta.IdPersona,
			&cuenta.Username,
			&cuenta.Email,
			&cuenta.Rol,
			&cuenta.Nombre,
			&cuenta.Apellidos,
			&cuenta.Sexo,
			&cuenta.TipoDocumento,
			&cuenta.NroDocumento,
			&cuenta.FechaNacimiento,
			&cuenta.Telefono,
			&cuenta.Pais,
			&cuenta.Provincia,
			&cuenta.Distrito,
			&cuenta.TipoVia,
			&cuenta.Direccion,
			&cuenta.Referencia,
			&cuenta.Ciudad,
			&cuenta.CodigoPostal,
		)
		if err != nil {
			logs.Logger.Println("Error al escanear fila de administrador: ", err)
			return nil, err
		}
		administradores = append(administradores, cuenta)
	}
	if err := rows.Err(); err != nil {
		logs.Logger.Println("Error al iterar filas de administradores: ", err)
		return nil, err
	}
	return administradores, nil
}
