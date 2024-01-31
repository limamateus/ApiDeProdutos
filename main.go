package main

import (
	"apideprodutos/internal/database"
	"apideprodutos/internal/service"
	"apideprodutos/internal/webserver"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/produtos")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)

	productDB := database.NewProductDb(db)
	productService := service.NewProductService(*productDB)

	webCategoryHandler := webserver.NewWebCategoryHandler(categoryService)
	webPoductHandler := webserver.NewWebProductHandler(productService)

	c := chi.NewRouter()
	//  Endpoit de Categorias
	c.Use(middleware.Recoverer)                                // Middleware que serve para não deixar morrar api
	c.Use(middleware.Logger)                                   // Para criar logs, então tomar cuidado quando for usar ele
	c.Get("/category/{id}", webCategoryHandler.GetCategory)    // Buscar categoria por id
	c.Get("/category", webCategoryHandler.GetCategories)       // Buscar todas as categorias
	c.Post("/category/new", webCategoryHandler.CreateCategory) // Criar uma categoria

	// Endpoit de Produtos

	c.Get("/product/{id}", webPoductHandler.GetProduct)                              // Buscar um produto por id
	c.Get("/product", webPoductHandler.GetProducts)                                  // Buscar todos os produtos
	c.Get("/product/category/{categoryId}", webPoductHandler.GetProductByCategoryId) // buscar produtos por categoria
	c.Post("/product/new", webPoductHandler.Createproduct)                           // Criar um novo porduto.

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", c)
}
