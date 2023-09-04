package main

import (
	"net/http"
	"strconv"
	"time"

	producto "github.com/ekn31/GO-WEB.git/internal/producto"
	"github.com/gin-gonic/gin"
)

const (
	port = ":8090"
)

func main() {

	storage := producto.Storage{ // le asigono al storage, la lista de prodcutos
		Productos: loadData(),
	}

	storage.PrintInfo()

	// creacion del sevidor

	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		{
			ctx.JSON(http.StatusOK, gin.H{
				"mensaje": "pong",
			})
		}
	})

	// Endpoint para consultar productos por precio
	router.GET("/producto/search", func(ctx *gin.Context) {

		precioQuery := ctx.Query("priceGT")

		if precioQuery != "" {
			precio, err := strconv.ParseFloat(precioQuery, 64) // Convierte el parametro
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"mensaje": "precio invalido",
				})
			}
			data := storage.GetProductoMayorPrecio(precio)
			ctx.JSON(http.StatusOK, gin.H{
				"data": data,
			})
			return

		}

		// si no pasamos parametros, devuelve todo lo que esta en storage
		ctx.JSON(http.StatusOK, gin.H{
			"data": storage.GetAll(),
		})

	})

	router.Run(port)

}

// CREACION DE EL SLICE DE PRODUCTOS
func loadData() []producto.Producto {
	productos := []producto.Producto{
		{
			ID:          1,
			Name:        "Banana",
			CodeValue:   "SD6ASD",
			Quantity:    10,
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       10.0,
		},
		{
			ID:          2,
			Name:        "Manzana",
			CodeValue:   "SD6FSD",
			Quantity:    5,
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       12.0,
		},
		{
			ID:          3,
			Name:        "Pera",
			CodeValue:   "SD6ASD",
			Quantity:    8,
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       7.0,
		},
	}

	return productos
}
