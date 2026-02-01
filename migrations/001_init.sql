-- =============================================
-- Kasir API Database Schema
-- =============================================

-- Drop tables if exist (untuk fresh install)
-- DROP TABLE IF EXISTS products;
-- DROP TABLE IF EXISTS categories;

-- Create categories table (harus dibuat duluan karena direferensi products)
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create products table
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price INT NOT NULL DEFAULT 0,
    stock INT NOT NULL DEFAULT 0,
    category_id INT REFERENCES categories(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create index for faster queries
CREATE INDEX IF NOT EXISTS idx_products_category_id ON products(category_id);
CREATE INDEX IF NOT EXISTS idx_products_name ON products(name);
CREATE INDEX IF NOT EXISTS idx_categories_name ON categories(name);

-- =============================================
-- Sample Data
-- =============================================

-- Insert sample categories
INSERT INTO categories (name, description) VALUES
    ('Electronics', 'Perangkat elektronik dan gadget'),
    ('Home Appliances', 'Peralatan rumah tangga'),
    ('Food & Beverages', 'Makanan dan minuman');

-- Insert sample products with category
INSERT INTO products (name, price, stock, category_id) VALUES
    ('Laptop', 15000000, 10, 1),
    ('Smartphone', 5000000, 25, 1),
    ('Tablet', 3000000, 15, 1),
    ('Washing Machine', 4500000, 8, 2),
    ('Rice Cooker', 500000, 20, 2),
    ('Mineral Water', 5000, 100, 3);

-- =============================================
-- Migration: Add category_id to existing products table
-- (jalankan ini jika table products sudah ada tanpa category_id)
-- =============================================
-- ALTER TABLE products ADD COLUMN IF NOT EXISTS category_id INT REFERENCES categories(id) ON DELETE SET NULL;
-- ALTER TABLE products ADD COLUMN IF NOT EXISTS created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
-- ALTER TABLE products ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
