// @title Kasir API
// @version 1.0
// @description API untuk aplikasi kasir dengan fitur produk dan kategori
// @host localhost:3000
// @BasePath /
// @schemes http
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"kasir-api/database"
	"kasir-api/docs"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"

	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Config struct {
	Port        string `mapstructure:"PORT"`
	DBConn      string `mapstructure:"DB_CONN"`
	SwaggerHost string `mapstructure:"SWAGGER_HOST"`
}

// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Load config from .env
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("gagal membaca file .env:", err)
			return
		}
	}

	config := Config{
		Port:        viper.GetString("PORT"),
		DBConn:      viper.GetString("DB_CONN"),
		SwaggerHost: viper.GetString("SWAGGER_HOST"),
	}

	// Override Swagger host/scheme for production
	if config.SwaggerHost != "" {
		docs.SwaggerInfo.Host = config.SwaggerHost
		docs.SwaggerInfo.Schemes = []string{"https"}
	}

	// Setup database
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Dependency Injection - Category (create first, needed by Product)
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// Dependency Injection - Product
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo, categoryRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Dependency Injection - Transaction
	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	// Dependency Injection - Report
	reportRepo := repositories.NewReportRepository(db)
	reportService := services.NewReportService(reportRepo)
	reportHandler := handlers.NewReportHandler(reportService)

	mux := http.NewServeMux()

	// Set DB for health check
	SetDB(db)

	// Home page
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
	mux.HandleFunc("GET /health/db", DBHealthCheckHandler)

	// Products routes (layered architecture)
	mux.HandleFunc("/api/produk", productHandler.HandleProducts)
	mux.HandleFunc("/api/produk/", productHandler.HandleProductByID)

	// Categories routes (layered architecture)
	mux.HandleFunc("/api/kategori", categoryHandler.HandleCategories)
	mux.HandleFunc("/api/kategori/", categoryHandler.HandleCategoryByID)

	// Transaction routes
	mux.HandleFunc("/api/checkout", transactionHandler.HandleCheckout)

	// Report routes
	mux.HandleFunc("/api/report/hari-ini", reportHandler.HandleDailyReport)
	mux.HandleFunc("/api/report", reportHandler.HandleReport)

	// Wrap with CORS middleware
	handler := corsMiddleware(mux)

	// Start server
	addr := "0.0.0.0:" + config.Port
	fmt.Printf("Server running di %s\n", addr)
	fmt.Printf("Swagger UI: http://localhost:%s/docs/\n", config.Port)
	err = http.ListenAndServe(addr, handler)
	if err != nil {
		fmt.Println("gagal running server:", err)
	}
}
