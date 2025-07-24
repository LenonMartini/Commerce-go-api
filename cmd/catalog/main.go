package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/LenonMartini/Commerce-go-api/GOAPI/internal/database"
	"github.com/LenonMartini/Commerce-go-api/GOAPI/internal/service"
	"github.com/LenonMartini/Commerce-go-api/GOAPI/internal/webserver"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/imersao")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)

	productDB := database.NewProductDB(db)
	productService := service.NewProductService(*productDB)

	webCategoryHandler := webserver.NewWebCategoryHandler(categoryService)
	webProductHandler := webserver.NewWebProductHandler(productService)

	/*Mapeamento de URLS*/
	c := chi.NewRouter()
	/*Usar middlewares*/
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)
	/*Rotas de Categorias*/
	c.Get("/category/{id}", webCategoryHandler.GetCategory)
	c.Get("/category", webCategoryHandler.GetCategories)
	c.Post("/category", webCategoryHandler.CreateCategory)

	/*Rotas de Produtos*/
	c.Get("/product/{id}", webProductHandler.GetProduct)
	c.Get("/product", webProductHandler.GetProducts)
	c.Get("/product/category/{categoryID}", webProductHandler.GetProductsByCategoryID)
	c.Post("/product", webProductHandler.CreateProduct)

	fmt.Println("Server is runing on port 8000")
	http.ListenAndServe(":8000", c)
}
