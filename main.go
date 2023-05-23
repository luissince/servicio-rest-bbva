package main

import (
	"fmt"
	"net/http"
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

	// Variable de entorno
	var go_port string = os.Getenv("GO_PORT")
	var swagger_host string = os.Getenv("SWAGGER_HOST")

	// Se establece la zana horaria
	time.LoadLocation("America/Lima")

	// Inicializa GIN para correr el servidor
	app := gin.Default()

	// Middleware para CORS
	app.Use(corsMiddleware())

	// Agregar el swagger
	basePath := ""
	docs.SwaggerInfo.BasePath = basePath
	docs.SwaggerInfo.Host = swagger_host

	// Configura ruta para la documentación de Swagger
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Configura ruta para el archivo JSON de Swagger
	app.GET("/swagger.json", func(c *gin.Context) {
		c.JSON(http.StatusOK, docs.SwaggerInfo)
	})

	// Carga las rutas de controladores
	routers.ServicioBbva(app.Group(basePath))

	// Funciona encargada de iniciar el servidor
	fmt.Println("Proceso en ejecución...")
	app.Run(go_port)
}
