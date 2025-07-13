package persona

type FamiliarRepository interface {
	InsertarFamiliar(f Familiar, idTitular int) error
	RegistrarFamiliares(req FamiliarResquest) (int, error)
	ObtenerIdsFamiliaresPorTitular(idTitular int) ([]Familiar, error)
	ObtenerFamiliarPorIDPersona(idPersona int) (*Familiar, error)
	RegistrarFamiliarConSolicitud(f Familiar, idTitular int) error 
}
