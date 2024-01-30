package service

import (
	"apideprodutos/internal/database"
	"apideprodutos/internal/entity"
)

type CategoryService struct {
	CategoryDB database.CategoryDB
}

func NewCategoryService(categoryDB database.CategoryDB) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

// Serviço para buscar a lista de categorias
func (cs *CategoryService) GetCategories() ([]*entity.Category, error) {
	categories, err := cs.CategoryDB.GetCategories()

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (cs *CategoryService) GetCategory(id string) (*entity.Category, error) {
	category, err := cs.CategoryDB.GetCategory(id)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (cs *CategoryService) CreateCategory(name string) (*entity.Category, error) {
	category := entity.NewCategory(name)

	_, err := cs.CategoryDB.CreateCategory(category)

	if err != nil {
		return nil, err
	}

	return category, nil
}
