package model

type Transaccion struct {
	NumeroReferenciaDeuda  string      `json:"numeroReferenciaDeuda" validate:"required"`
	NombreCliente          string      `json:"nombreCliente"`
	NumeroOperacionEmpresa int         `json:"numeroOperacionEmpresa"`
	IndMasDeuda            int         `json:"indMasDeuda"`
	CantidadDocsDeuda      int         `json:"cantidadDocsDeuda"`
	DatosEmpresa           string      `json:"datosEmpresa"`
	ListaDocumentos        []Documento `json:"listaDocumentos"`

	NumeroDocumento         string  `json:"numeroDocumento,omitempty"`
	ImporteDeudaPagada      float32 `json:"importeDeudaPagada,omitempty"`
	NumeroOperacionRecaudos int     `json:"numeroOperacionRecaudos,omitempty"`
	NumeroOperacionOriginal int     `json:"numeroOperacionOriginal,omitempty"`
	FechaOperacionOriginal  string  `json:"fechaOperacionOriginal,omitempty"`
	FormaPago               string  `json:"formaPago,omitempty"`
	CodigoMoneda            string  `json:"codigoMoneda,omitempty"`
	OtrosDatosEmpresa       string  `json:"otrosDatosEmpresa,omitempty"`
}
