package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// @Summary Get all categories
// @Description Mengambil semua daftar kategori
// @Tags Categories
// @Accept json
// @Produce json
// @Success 200 {array} Category
// @Router /categories [get]
func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

// @Summary Get category by ID
// @Description Mengambil kategori berdasarkan ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} Category
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /categories/{id} [get]
func GetCategoryByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.URL.Path[len("/categories/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if c, exists := categories[id]; exists {
		json.NewEncoder(w).Encode(c)
		return
	}
	http.Error(w, "Category not found", http.StatusNotFound)
}

// @Summary Add new category
// @Description Menambahkan kategori baru
// @Tags Categories
// @Accept json
// @Produce json
// @Param category body Category true "Category data"
// @Success 201 {object} Category
// @Failure 400 {object} map[string]string
// @Router /add-category [post]
func AddCategoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var newCategory Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if newCategory.ID == 0 {
		newCategory.ID = len(categories) + 1
	}
	categories[newCategory.ID] = newCategory
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCategory)
}

// @Summary Update category
// @Description Mengedit kategori berdasarkan ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body Category true "Category data"
// @Success 200 {object} Category
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /edit-category/{id} [put]
func UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	idStr := r.URL.Path[len("/edit-category/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var updatedCategory Category
	err = json.NewDecoder(r.Body).Decode(&updatedCategory)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if _, exists := categories[id]; exists {
		categories[id] = updatedCategory
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedCategory)
		return
	}
	http.Error(w, "Category not found", http.StatusNotFound)
}

// @Summary Delete category
// @Description Menghapus kategori berdasarkan ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /delete-category/{id} [delete]
func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	idStr := r.URL.Path[len("/delete-category/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if _, exists := categories[id]; exists {
		delete(categories, id)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
		return
	}
	http.Error(w, "Category not found", http.StatusNotFound)
}
