package model

type Operacion struct {
	CodigoOperacion     int    `json:"codigoOperacion"`
	NumeracionOperacion int    `json:"numeracionOperacion"`
	CodigoBanco         int    `json:"codigoBanco"`
	CodigoConvenio      int    `json:"codigoConvenio"`
	CanalOperacion      string `json:"canalOperacion"`
	CodigoOficina       string `json:"codigoOficina"`
	FechaOperacion      string `json:"fechaOperacion"`
	HoraOperacion       string `json:"horaOperacion"`
}
