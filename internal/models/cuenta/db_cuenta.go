package cuenta

import (
	"fmt"
	"strings"

	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/security"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/go-sql-driver/mysql"
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
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				// Error 1062: Duplicate entry
				if strings.Contains(mysqlErr.Message, "email") {
					return -1, fmt.Errorf("el DNI ya está registrado")
				}
				if strings.Contains(mysqlErr.Message, "username") {
					return -1, fmt.Errorf("el nombre de usuario ya está registrado")
				}

			}
		}
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

	// Función para manejar el rollback de forma segura
	rollbackSafe := func() {
		if err := tx.Rollback(); err != nil {
			logs.Logger.Println("Error en rollback: ", err)
		}
	}

	defer func() {
		if p := recover(); p != nil {
			rollbackSafe()
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
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				// Error 1062: Duplicate entry
				if strings.Contains(mysqlErr.Message, "nroDocumento") {
					rollbackSafe()
					return fmt.Errorf("el nroDocumento ya está registrado")
				}

			}
		}
		rollbackSafe()
		return err
	}

	var idPersona int64
	err = tx.QueryRow("SELECT LAST_INSERT_ID()").Scan(&idPersona)
	if err != nil {
		logs.Logger.Println("Error al obtener el ID de la persona: ", err)
		rollbackSafe()
		return err
	}

	// 2. Insertar cuenta admin
	passwordEncriptado := security.Hash256(cuenta.Contrasena)
	query = "CALL ingesoft.InsertarCuentaAdmin(?, ?, ?, ?, ?)"
	_, err = tx.Exec(query, cuenta.Username, cuenta.Email, passwordEncriptado, cuenta.Rol.String(), idPersona)
	if err != nil {
		logs.Logger.Println("Error al crear cuenta de administrador: ", err)
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				// Error 1062: Duplicate entry
				if strings.Contains(mysqlErr.Message, "email") {
					rollbackSafe()
					return fmt.Errorf("el DNI ya está registrado")
				}
				if strings.Contains(mysqlErr.Message, "username") {
					rollbackSafe()
					return fmt.Errorf("el nombre de usuario ya está registrado")
				}

			}
		}
		rollbackSafe()
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

func (r *cuentaRepositoryDB) ObtenerPerfilPorIdCuenta(idCuenta int64) (CuentaAdminDTO, error) {
	query := "CALL ingesoft.ObtenerPerfilPorIdCuenta(?)"
	row := database.DB.QueryRow(query, idCuenta)

	var cuentaDTO CuentaAdminDTO
	err := row.Scan(
		&cuentaDTO.Username,
		&cuentaDTO.Email,
		&cuentaDTO.Rol,
		&cuentaDTO.Persona.Id,
		&cuentaDTO.Persona.Nombre,
		&cuentaDTO.Persona.Apellidos,
		&cuentaDTO.Persona.Sexo,
		&cuentaDTO.Persona.TipoDocumento,
		&cuentaDTO.Persona.NroDocumento,
		&cuentaDTO.Persona.FechaNacimiento,
		&cuentaDTO.Persona.Telefono,
		&cuentaDTO.Persona.Pais,
		&cuentaDTO.Persona.Provincia,
		&cuentaDTO.Persona.Distrito,
		&cuentaDTO.Persona.TipoVia,
		&cuentaDTO.Persona.Direccion,
		&cuentaDTO.Persona.Referencia,
		&cuentaDTO.Persona.Ciudad,
		&cuentaDTO.Persona.CodigoPostal,
		&cuentaDTO.Persona.Ocupacion,
		&cuentaDTO.Persona.NombreEmpresa,
		&cuentaDTO.Persona.DireccionEmpresa,
		&cuentaDTO.Persona.IngresoPromedio,
	)
	if err != nil {
		logs.Logger.Println("Error al obtener perfil por ID de cuenta: ", err)
		return cuentaDTO, err
	}

	return cuentaDTO, nil
}

