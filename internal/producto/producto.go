package producto

import (
	"fmt"
	"time"
)

// Estructura de un producto
type Producto struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	CodeValue   string    `json:"code_value"`
	IsPublished bool      `json:"is_published"`
	Expiration  time.Time `json:"expiration"`
	Price       float64   `json:"price"`
}

// almacena productos
type Storage struct {
	Productos []Producto
}

func (s *Storage) PrintInfo() {
	fmt.Println(s.Productos)
}

// Funcion para retornar todo

func (s *Storage) GetAll() []Producto {
	return s.Productos
}

// funcion para traer productos mayor al precio pasado por parametro
func (s *Storage) GetProductoMayorPrecio(precio float64) []Producto {
	var resultado []Producto // creamos un slice de productos para almacenar

	for _, producto := range s.Productos { // recorre todos los productos que estan en el storage
		if producto.Price >= precio { // pregunta si el producto encontrado, en su propiedad precio es mayor o igual al pasado por parametro
			resultado = append(resultado, producto) //  aprega el producto encontrado al slice resultado
		}

	}
	return resultado

}
