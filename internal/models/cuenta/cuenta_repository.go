package cuenta

type CuentaRepository interface {
	CrearCuenta(cuenta Cuenta) (int64, error)
	LogIn(cuenta Cuenta) (DTOCuenta, error)
	CrearCuentaAdministrador(cuenta CuentaAdminDTO) (int64, error)
}
