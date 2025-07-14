package tarifas

type TarifaMora struct {
	IdTarifaMora       int     `json:"id_tarifa_mora"`
	FidTarifaMembresia int     `json:"id_tarifa_membresia"`
	NombreMora         string  `json:"nombre_mora"`
	CostoMora          float64 `json:"costo_mora"`
	Estado             bool    `json:"estado"`
}
