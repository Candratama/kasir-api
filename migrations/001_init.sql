-- Create products table
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price INT NOT NULL DEFAULT 0,
    stock INT NOT NULL DEFAULT 0
);

-- Create categories table
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT
);

-- Insert sample data for products
INSERT INTO products (name, price, stock) VALUES
    ('Laptop', 15000000, 10),
    ('Smartphone', 5000000, 25),
    ('Tablet', 3000000, 15);

-- Insert sample data for categories
INSERT INTO categories (name, description) VALUES
    ('Electronics', 'Perangkat elektronik dan gadget'),
    ('Home Appliances', 'Peralatan rumah tangga');
