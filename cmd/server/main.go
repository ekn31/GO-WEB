package main

import (
	"net/http"

	handlerProducto "github.com/ekn31/GO-WEB.git/cmd/server/handler/producto"
	"github.com/ekn31/GO-WEB.git/internal/producto/domain/producto"
	"github.com/gin-gonic/gin"
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
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		{
			ctx.JSON(http.StatusOK, gin.H{
				"mensaje": "pong",
			})
		}
	})

	// api para devolver todos los productos
	router.GET("/productos", controlador.GetAll())

	// api para eliminar producto
	router.DELETE("/productos/:id", controlador.Delete())

	// inicializador del servidor
	if err := router.Run(port); err != nil {
		panic(err)
	}
}
