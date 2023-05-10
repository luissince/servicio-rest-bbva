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
// @Summary 	 Registrar pago
// @Schemes
// @Description  Proceso para realizar el pago correspondientes
// @Tags 		 Notificar Pago
// @Accept 		 json
// @Produce 	 json
// @Param opcion body  model.BodyNotificarPago true "Estructura para realizar la consulta"
// @Success 	 200  {object}  model.BodyNotificarPagoResponse
// @Failure 	 400  {object}  model.Error
// @Failure 	 500  {object}  model.Error
// @Router /NotificarPago [post]
func NotificarPago(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var body model.BodyNotificarPago

	transaccion := &model.Transaccion{}
	body.NotificarPago.RecaudosRq.Detalle.Transaccion = transaccion

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Cuerpo de solictud vacío.",
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(&body); err != nil {
		consultarDeudaError := model.BodyNotificarPagoResponse{}
		consultarDeudaError.NotificarPagoResponse.RecaudosRs.Detalle.Respuesta.Codigo = helper.CodigoResputas["3000"].Codigo
		consultarDeudaError.NotificarPagoResponse.RecaudosRs.Detalle.Respuesta.Descripcion = helper.CodigoResputas["3000"].Descripcion

		c.IndentedJSON(http.StatusOK, consultarDeudaError)
		return
	}

	bodyResponse := service.NotificarPago(&body.NotificarPago.RecaudosRq)

	c.IndentedJSON(http.StatusOK, bodyResponse)
}