func (r *cuentaRepositoryDB) ActualizarCuentaAParcial(idCuenta int64, dto CuentaAdminUpdateDTO) error {
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
	// 1. Actualizar tabla Cuenta
	cuentaSet := []string{}
	cuentaArgs := []interface{}{}

	if dto.Username != nil {
		cuentaSet = append(cuentaSet, "username = ?")
		cuentaArgs = append(cuentaArgs, *dto.Username)
	}
	if dto.Email != nil {
		cuentaSet = append(cuentaSet, "email = ?")
		cuentaArgs = append(cuentaArgs, *dto.Email)
	}
	if dto.Contrasena != nil {
		passwordEncriptado := security.Hash256(*dto.Contrasena)
		cuentaSet = append(cuentaSet, "contrasena = ?")
		cuentaArgs = append(cuentaArgs, passwordEncriptado)
	}
	if dto.Rol != nil {
		cuentaSet = append(cuentaSet, "rol = ?")
		cuentaArgs = append(cuentaArgs, dto.Rol.String())
	}

	if len(cuentaSet) > 0 {
		query := "UPDATE Cuenta SET " + strings.Join(cuentaSet, ", ") + " WHERE id = ?"
		cuentaArgs = append(cuentaArgs, idCuenta)
		_, err := tx.Exec(query, cuentaArgs...)
		if err != nil {
			logs.Logger.Println("Error al actualizar cuenta: ", err)
			tx.Rollback()
			return err
		}
	}

	// 2. Actualizar tabla Persona (si se proporcionaron datos)
	if dto.Titular != nil {
		// Primero obtener el id_persona de la cuenta
		var idPersona int
		err := tx.QueryRow("SELECT fid_Persona FROM Cuenta WHERE id = ?", idCuenta).Scan(&idPersona)
		if err != nil {
			logs.Logger.Println("Error al obtener id_persona: ", err)
			tx.Rollback()
			return err
		}

		personaSet := []string{}
		personaArgs := []interface{}{}

		if dto.Titular.Nombre != nil {
			personaSet = append(personaSet, "nombres = ?")
			personaArgs = append(personaArgs, *dto.Titular.Nombre)
		}
		if dto.Titular.Apellidos != nil {
			personaSet = append(personaSet, "apellidos = ?")
			personaArgs = append(personaArgs, *dto.Titular.Apellidos)
		}
		if dto.Titular.Sexo != nil {
			personaSet = append(personaSet, "sexo = ?")
			personaArgs = append(personaArgs, dto.Titular.Sexo.String())
		}
		if dto.Titular.TipoDocumento != nil {
			personaSet = append(personaSet, "tipoDocumento = ?")
			personaArgs = append(personaArgs, dto.Titular.TipoDocumento.String())
		}
		if dto.Titular.NroDocumento != nil {
			personaSet = append(personaSet, "nroDocumento = ?")
			personaArgs = append(personaArgs, *dto.Titular.NroDocumento)
		}
		if dto.Titular.FechaNacimiento != nil {
			personaSet = append(personaSet, "fechaNacimiento = ?")
			personaArgs = append(personaArgs, *dto.Titular.FechaNacimiento)
		}
		if dto.Titular.Telefono != nil {
			personaSet = append(personaSet, "telefono = ?")
			personaArgs = append(personaArgs, *dto.Titular.Telefono)
		}
		if dto.Titular.Pais != nil {
			personaSet = append(personaSet, "pais = ?")
			personaArgs = append(personaArgs, *dto.Titular.Pais)
		}
		if dto.Titular.Provincia != nil {
			personaSet = append(personaSet, "provincia = ?")
			personaArgs = append(personaArgs, *dto.Titular.Provincia)
		}
		if dto.Titular.Distrito != nil {
			personaSet = append(personaSet, "distrito = ?")
			personaArgs = append(personaArgs, *dto.Titular.Distrito)
		}
		if dto.Titular.TipoVia != nil {
			personaSet = append(personaSet, "tipoVia = ?")
			personaArgs = append(personaArgs, dto.Titular.TipoVia.String())
		}
		if dto.Titular.Direccion != nil {
			personaSet = append(personaSet, "direccion = ?")
			personaArgs = append(personaArgs, *dto.Titular.Direccion)
		}
		if dto.Titular.Referencia != nil {
			personaSet = append(personaSet, "referencia = ?")
			personaArgs = append(personaArgs, *dto.Titular.Referencia)
		}
		if dto.Titular.Ciudad != nil {
			personaSet = append(personaSet, "ciudad = ?")
			personaArgs = append(personaArgs, *dto.Titular.Ciudad)
		}
		if dto.Titular.CodigoPostal != nil {
			personaSet = append(personaSet, "codigoPostal = ?")
			personaArgs = append(personaArgs, *dto.Titular.CodigoPostal)
		}

		if len(personaSet) > 0 {
			query := "UPDATE Persona SET " + strings.Join(personaSet, ", ") + " WHERE id = ?"
			personaArgs = append(personaArgs, idPersona)
			_, err := tx.Exec(query, personaArgs...)
			if err != nil {
				logs.Logger.Println("Error al actualizar persona: ", err)
				tx.Rollback()
				return err
			}
		}
	}

	// 3. Commit si todo salió bien
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *cuentaRepositoryDB) ObtenerIdCuentaPorPersona(idPersona int64) (int64, error) {
	var idCuenta int64

	err := database.DB.QueryRow("CALL ingesoft.ObtenerCuentaPorPersona(?)", idPersona).Scan(&idCuenta)
	if err != nil {
		logs.Logger.Println("Error al ejecutar procedure ObtenerCuentaPorPersona: ", err)
		return 0, err
	}

	return idCuenta, nil
}

