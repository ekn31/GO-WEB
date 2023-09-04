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
