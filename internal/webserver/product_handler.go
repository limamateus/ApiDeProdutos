package webserver

import (
	"apideprodutos/internal/entity"
	"apideprodutos/internal/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type WebProductHandler struct {
	ProductService *service.ProductService
}

func NewWebProductHandler(productService *service.ProductService) *WebProductHandler {
	return &WebProductHandler{ProductService: productService}
}

func (wph *WebProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := wph.ProductService.GetProducts()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)

}

func (wph *WebProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id") // Aqui eu pego o id que está sendo passado

	if id == "" { // Aqui estou realizando uma validação se o Id é vazio e retornado um badrequest
		http.Error(w, "id is required", http.StatusBadGateway)
		return
	}
	product, err := wph.ProductService.GetProduct(id) // Aqui estou realizando a consulta

	if err != nil { // Aqui estou realizando uma validação caso teve algum erro
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product) // Aqui eu retornando a produto caso de tudo certo.

}

func (wph *WebProductHandler) GetProductByCategoryId(w http.ResponseWriter, r *http.Request) {
	categoryId := chi.URLParam(r, "categoryId") // Aqui eu pego o id que está sendo passado

	if categoryId == "" { // Aqui estou realizando uma validação se o Id é vazio e retornado um badrequest
		http.Error(w, "categoryId is required", http.StatusBadGateway)
		return
	}

	product, err := wph.ProductService.GetProductByCategoryId(categoryId) // Aqui estou realizando a consulta

	if err != nil { // Aqui estou realizando uma validação caso teve algum erro
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product) // Aqui eu retornando a produto caso de tudo certo.

}
func (wph *WebProductHandler) Createproduct(w http.ResponseWriter, r *http.Request) {

	var product entity.Product

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := wph.ProductService.CreateProduct(product.Name, product.Description, product.CategoryID, product.ImageUrl, product.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}
