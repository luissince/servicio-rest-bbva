package controller

import (
	"net/http"
	"servicio-rest-bbva/main/src/helper"
	"servicio-rest-bbva/main/src/model"
	"servicio-rest-bbva/main/src/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ObtenerDeuda(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var body model.BodyConsultarDeuda

	transaccion := &model.Transaccion{
		NumeroReferenciaDeuda: "",
	}
	body.ConsultarDeuda.RecaudosRq.Detalle.Transaccion = transaccion

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cuerpo de solicitud vacío.",
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
