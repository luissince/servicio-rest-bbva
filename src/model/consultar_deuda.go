package model

type ConsultarDeuda struct {
	RecaudosRq RecaudosRq `json:"recaudosRq"`
	Detalle    Detalle    `json:"detalle"`
}
