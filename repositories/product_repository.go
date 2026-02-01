package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAll() ([]models.Product, error) {
	query := `SELECT products.id, products.name, products.price, products.stock, 
			  products.category_id, categories.name AS category_name 
			  FROM products 
			  LEFT JOIN categories ON products.category_id = categories.id`
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]models.Product, 0)
	for rows.Next() {
		var p models.Product
		var categoryID sql.NullInt64
		var categoryName sql.NullString
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &categoryID, &categoryName)
		if err != nil {
			return nil, err
		}
		if categoryID.Valid {
			p.CategoryID = int(categoryID.Int64)
		}
		if categoryName.Valid {
			p.CategoryName = categoryName.String
		}
		products = append(products, p)
	}

	return products, nil
}

func (repo *ProductRepository) Create(product *models.Product) error {
	query := "INSERT INTO products (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING id"
	err := repo.db.QueryRow(query, product.Name, product.Price, product.Stock, product.CategoryID).Scan(&product.ID)
	return err
}

// GetByID - ambil produk by ID
func (repo *ProductRepository) GetByID(id int) (*models.Product, error) {
	query := `SELECT products.id, products.name, products.price, products.stock, 
			  products.category_id, categories.name AS category_name 
			  FROM products 
			  LEFT JOIN categories ON products.category_id = categories.id 
			  WHERE products.id = $1`

	var p models.Product
	var categoryID sql.NullInt64
	var categoryName sql.NullString
	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &categoryID, &categoryName)
	if err == sql.ErrNoRows {
		return nil, errors.New("produk tidak ditemukan")
	}
	if err != nil {
		return nil, err
	}
	if categoryID.Valid {
		p.CategoryID = int(categoryID.Int64)
	}
	if categoryName.Valid {
		p.CategoryName = categoryName.String
	}

	return &p, nil
}

func (repo *ProductRepository) Update(product *models.Product) error {
	query := "UPDATE products SET name = $1, price = $2, stock = $3, category_id = $4 WHERE id = $5"
	result, err := repo.db.Exec(query, product.Name, product.Price, product.Stock, product.CategoryID, product.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return nil
}

func (repo *ProductRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return err
}