func (r *cuentaRepositoryDB) ListarUsuarios() ([]CuentaUsuariosRequest, error) {
	query := "CALL ingesoft.ListarUsuarios()"
	rows, err := database.DB.Query(query)
	if err != nil {
		logs.Logger.Println("Error al obtener usuarios: ", err)
		return nil, err
	}
	defer rows.Close()
	var usuarios []CuentaUsuariosRequest
	for rows.Next() {
		var cuenta CuentaUsuariosRequest
		err := rows.Scan(
			&cuenta.IdCuenta,
			&cuenta.Username,
			&cuenta.Email,
			&cuenta.Nombre,
			&cuenta.NroDocumento,
			&cuenta.Telefono,
			&cuenta.Rol,
		)
		if err != nil {
			logs.Logger.Println("Error al escanear fila de usuario: ", err)
			return nil, err
		}
		usuarios = append(usuarios, cuenta)
	}
	if err := rows.Err(); err != nil {
		logs.Logger.Println("Error al iterar filas de usuario: ", err)
		return nil, err
	}
	return usuarios, nil
}

func (r *cuentaRepositoryDB) RegistrarGmail(cuenta CuentaGmailDTO) (int, error) {
	if !cuenta.EmailVerified {
		logs.Logger.Println("Error: El correo electrónico no ha sido verificado")
		return -1, nil
	}
	query := "CALL ingesoft.RegistrarCuentaGmail(?, ?,@p_id)"
	_, err := database.DB.Exec(query, cuenta.Username, cuenta.Email)
	if err != nil {
		logs.Logger.Println("Error al registrar cuenta Gmail: ", err)
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				// Error 1062: Duplicate entry
				if strings.Contains(mysqlErr.Message, "email") {
					return -1, fmt.Errorf("el DNI ya está registrado")
				}
				if strings.Contains(mysqlErr.Message, "username") {
					return -1, fmt.Errorf("el nombre de usuario ya está registrado")
				}

			}
		}
		return -1, err
	}

	var idCuenta int
	err = database.DB.QueryRow("SELECT @p_id").Scan(&idCuenta)
	if err != nil {
		logs.Logger.Println("Error al obtener el ID de la cuenta Gmail: ", err)
		return -1, err
	}

	return idCuenta, nil
}

