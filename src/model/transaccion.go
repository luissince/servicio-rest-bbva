package model

type Transaccion struct {
	// NumeroReferenciaDeuda string `json:"numeroReferenciaDeuda"`
	NumeroReferenciaDeuda  string      `json:"numeroReferenciaDeuda" validate:"required"`
	NombreCliente          string      `json:"nombreCliente"`
	NumeroOperacionEmpresa int         `json:"numeroOperacionEmpresa"`
	IndMasDeuda            int         `json:"indMasDeuda"`
	CantidadDocsDeuda      int         `json:"cantidadDocsDeuda"`
	DatosEmpresa           string      `json:"datosEmpresa"`
	ListaDocumentos        []Documento `json:"listaDocumentos"`
}
