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

	_ "kasir-api/docs"

	httpSwagger "github.com/swaggo/http-swagger/v2"
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
	mux := http.NewServeMux()

	// Home page (exact match root only)
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `<html>
		<head><title>Kasir API</title></head>
		<body>
		<h1>Selamat datang di Kasir API</h1>
		<p>Untuk dokumentasi API, kunjungi <a href="/docs/index.html">Swagger UI</a></p>
		</body>
		</html>`)
	})

	// Swagger docs
	mux.Handle("/docs/", httpSwagger.Handler(
		httpSwagger.URL("/docs/doc.json"),
	))

	// Health check
	mux.HandleFunc("GET /health", HealthCheckHandler)

	// Products
	mux.HandleFunc("GET /products", GetProductsHandler)
	mux.HandleFunc("POST /products", AddProductHandler)
	mux.HandleFunc("GET /products/{id}", GetProductByIDHandler)
	mux.HandleFunc("PUT /products/{id}", UpdateProductHandler)
	mux.HandleFunc("DELETE /products/{id}", DeleteProductHandler)

	// Categories
	mux.HandleFunc("GET /categories", GetCategoriesHandler)
	mux.HandleFunc("POST /categories", AddCategoryHandler)
	mux.HandleFunc("GET /categories/{id}", GetCategoryByIDHandler)
	mux.HandleFunc("PUT /categories/{id}", UpdateCategoryHandler)
	mux.HandleFunc("DELETE /categories/{id}", DeleteCategoryHandler)

	port := getPort()
	fmt.Printf("Server running di port %s\n", port)
	fmt.Printf("Swagger UI: http://localhost%s/docs/\n", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		fmt.Println("gagal running server:", err)
	}
}
