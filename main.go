package main

import (
	"os"
	"servicio-rest-bbva/src/routers"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	docs "servicio-rest-bbva/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		c.Next()
	}
}

func main() {
	// Cargar las variables de entorno
	godotenv.Load()

	//
	var go_port string = os.Getenv("GO_PORT")

	// Se establece la zana horaria
	time.LoadLocation("America/Lima")

	// Inicializa GIN para correr el servidor
	app := gin.Default()

	// Middleware para CORS
	app.Use(corsMiddleware())

	// Agregar el swagger
	basePath := "/v1"
	docs.SwaggerInfo.BasePath = basePath

	// app.POST("/ConsultarDeuda", controller.ObtenerDeuda)
	// app.POST("/NotificarPago", controller.NotificarPago)
	// app.POST("/ExtornarPago", controller.ExtornarPago)

	// Carga de rutas
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routers.ServicioBbva(app.Group(basePath))

	// Funciona encargada de iniciar el servidor
	app.Run(go_port)
}
