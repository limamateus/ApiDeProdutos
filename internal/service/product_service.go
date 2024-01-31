package service

import (
	"apideprodutos/internal/database"
	"apideprodutos/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(productDB database.ProductDB) *ProductService {
	return &ProductService{ProductDB: productDB}
}

// Serviço para buscar a lista de produtos
func (ps *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProducts()

	if err != nil {
		return nil, err
	}

	return products, nil
}

// Serviço para buscar um produto
func (ps *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := ps.ProductDB.GetProduct(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

// Serviço para buscar um produto pela id da categoria
func (ps *ProductService) GetProductByCategoryId(categoryID string) ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProductByCategoryId(categoryID)

	if err != nil {
		return nil, err
	}

	return products, nil
}

// Serviço para criar um produto
func (ps *ProductService) CreateProduct(name, description, category_id, image_url string, price float64) (*entity.Product, error) {
	Product := entity.NewProduct(name, description, price, category_id, image_url)

	_, err := ps.ProductDB.CreateProduct(Product)

	if err != nil {
		return nil, err
	}

	return Product, nil
}