func (r *cuentaRepositoryDB) LoginGmail(cuenta CuentaGmailDTO) (DTOCuenta, error) {
	if !cuenta.EmailVerified {
		logs.Logger.Println("Error: El correo electrónico no ha sido verificado")
		return DTOCuenta{}, nil
	}
	query := "CALL ingesoft.LoginGmail(?,@c_fid_persona,@c_rol,@c_esPostulante, @c_estadoSolicitud, @c_id_membresia, @c_id_solicitud)"
	_, err := database.DB.Exec(query, cuenta.Email)
	if err != nil {
		logs.Logger.Println("Error al iniciar sesión con Gmail: ", err)
		return DTOCuenta{}, err
	}
	var cuentaDTO DTOCuenta

	err = database.DB.QueryRow("SELECT @c_fid_persona, @c_rol,@c_esPostulante, @c_estadoSolicitud, @c_id_membresia,@c_id_solicitud").Scan(&cuentaDTO.IdPersona, &cuentaDTO.Rol, &cuentaDTO.Postulante, &cuentaDTO.EstadoSolicitud, &cuentaDTO.IdMembresia, &cuentaDTO.IdSolicitud)
	if err != nil {
		logs.Logger.Println("Error al obtener idPersona, rol y postulante: ", err)
		return cuentaDTO, err
	}
	return cuentaDTO, nil
}

func (r *cuentaRepositoryDB) ListarCuentasSocios() ([]CuentaSocioRequest, error) {
	query := "CALL ingesoft.ListarCuentasSocios()"
	rows, err := database.DB.Query(query)
	if err != nil {
		logs.Logger.Println("Error al obtener cuentas de socios: ", err)
		return nil, err
	}
	defer rows.Close()
	var cuentas []CuentaSocioRequest
	for rows.Next() {
		var cuenta CuentaSocioRequest
		err := rows.Scan(
			&cuenta.IDCuenta,
			&cuenta.Username,
			&cuenta.Email,
			&cuenta.Nombre,
			&cuenta.Apellidos,
			&cuenta.FechaInicioMembresia,
			&cuenta.EstadoMembresia,
			&cuenta.Activo,
		)
		if err != nil {
			logs.Logger.Println("Error al escanear fila de cuenta de socio: ", err)
			return nil, err
		}
		cuentas = append(cuentas, cuenta)
	}
	if err := rows.Err(); err != nil {
		logs.Logger.Println("Error al iterar filas de cuentas de socios: ", err)
		return nil, err
	}
	return cuentas, nil

}

func (r *cuentaRepositoryDB)VisualizarCuentaSocioPorIdCuenta(idCuenta int64) (CuentaSocioDTO, error){
	query := "CALL ingesoft.VisualizarCuentaSocioPorIdCuenta(?)"
	row := database.DB.QueryRow(query, idCuenta)

	var cuentaDTO CuentaSocioDTO
	err := row.Scan(
		&cuentaDTO.IdCuenta,
		&cuentaDTO.IdPersona,
		&cuentaDTO.Username,
		&cuentaDTO.Email,
		&cuentaDTO.Rol,
		&cuentaDTO.Nombre,
		&cuentaDTO.Apellidos,
		&cuentaDTO.Sexo,
		&cuentaDTO.TipoDocumento,
		&cuentaDTO.NroDocumento,
		&cuentaDTO.FechaNacimiento,
		&cuentaDTO.Telefono,
		&cuentaDTO.Pais,
		&cuentaDTO.Provincia,
		&cuentaDTO.Distrito,
		&cuentaDTO.TipoVia,
		&cuentaDTO.Direccion,
		&cuentaDTO.Referencia,
		&cuentaDTO.Ciudad,
		&cuentaDTO.CodigoPostal,
		&cuentaDTO.FechaInicioMembresia,
		&cuentaDTO.EstadoMembresia,
	)
	if err != nil {
		logs.Logger.Println("Error al obtener cuenta de socio por ID de cuenta: ", err)
		return cuentaDTO, err
	}

	return cuentaDTO, nil
}
