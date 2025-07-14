package persona

type FamiliarConSolicitud struct {
	FidPersona               *int    `json:"fid_persona"`
	Nombre                   *string `json:"nombre"`
	Apellidos                *string `json:"apellidos"`
	Sexo                     *string `json:"sexo"`
	TipoDocumento            *string `json:"tipo_documento"`
	NroDocumento             *string `json:"nro_documento"`
	FechaNacimiento          *string `json:"fecha_nacimiento"`
	Telefono                 *string `json:"telefono"`
	Pais                     *string `json:"pais"`
	Provincia                *string `json:"provincia"`
	Distrito                 *string `json:"distrito"`
	Direccion                *string `json:"direccion"`
	TipoVia                  *string `json:"tipo_via"`
	Referencia               *string `json:"referencia"`
	Ciudad                   *string `json:"ciudad"`
	CodigoPostal             *string `json:"codigo_postal"`
	MismaDireccionPostulante *bool   `json:"misma_direccion_postulante"`
	EsConyuge                *bool   `json:"es_conyuge"`
	TipoFamiliar             *string `json:"tipo_familiar"`
	FechaSolicitud           *string `json:"fecha_solicitud"`
	FechaDecision            *string `json:"fecha_decision"`
	Estado                   *string `json:"estado"`
	Motivo                   *string `json:"motivo"`
	TipoSolicitud            *string `json:"tipo_solicitud"`
}

