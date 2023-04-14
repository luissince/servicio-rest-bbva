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

	var body model.Body

	transaccion := &model.Transaccion{
		NumeroReferenciaDeuda: "",
	}
	body.ConsultarDeuda.Detalle.Transaccion = transaccion

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cuerpo de solictud vacío.",
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(&body); err != nil {
		consultarDeudaError := model.ConsultarDeuda{}
		consultarDeudaError.Detalle.Respuesta.Codigo = helper.CodigoResputas["3000"].Codigo
		consultarDeudaError.Detalle.Respuesta.Descripcion = helper.CodigoResputas["3000"].Descripcion

		c.IndentedJSON(http.StatusOK, gin.H{"ConsultarDeuda": consultarDeudaError})
		return
	}

	consultarDeuda, result := service.ObtenerDeuda(body.ConsultarDeuda.Detalle.Transaccion.NumeroReferenciaDeuda)
	if result != "" && result != "ok" {
		c.IndentedJSON(http.StatusInternalServerError, model.Error{
			Message: result,
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"ConsultarDeuda": consultarDeuda})
}
