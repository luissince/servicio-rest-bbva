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
	})

	r.Run(go_port)
}