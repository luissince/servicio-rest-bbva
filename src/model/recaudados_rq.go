package model

type RecaudosRq struct {
	Cabecera Cabecera `json:"cabecera"`
	Detalle  Detalle  `json:"detalle"`
}
