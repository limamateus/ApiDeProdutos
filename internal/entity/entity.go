package entity

import "github.com/google/uuid"

type Category struct { // Categoria
	ID, Name string
}

func NewCategory(name string) *Category { // Função que irá criar uma categoria
	return &Category{
		ID:   uuid.New().String(),
		Name: name,
	}
}

type Product struct { // Produto
	ID, Name, Description string
	Price                 float64
	CategoryID, ImageUrl  string
}

func NewProduct(name, description string, price float64, categoryID, imageUrl string) *Product {
	return &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
		ImageUrl:    imageUrl,
	}
}
