package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	handlerProducto "github.com/ekn31/GO-WEB.git/cmd/server/handler/producto"
	"github.com/ekn31/GO-WEB.git/internal/producto/domain/producto"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const (
	port = ":8090"
)

func main() {

	// Creacion del repositorio
	repository := producto.NewRepository()
	service := producto.NewService(repository)
	controlador := handlerProducto.NewControladorProducto(service)

	// creacion del sevidor
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	token := os.Getenv("TOKEN")
	fmt.Println(token)

	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		{
			ctx.JSON(http.StatusOK, gin.H{
				"mensaje": "pong",
			})
		}
	})

	// api para devolver todos los productos
	router.GET("/productos", controlador.GetAll(token))

	// api para eliminar producto
	router.DELETE("/productos/:id", controlador.Delete())

	// inicializador del servidor
	if err := router.Run(port); err != nil {
		panic(err)
	}
}
