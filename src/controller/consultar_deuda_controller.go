package controller

import (
	"net/http"
	"servicio-rest-bbva/src/helper"
	"servicio-rest-bbva/src/model"
	"servicio-rest-bbva/src/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// PingExample   godoc
// @Summary 	 Ruta inicial
// @Schemes
// @Description  Ruta para testear el punto inicial
// @Tags 		 Inicio
// @Accept 		 json
// @Produce 	 json
// @Success 	 200  {string}  string "API GO LANG"
// @Router / [get]
func InitApp(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "API GO LANG",
	})
}

// PingExample   godoc
// @Summary 	 Obtener lista de deudas del alumno
// @Schemes
// @Description  Obtener el listado de deudas del alumno con su código o dni
// @Tags 		 Consultar
// @Accept 		 json
// @Produce 	 json
// @Param opcion body  model.BodyConsultarDeuda true "Estructura para realizar la consulta"
// @Success 	 200  {object}  model.BodyConsultarDeudaResponse
// @Failure 	 400  {object}  model.Error
// @Failure 	 500  {object}  model.Error
// @Router /ConsultarDeuda [post]
func ObtenerDeuda(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var body model.BodyConsultarDeuda

	transaccion := &model.Transaccion{
		NumeroReferenciaDeuda: "",
	}
	body.ConsultarDeuda.RecaudosRq.Detalle.Transaccion = transaccion

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Cuerpo de solicitud vacío.",
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(&body); err != nil {
		consultarDeudaError := model.BodyConsultarDeudaResponse{}
		consultarDeudaError.ConsultarDeudaResponse.RecaudosRs.Detalle.Respuesta.Codigo = helper.CodigoResputas["3000"].Codigo
		consultarDeudaError.ConsultarDeudaResponse.RecaudosRs.Detalle.Respuesta.Descripcion = helper.CodigoResputas["3000"].Descripcion

		c.IndentedJSON(http.StatusOK, consultarDeudaError)
		return
	}

	bodyResponse := service.ObtenerDeuda(body.ConsultarDeuda.RecaudosRq.Detalle.Transaccion.NumeroReferenciaDeuda)
	c.IndentedJSON(http.StatusOK, bodyResponse)
}
