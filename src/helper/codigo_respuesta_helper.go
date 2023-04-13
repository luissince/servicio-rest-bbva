package helper

type CodigoRespuesta struct {
	Codigo      string
	Descripcion string
}

var CodigoResputas = map[string]CodigoRespuesta{
	"0001": CodigoRespuesta{
		Codigo:      "0001",
		Descripcion: "TRANSACCION REALIZADA CON EXITO",
	},
	"3000": CodigoRespuesta{
		Codigo:      "3000",
		Descripcion: "FORMATO DE TRAMA NO VALIDO",
	},
	"3002": CodigoRespuesta{
		Codigo:      "3002",
		Descripcion: "NO SE PUEDE REALIZAR LA TRANSACCION",
	},
	"3004": CodigoRespuesta{
		Codigo:      "3004",
		Descripcion: "NO SE PUEDE REALIZAR EL REGISTRO DE EXTORNO",
	},
	"3009": CodigoRespuesta{
		Codigo:      "3009",
		Descripcion: "NO TIENE DEUDAS PENDIENTES",
	},
	"0101": CodigoRespuesta{
		Codigo:      "0101",
		Descripcion: "NUMERO DE REFERENCIA NO EXISTE",
	},
	"0102": CodigoRespuesta{
		Codigo:      "0102",
		Descripcion: "NUMERO DE REFERENCIA EXPIRADA",
	},
	"0106": CodigoRespuesta{
		Codigo:      "0106",
		Descripcion: "NUMERO DE REFERENCIA CON ESTADO PAGADO",
	},
	"0290": CodigoRespuesta{
		Codigo:      "0290",
		Descripcion: "ERROR DEBE PAGAR LA CUOTA MAS ANTIGUA",
	},
	"3013": CodigoRespuesta{
		Codigo:      "3013",
		Descripcion: "ESTADO DE NRO DE REFERENCIA NO VALIDO",
	},
	"3014": CodigoRespuesta{
		Codigo:      "3014",
		Descripcion: "EXTORNO NO PRECESADO PORQUE NO EXISTE REGISTRO DE PAGO",
	},
	"3051": CodigoRespuesta{
		Codigo:      "3051",
		Descripcion: "MONTO DE PAGO DEBE SER MINIMO O MAXIMO",
	},
}
