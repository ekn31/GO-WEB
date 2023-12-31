package producto

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ekn31/GO-WEB.git/internal/producto/domain/producto"
	"github.com/gin-gonic/gin"
)

type Controlador struct {
	service producto.Service
}

func NewControladorProducto(service producto.Service) *Controlador {
	return &Controlador{
		service: service,
	}

}

func (c *Controlador) GetAll(tokenEnv string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.GetHeader("tokenPostman")

		if tokenHeader == tokenEnv {
			fmt.Println("token valido")
		}
		if tokenHeader != tokenEnv {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"mensaje": "Token invalido",
			})
			return
		}

		productos, err := c.service.GetAll(ctx)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"mensaje": "Internal server error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data:": productos,
		})
	}
}

func (c *Controlador) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"mensaje": "id invalid",
			})
			return
		}
		err = c.service.Delete(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"mensaje": "no se pudo eliminar el producto",
			})
			return
		}
		ctx.JSON(http.StatusNoContent, gin.H{
			"mensaje": "producto eliminado",
		})
	}
}
