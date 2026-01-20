package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Use map for O(1) lookup instead of O(n) with slice
var products = make(map[int]Product)
var nextID = 1

func init() {
	// Initialize with sample data
	products[1] = Product{ID: 1, Name: "Laptop", Price: 999.99}
	products[2] = Product{ID: 2, Name: "Smartphone", Price: 499.99}
	products[3] = Product{ID: 3, Name: "Tablet", Price: 299.99}
	nextID = 4
}

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
	// Endpoint untuk health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
	})

	// Endpoint untuk mendapatkan daftar produk
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		productList := make([]Product, 0, len(products))
		for _, p := range products {
			productList = append(productList, p)
		}
		json.NewEncoder(w).Encode(productList)
	})

	// Endpoint untuk mendapatkan produk berdasarkan ID
	http.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		idStr := r.URL.Path[len("/products/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		if p, exists := products[id]; exists {
			json.NewEncoder(w).Encode(p)
			return
		}
		http.Error(w, "Product not found", http.StatusNotFound)
	})

	// Endpoint untuk menambahkan produk baru
	http.HandleFunc("/add-product", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var newProduct Product
		err := json.NewDecoder(r.Body).Decode(&newProduct)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		// Auto-generate ID if not provided
		if newProduct.ID == 0 {
			newProduct.ID = nextID
			nextID++
		}
		products[newProduct.ID] = newProduct
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newProduct)
	})

	// Endpoint untuk mengedit produk berdasarkan ID
	http.HandleFunc("/edit-product/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		idStr := r.URL.Path[len("/edit-product/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		var updatedProduct Product
		err = json.NewDecoder(r.Body).Decode(&updatedProduct)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		if _, exists := products[id]; exists {
			products[id] = updatedProduct
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedProduct)
			return
		}
		http.Error(w, "Product not found", http.StatusNotFound)
	})

	// Endpoint untuk menghapus produk berdasarkan ID
	http.HandleFunc("/delete-product/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		idStr := r.URL.Path[len("/delete-product/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		if _, exists := products[id]; exists {
			delete(products, id)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
			return
		}
		http.Error(w, "Product not found", http.StatusNotFound)
	})

	// Menjalankan server di port yang ditentukan
	port := getPort()
	fmt.Printf("Server running di port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("gagal running server:", err)
	}
}