package main

import (
	"fmt"
	// "io/ioutil"
	// "log"
	"net/http"
	// "strings"
	"time"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Documento struct{
	Numero string `json:"numero"`
	Descripcion string `json:"descripcion"`
	FechaEmision string `json:"fechaEmision"`
	FechaVencimiento string `json:"fechaVencimiento"`
	ImporteDeuda string `json:"importeDeuda"`
	ImporteDeudaMinima string `json:"importeDeudaMinima"`
	IndicadorRestriccPago string `json:"indicadorRestriccPago"`
	CantidadSubconceptos string `json:"cantidadSubconceptos"`
}

type Transaccion struct{
	NumeroReferenciaDeuda string `json:"numeroReferenciaDeuda"`
	NombreCliente string `json:"nombreCliente"`
	NumeroOperacionEmpresa int `json:"numeroOperacionEmpresa"`
	IndMasDeuda int `json:"indMasDeuda"`
	CantidadDocsDeuda int `json:"cantidadDocsDeuda"`
	DatosEmpresa string `json:"datosEmpresa"`
	ListaDocumentos []Documento `json:"listaDocumentos"`
}

type Respuesta struct{
	Codigo int `json:"codigo"`
	Descripcion string `json:"descripcion"`
}

type Operacion struct{
	CodigoOperacion int `json:"codigoOperacion"`
	NumeracionOperacion int `json:"numeracionOperacion"`
	CodigoBanco int `json:"codigoBanco"`
	CodigoConvenio int `json:"codigoConvenio"`
	CanalOperacion string `json:"canalOperacion"`
	CodigoOficina string `json:"codigoOficina"`
	FechaOperacion string `json:"fechaOperacion"`
	HoraOperacion string `json:"horaOperacion"`
}

type Cabecera struct{
	Operacion Operacion
}

type RecaudosRq struct{
	Cabecera Cabecera
}


type Transacciones struct{
	NumeroReferenciaDeuda string `json:"numeroReferenciaDeuda"`
}

type Detalle struct{
	Respuesta Respuesta
	Transaccion Transaccion
}


type ConsultarDeuda struct{
	RecaudosRq RecaudosRq
	Detalle Detalle 
}

func main() {
	time.LoadLocation("America/Lima")
	godotenv.Load();

    fmt.Println("Hello, World!")

	var go_port string = os.Getenv("GO_PORT");

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API GO LANG",
		})
	});

	r.POST("/ConsultarDeuda", func(c *gin.Context) {

		// var ConsultarDeuda ConsultarDeuda{}
		consultarDeuda := ConsultarDeuda{}

		// c.JSON(http.StatusOK, gin.H{
		// 	"message": "API GO LANG",
		// })

		c.IndentedJSON(http.StatusOK, gin.H{"ConsultarDeuda":consultarDeuda})
	});

	r.Run(go_port)
}