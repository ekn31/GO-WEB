package producto

import (
	"context"
	"errors"
	"time"
)

var (
	ErrEmptyList = errors.New("la lista de productos esta vacia")
	ErrNotFound  = errors.New("Producto no encontrado")
)

// Base de datos en memoria
var (
	productos = []Producto{
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
)

// Interface
type Repository interface {
	GetAll(ctx context.Context) ([]Producto, error) // metodo para obtener todos los productos
	Delete(ctx context.Context, id int) error
}

// Implemetacion de la interface
func NewRepository() Repository {
	return &repository{
		db: productos,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]Producto, error) {
	if len(r.db) < 1 {
		return []Producto{}, ErrEmptyList
	}

	return r.db, nil
}

type repository struct {
	db []Producto
}

func (r *repository) Delete(ctx context.Context, id int) error {

	for key, producto := range r.db {
		if producto.ID == id {
			r.db = append(r.db[:key], r.db[key+1:]...)
			return nil
		}

	}
	return ErrNotFound
}
