// @title Kasir API
// @version 1.0
// @description API untuk aplikasi kasir dengan fitur produk dan kategori
// @host localhost:3000
// @BasePath /
// @schemes http
package main

import (
	"fmt"
	"net/http"
	"os"

	httpSwagger "github.com/swaggo/http-swagger/v2"
	_ "kasir-api/docs"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}
	return port
}

func main() {
	http.Handle("/docs/", httpSwagger.Handler(
		httpSwagger.URL("/docs/doc.json"),
	))

	http.HandleFunc("/health", HealthCheckHandler)
	http.HandleFunc("/products", GetProductsHandler)
	http.HandleFunc("/products/", GetProductByIDHandler)
	http.HandleFunc("/add-product", AddProductHandler)
	http.HandleFunc("/edit-product/", UpdateProductHandler)
	http.HandleFunc("/delete-product/", DeleteProductHandler)
	http.HandleFunc("/categories", GetCategoriesHandler)
	http.HandleFunc("/categories/", GetCategoryByIDHandler)
	http.HandleFunc("/add-category", AddCategoryHandler)
	http.HandleFunc("/edit-category/", UpdateCategoryHandler)
	http.HandleFunc("/delete-category/", DeleteCategoryHandler)

	port := getPort()
	fmt.Printf("Server running di port %s\n", port)
	fmt.Printf("Swagger UI: http://localhost%s/docs/\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("gagal running server:", err)
	}
}
