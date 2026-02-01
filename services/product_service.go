package services

import (
	"errors"
	"kasir-api/models"
	"kasir-api/repositories"
)

type ProductService struct {
	repo         *repositories.ProductRepository
	categoryRepo *repositories.CategoryRepository
}

func NewProductService(repo *repositories.ProductRepository, categoryRepo *repositories.CategoryRepository) *ProductService {
	return &ProductService{repo: repo, categoryRepo: categoryRepo}
}

func (s *ProductService) GetAll() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) Create(data *models.Product) error {
	// Validasi category_id jika diisi
	if data.CategoryID > 0 {
		_, err := s.categoryRepo.GetByID(data.CategoryID)
		if err != nil {
			return errors.New("category_id tidak ditemukan")
		}
	}
	return s.repo.Create(data)
}

func (s *ProductService) GetByID(id int) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) Update(product *models.Product) error {
	// Validasi category_id jika diisi
	if product.CategoryID > 0 {
		_, err := s.categoryRepo.GetByID(product.CategoryID)
		if err != nil {
			return errors.New("category_id tidak ditemukan")
		}
	}
	return s.repo.Update(product)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
