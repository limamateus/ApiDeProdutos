package webserver

import (
	"apideprodutos/internal/entity"
	"apideprodutos/internal/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type WebCategoryHandler struct {
	CategoryService *service.CategoryService
}

func NewWebCategoryHandler(categoryService *service.CategoryService) *WebCategoryHandler {
	return &WebCategoryHandler{CategoryService: categoryService}
}

func (wch *WebCategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := wch.CategoryService.GetCategories()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categories)

}

func (wch *WebCategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id") // Aqui eu pego o id que está sendo passado

	if id == "" { // Aqui estou realizando uma validação se o Id é vazio e retornado um badrequest
		http.Error(w, "id is required", http.StatusBadGateway)
		return
	}
	category, err := wch.CategoryService.GetCategory(id) // Aqui estou realizando a consulta

	if err != nil { // Aqui estou realizando uma validação caso teve algum erro
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(category) // Aqui eu retornando a categoria caso de tudo certo.

}

func (wch *WebCategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {

	var category entity.Category

	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := wch.CategoryService.CreateCategory(category.Name)
	println(category.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}
