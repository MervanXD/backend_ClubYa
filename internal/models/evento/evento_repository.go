package evento

type EventoRepository interface{
	ListarEventos() ([]Evento, error)
	BuscarEventoPorID(id int) (*Evento, error) 
	InsertarEvento(req EventoRequest) (int, error)
	ModificarEvento(req EventoRequest) error
	CancelarEvento(idEvento int) error
	EliminarEvento(idEvento int) error
}