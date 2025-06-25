package academia

type AcademiaListarRequest struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Deporte     string `json:"deporte"`
	Entrenador  string `json:"entrenador"`
	Imagen      string `json:"imagen"`
	FechaInicio string `json:"fecha_inicio"`
	FechaFin    string `json:"fecha_fin"`
	EdadMinima  int    `json:"edad_minima"`
	EdadMaxima  int    `json:"edad_maxima"`
	Vacantes    int    `json:"vacantes"`
	Inscritos   int    `json:"inscritos"`
}
