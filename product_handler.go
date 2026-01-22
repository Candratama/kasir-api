package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// @Summary Get all products
// @Description Mengambil semua daftar produk
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {array} Product
// @Router /products [get]
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// @Summary Get product by ID
// @Description Mengambil produk berdasarkan ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} Product
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /products/{id} [get]
func GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if p, exists := products[id]; exists {
		json.NewEncoder(w).Encode(p)
		return
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}

// @Summary Add new product
// @Description Menambahkan produk baru
// @Tags Products
// @Accept json
// @Produce json
// @Param product body Product true "Product data"
// @Success 201 {object} Product
// @Failure 400 {object} map[string]string
// @Router /products [post]
func AddProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if newProduct.ID == 0 {
		newProduct.ID = nextID
		nextID++
	}
	products[newProduct.ID] = newProduct
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProduct)
}

// @Summary Update product
// @Description Mengedit produk berdasarkan ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body Product true "Product data"
// @Success 200 {object} Product
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /products/{id} [put]
func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
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
		updatedProduct.ID = id
		products[id] = updatedProduct
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedProduct)
		return
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}

// @Summary Delete product
// @Description Menghapus produk berdasarkan ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /products/{id} [delete]
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
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
}
