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
// @Summary 	 Extornar pago
// @Schemes
// @Description  Proceso para realizar el extorno del pago
// @Tags 		 Extornar Pago
// @Accept 		 json
// @Produce 	 json
// @Param opcion body  model.BodyExtornarPago true "Estructura para realizar la consulta"
// @Success 	 200  {object}  model.BodyExtornarPagoResponse
// @Failure 	 400  {object}  model.Error
// @Failure 	 500  {object}  model.Error
// @Router /ExtornarPago [post]
func ExtornarPago(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var body model.BodyExtornarPago

	transaccion := &model.Transaccion{}
	body.ExtornarPago.RecaudosRq.Detalle.Transaccion = transaccion

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Cuerpo de solictud vacío.",
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(&body); err != nil {
		consultarDeudaError := model.BodyExtornarPagoResponse{}
		consultarDeudaError.ExtornarPagoResponse.RecaudosRs.Detalle.Respuesta.Codigo = helper.CodigoResputas["3000"].Codigo
		consultarDeudaError.ExtornarPagoResponse.RecaudosRs.Detalle.Respuesta.Descripcion = helper.CodigoResputas["3000"].Descripcion

		c.IndentedJSON(http.StatusOK, consultarDeudaError)
		return
	}

	bodyResponse := service.ExtornarPago(&body.ExtornarPago.RecaudosRq)

	c.IndentedJSON(http.StatusOK, bodyResponse)
}
