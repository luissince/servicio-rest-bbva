package routers

import (
	"github.com/gin-gonic/gin"

	"servicio-rest-bbva/src/controller"
)

func ServicioBbva(router *gin.RouterGroup) {
	router.GET("/", controller.InitApp)
	router.POST("/ConsultarDeuda", controller.ObtenerDeuda)
	router.POST("/NotificarPago", controller.NotificarPago)
	router.POST("/ExtornarPago", controller.ExtornarPago)
}
