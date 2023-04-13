package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"servicio-rest-bbva/main/src/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	time.LoadLocation("America/Lima")
	godotenv.Load()

	fmt.Println("Hello, World!")

	var go_port string = os.Getenv("GO_PORT")

	r := gin.Default()

	// Middleware para CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))

	// Rutas
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API GO LANG",
		})
	})

	r.POST("/ConsultarDeuda", controller.ObtenerDeuda)

	// r.POST("/ConsultarDeuda", func(c *gin.Context) {
	// 	consultarDeuda := model.ConsultarDeuda{}

	// 	c.IndentedJSON(http.StatusOK, gin.H{"ConsultarDeuda": consultarDeuda})
	// })

	r.Run(go_port)
}
