package database

import (
	"apideprodutos/internal/entity"
	"database/sql"
) //Banco de Dados SQL

type CategoryDB struct { // Conexão com banco
	db *sql.DB
}

func NewCategoryDB(db *sql.DB) *CategoryDB { // Iniciando uma conexão com função
	return &CategoryDB{db: db}
}

func (cd *CategoryDB) GetCategories() ([]*entity.Category, error) { // Metodo que irá retornar uma lista de categorias ou um erro
	rows, err := cd.db.Query("Select id,name FROM categories")

	if err != nil { // validação se teve algum erro
		return nil, err
	}
	defer rows.Close() // Irá fechar o conexão no final de tudo

	var categories []*entity.Category // Variavel que irá ser usada para armazenar uma lista de categorias

	for rows.Next() {
		var category entity.Category                                    // variavel que irá ser usada para armazenar um categoria
		if err := rows.Scan(&category.ID, &category.Name); err != nil { // Aqui estou lendo resultado da consulta no banco e passando para uma variavel, caso de algum erro eu retorno
			return nil, err
		}
		// Se der tudo certo, adiciono na lista a categoria
		categories = append(categories, &category)

	}
	return categories, nil
}

// Metodo para consultar uma categoria
func (cd *CategoryDB) GetCategory(id string) (*entity.Category, error) {
	var category entity.Category // Varivavel que será usada para armazernar informação vinda do banco

	err := cd.db.QueryRow("SELECT id, name FROM cotegories WHERE id = ?", id).Scan(&category.ID, &category.Name) // Consulta no banco e armazeno os dados
	if err != nil {
		return nil, err
	}

	return &category, nil
}

// Metodo para criar uma categoria nova
func (cd *CategoryDB) CreateCategory(category *entity.Category) (string, error) {

	println(category.ID, category.Name)
	_, err := cd.db.Exec("INSERT INTO categories (id,name) VALUES(?,?)", category.ID, category.Name)

	if err != nil {
		println(err.Error())
		return "Error ao salvar no banco", err
	}

	return category.ID, nil
}
