package cuenta

type CuentaRepository interface {
	CrearCuenta(cuenta Cuenta) (int64, error)
	LogIn(cuenta Cuenta) (DTOCuenta, error)
	CrearCuentaAdministrador(cuenta CuentaAdminDTO) error
	ObtenerAdministradores() ([]CuentaAdminRequest, error)
	ObtenerPerfilPorIdCuenta(idCuenta int64) (CuentaAdminDTO, error)
	ActualizarCuentaAParcial(idCuenta int64, dto CuentaAdminUpdateDTO) error
	ObtenerIdCuentaPorPersona(idPersona int64) (int64, error)
	ListarUsuarios() ([]CuentaUsuariosRequest, error)
	RegistrarGmail(cuenta CuentaGmailDTO) (int, error)
	LoginGmail(cuenta CuentaGmailDTO) (DTOCuenta, error)
	ListarCuentasSocios() ([]CuentaSocioRequest, error)
	VisualizarCuentaSocioPorIdCuenta(idCuenta int64) (CuentaSocioDTO, error)

}
