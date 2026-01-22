package main

var products = make(map[int]Product)
var categories = make(map[int]Category)
var nextID = 1

func init() {
	products[1] = Product{ID: 1, Name: "Laptop", Price: 999.99}
	products[2] = Product{ID: 2, Name: "Smartphone", Price: 499.99}
	products[3] = Product{ID: 3, Name: "Tablet", Price: 299.99}

	categories[1] = Category{ID: 1, Name: "Electronics", Description: "Devices and gadgets"}
	categories[2] = Category{ID: 2, Name: "Home Appliances", Description: "Appliances for home use"}
}
