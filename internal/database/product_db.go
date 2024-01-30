package database

import (
	"apideprodutos/internal/entity"
	"database/sql"
)

type ProductDB struct { // Conexão com banco
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

// Metodo para buscar uma lista de produtos
func (pd *ProductDB) GetProducts() ([]*entity.Product, error) {

	rows, err := pd.db.Query("SELECT id,name,price,category_id FROM products") // Consulta no banco

	if err != nil { // Validação caso teve algum erro
		return nil, err
	}

	defer rows.Close() // defino o fechamento para garantir

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product

		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil

}

// Metodo Para Criar uma api
func (pd *ProductDB) GetProduct(id string) (*entity.Product, error) {
	var product entity.Product

	err := pd.db.QueryRow("SELECT id,name,price,category_id, image_url FROM products WHERE id = ?", id).
		Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID, &product.ImageUrl)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (pd *ProductDB) GetProductByCategoryId(categoryId string) ([]*entity.Product, error) {
	rows, err := pd.db.Query("SELECT id,name,price,category_id,image_url FROM products WHERE category_id = ?", categoryId) // Consulta no banco

	if err != nil { // Validação caso teve algum erro
		return nil, err
	}

	defer rows.Close() // defino o fechamento para garantir

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product

		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

// Metodo para criar uma categoria nova
func (cd *ProductDB) CreateProduct(product *entity.Product) (*entity.Product, error) {
	_, err := cd.db.Exec("INSERT INTO products (id,name,description,price,category_id,image_url) VALUES(?,?,?,?,?,?)", product.ID, product.Name, product.Description, product.Price, product.CategoryID, product.ImageUrl)

	if err != nil {
		return nil, err
	}

	return product, nil
}
