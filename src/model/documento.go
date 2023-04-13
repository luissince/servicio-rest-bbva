package model

type Documento struct {
	Numero                string  `json:"numero"`
	Descripcion           string  `json:"descripcion"`
	FechaEmision          string  `json:"fechaEmision"`
	FechaVencimiento      string  `json:"fechaVencimiento"`
	ImporteDeuda          float32 `json:"importeDeuda"`
	ImporteDeudaMinima    float32 `json:"importeDeudaMinima"`
	IndicadorRestriccPago int     `json:"indicadorRestriccPago"`
	CantidadSubconceptos  int     `json:"cantidadSubconceptos"`
}
