package model

type Detalle struct {
	Respuesta   Respuesta    `json:"respuesta"`
	Transaccion *Transaccion `json:"transaccion,omitempty,selective"`
}
